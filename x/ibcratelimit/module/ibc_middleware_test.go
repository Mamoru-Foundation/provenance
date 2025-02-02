package module_test

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"

	dbm "github.com/cosmos/cosmos-db"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	ibctesting "github.com/cosmos/ibc-go/v8/testing"

	"github.com/provenance-io/provenance/app"
	"github.com/provenance-io/provenance/internal/pioconfig"
	testutil "github.com/provenance-io/provenance/testutil/ibc"
	"github.com/provenance-io/provenance/x/ibcratelimit"
)

type MiddlewareTestSuite struct {
	suite.Suite

	*app.App
	Ctx sdk.Context

	coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	chainA *testutil.TestChain
	chainB *testutil.TestChain
	path   *ibctesting.Path
}

// Setup
func TestMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareTestSuite))
}

func SetupSimApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	pioconfig.SetProvenanceConfig(sdk.DefaultBondDenom, 0)
	db := dbm.NewMemDB()
	provenanceApp := app.New(log.NewNopLogger(), db, nil, true, map[int64]bool{}, app.DefaultNodeHome, 5, simtestutil.EmptyAppOptions{})
	genesis := provenanceApp.DefaultGenesis()
	return provenanceApp, genesis
}

func (s *MiddlewareTestSuite) SetupTest() {
	SkipIfWSL(s.T())
	ibctesting.DefaultTestingAppInit = SetupSimApp
	s.coordinator = ibctesting.NewCoordinator(s.T(), 2)

	s.chainA = &testutil.TestChain{
		TestChain: s.coordinator.GetChain(ibctesting.GetChainID(1)),
	}
	s.chainB = &testutil.TestChain{
		TestChain: s.coordinator.GetChain(ibctesting.GetChainID(2)),
	}
	s.path = NewTransferPath(s.chainA, s.chainB)
	s.coordinator.Setup(s.path)

	params, err := s.chainA.GetProvenanceApp().MintKeeper.Params.Get(s.chainA.GetContext())
	s.Require().NoError(err, "getting mint keeper params")
	params.InflationMax = sdkmath.LegacyNewDec(0)
	params.InflationRateChange = sdkmath.LegacyNewDec(1)
	params.InflationMin = sdkmath.LegacyNewDec(0)
	err = s.chainA.GetProvenanceApp().MintKeeper.Params.Set(s.chainA.GetContext(), params)
	s.Require().NoError(err, "setting mint keeper params for chainA")
	err = s.chainB.GetProvenanceApp().MintKeeper.Params.Set(s.chainB.GetContext(), params)
	s.Require().NoError(err, "setting mint keeper params for chainB")
}

// MessageFromAToB sends a message from chain A to chain B.
func (s *MiddlewareTestSuite) MessageFromAToB(denom string, amount sdkmath.Int) sdk.Msg {
	coin := sdk.NewCoin(denom, amount)
	port := s.path.EndpointA.ChannelConfig.PortID
	channel := s.path.EndpointA.ChannelID
	accountFrom := s.chainA.SenderAccount.GetAddress().String()
	accountTo := s.chainB.SenderAccount.GetAddress().String()
	timeoutHeight := clienttypes.NewHeight(0, 100)
	memo := ""
	return transfertypes.NewMsgTransfer(
		port,
		channel,
		coin,
		accountFrom,
		accountTo,
		timeoutHeight,
		0,
		memo,
	)
}

// MessageFromAToB sends a message from chain B to chain A.
func (s *MiddlewareTestSuite) MessageFromBToA(denom string, amount sdkmath.Int) sdk.Msg {
	coin := sdk.NewCoin(denom, amount)
	port := s.path.EndpointB.ChannelConfig.PortID
	channel := s.path.EndpointB.ChannelID
	accountFrom := s.chainB.SenderAccount.GetAddress().String()
	accountTo := s.chainA.SenderAccount.GetAddress().String()
	timeoutHeight := clienttypes.NewHeight(0, 100)
	memo := ""
	return transfertypes.NewMsgTransfer(
		port,
		channel,
		coin,
		accountFrom,
		accountTo,
		timeoutHeight,
		0,
		memo,
	)
}

// Tests that a receiver address longer than 4096 is not accepted
func (s *MiddlewareTestSuite) TestInvalidReceiver() {
	msg := transfertypes.NewMsgTransfer(
		s.path.EndpointB.ChannelConfig.PortID,
		s.path.EndpointB.ChannelID,
		sdk.NewInt64Coin(sdk.DefaultBondDenom, 1),
		s.chainB.SenderAccount.GetAddress().String(),
		strings.Repeat("x", 4097),
		clienttypes.NewHeight(0, 100),
		0,
		"",
	)
	_, ack, _ := s.FullSendBToA(msg)
	s.Assert().Contains(ack, "error",
		"acknowledgment is not an error")
	s.Assert().Contains(ack, fmt.Sprintf("ABCI code: %d", ibcratelimit.ErrBadMessage.ABCICode()),
		"acknowledgment error is not of the right type")
}

// FullSendBToA does the entire logic from sending a message from chain B to chain A.
func (s *MiddlewareTestSuite) FullSendBToA(msg sdk.Msg) (*sdk.Result, string, error) {
	sendResult, err := s.chainB.SendMsgsNoCheck(msg)
	s.Assert().NoError(err)

	packet, err := ibctesting.ParsePacketFromEvents(sendResult.Events)
	s.Assert().NoError(err)

	err = s.path.EndpointA.UpdateClient()
	s.Assert().NoError(err)

	res, err := s.path.EndpointA.RecvPacketWithResult(packet)
	s.Assert().NoError(err)

	ack, err := ibctesting.ParseAckFromEvents(res.GetEvents())
	s.Assert().NoError(err)

	err = s.path.EndpointA.UpdateClient()
	s.Assert().NoError(err)
	err = s.path.EndpointB.UpdateClient()
	s.Assert().NoError(err)

	return sendResult, string(ack), err
}

// FullSendAToB does the entire logic from sending a message from chain A to chain B.
func (s *MiddlewareTestSuite) FullSendAToB(msg sdk.Msg) (*sdk.Result, string, error) {
	sendResult, err := s.chainA.SendMsgsNoCheck(msg)
	if err != nil {
		return nil, "", err
	}

	packet, err := ibctesting.ParsePacketFromEvents(sendResult.Events)
	if err != nil {
		return nil, "", err
	}

	err = s.path.EndpointB.UpdateClient()
	if err != nil {
		return nil, "", err
	}

	res, err := s.path.EndpointB.RecvPacketWithResult(packet)
	if err != nil {
		return nil, "", err
	}

	ack, err := ibctesting.ParseAckFromEvents(res.GetEvents())
	if err != nil {
		return nil, "", err
	}

	err = s.path.EndpointA.UpdateClient()
	if err != nil {
		return nil, "", err
	}
	err = s.path.EndpointB.UpdateClient()
	if err != nil {
		return nil, "", err
	}

	return sendResult, string(ack), nil
}

// AssertSend checks that a receive on A from B was successful.
func (s *MiddlewareTestSuite) AssertReceive(success bool, msg sdk.Msg) (string, error) {
	_, ack, err := s.FullSendBToA(msg)
	if success {
		s.Assert().NoError(err)
		s.Assert().NotContains(ack, "error",
			"acknowledgment is an error")
	} else {
		s.Assert().Contains(ack, "error",
			"acknowledgment is not an error")
		s.Assert().Contains(ack, fmt.Sprintf("ABCI code: %d", ibcratelimit.ErrRateLimitExceeded.ABCICode()),
			"acknowledgment error is not of the right type")
	}
	return ack, err
}

// AssertSend checks that a send from A to B was successful.
func (s *MiddlewareTestSuite) AssertSend(success bool, msg sdk.Msg) (*sdk.Result, error) {
	r, _, err := s.FullSendAToB(msg)
	if success {
		s.Assert().NoError(err, "IBC send failed. Expected success. %s", err)
	} else {
		s.Assert().Error(err, "IBC send succeeded. Expected failure")
		s.ErrorContains(err, ibcratelimit.ErrRateLimitExceeded.Error(), "Bad error type")
	}
	return r, err
}

// BuildChannelQuota creates a quota message.
func (s *MiddlewareTestSuite) BuildChannelQuota(name, channel, denom string, duration, send_percentage, recv_percentage uint32) string {
	return fmt.Sprintf(`
          {"channel_id": "%s", "denom": "%s", "quotas": [{"name":"%s", "duration": %d, "send_recv":[%d, %d]}] }
    `, channel, denom, name, duration, send_percentage, recv_percentage)
}

// Tests

// Test that Sending IBC messages works when the middleware isn't configured
func (s *MiddlewareTestSuite) TestSendTransferNoContract() {
	one := sdkmath.NewInt(1)
	_, err := s.AssertSend(true, s.MessageFromAToB(sdk.DefaultBondDenom, one))
	s.Assert().NoError(err)
}

// Test that Receiving IBC messages works when the middleware isn't configured
func (s *MiddlewareTestSuite) TestReceiveTransferNoContract() {
	one := sdkmath.NewInt(1)
	_, err := s.AssertReceive(true, s.MessageFromBToA(sdk.DefaultBondDenom, one))
	s.Assert().NoError(err)
}

// initializeEscrow sets up the escrow on the chain.
func (s *MiddlewareTestSuite) initializeEscrow() (totalEscrow, expectedSed sdkmath.Int) {
	provenanceApp := s.chainA.GetProvenanceApp()
	supply := provenanceApp.BankKeeper.GetSupply(s.chainA.GetContext(), sdk.DefaultBondDenom)

	// Move some funds from chainA to chainB so that there is something in escrow
	// Each user has 10% of the supply, so we send most of the funds from one user to chainA
	transferAmount := supply.Amount.QuoRaw(20)

	// When sending, the amount we're sending goes into escrow before we enter the middleware and thus
	// it's used as part of the channel value in the rate limiting contract
	// To account for that, we subtract the amount we'll send first (2.5% of transferAmount) here
	sendAmount := transferAmount.QuoRaw(40)

	// Send from A to B
	_, _, err := s.FullSendAToB(s.MessageFromAToB(sdk.DefaultBondDenom, transferAmount.Sub(sendAmount)))
	s.Assert().NoError(err)
	// Send from A to B
	_, _, err = s.FullSendBToA(s.MessageFromBToA(sdk.DefaultBondDenom, transferAmount.Sub(sendAmount)))
	s.Assert().NoError(err)

	return transferAmount, sendAmount
}

func (s *MiddlewareTestSuite) fullSendTest(native bool) map[string]string {
	quotaPercentage := 5
	s.initializeEscrow()
	// Get the denom and amount to send
	denom := sdk.DefaultBondDenom
	channel := "channel-0"
	if !native {
		denomTrace := transfertypes.ParseDenomTrace(transfertypes.GetPrefixedDenom("transfer", "channel-0", denom))
		fmt.Println(denomTrace)
		denom = denomTrace.IBCDenom()
	}

	provenanceApp := s.chainA.GetProvenanceApp()

	// This is the first one. Inside the tests. It works as expected.
	channelValue := CalculateChannelValue(s.chainA.GetContext(), denom, provenanceApp.BankKeeper)

	// The amount to be sent is send 2.5% (quota is 5%)
	quota := channelValue.QuoRaw(int64(100 / quotaPercentage))
	sendAmount := quota.QuoRaw(2)

	fmt.Printf("Testing send rate limiting for denom=%s, channelValue=%s, quota=%s, sendAmount=%s\n", denom, channelValue, quota, sendAmount)

	// Setup contract
	s.chainA.StoreContractRateLimiterDirect(&s.Suite)
	quotas := s.BuildChannelQuota("weekly", channel, denom, 604800, 5, 5)
	fmt.Println(quotas)
	initMsg := CreateRateLimiterInitMessage(s.chainA, quotas)
	addr := s.chainA.InstantiateContract(&s.Suite, initMsg, 1)
	s.chainA.RegisterRateLimiterContract(&s.Suite, addr)

	// send 2.5% (quota is 5%)
	fmt.Printf("Sending %s from A to B. Represented in chain A as wrapped? %v\n", denom, !native)
	_, err := s.AssertSend(true, s.MessageFromAToB(denom, sendAmount))
	s.Assert().NoError(err)

	// send 2.5% (quota is 5%)
	fmt.Println("trying to send ", sendAmount)
	r, _ := s.AssertSend(true, s.MessageFromAToB(denom, sendAmount))

	// Calculate remaining allowance in the quota
	attrs := ExtractAttributes(FindEvent(r.GetEvents(), "wasm"))

	used, ok := sdkmath.NewIntFromString(attrs["weekly_used_out"])
	s.Assert().True(ok)

	s.Assert().Equal(used, sendAmount.MulRaw(2))

	// Sending above the quota should fail. We use 2 instead of 1 here to avoid rounding issues
	_, err = s.AssertSend(false, s.MessageFromAToB(denom, sdkmath.NewInt(2)))
	s.Assert().Error(err)
	return attrs
}

// Test rate limiting on sends
func (s *MiddlewareTestSuite) TestSendTransferWithRateLimitingNative() {
	// Sends denom=stake from A->B. Rate limit receives "stake" in the packet. Nothing to do in the contract
	s.fullSendTest(true)
}

// Test rate limiting on sends
func (s *MiddlewareTestSuite) TestSendTransferWithRateLimitingNonNative() {
	// Sends denom=ibc/C053D637CCA2A2BA030E2C5EE1B28A16F71CCB0E45E8BE52766DC1B241B77878 from A->B.
	// Rate limit receives "transfer/channel-0/stake" in the packet (because transfer.relay.SendTransfer is called before the middleware)
	// and should hash it before calculating the value
	s.fullSendTest(false)
}

// Test rate limits are reset when the specified time period has passed
func (s *MiddlewareTestSuite) TestSendTransferReset() {
	// Same test as above, but the quotas get reset after time passes
	attrs := s.fullSendTest(true)
	parts := strings.Split(attrs["weekly_period_end"], ".") // Splitting timestamp into secs and nanos
	secs, err := strconv.ParseInt(parts[0], 10, 64)
	s.Assert().NoError(err)
	nanos, err := strconv.ParseInt(parts[1], 10, 64)
	s.Assert().NoError(err)
	resetTime := time.Unix(secs, nanos)

	// Move chainA forward one block
	s.chainA.NextBlock()
	err = s.chainA.SenderAccount.SetSequence(s.chainA.SenderAccount.GetSequence() + 1)
	s.Assert().NoError(err)

	// Reset time + one second
	oneSecAfterReset := resetTime.Add(time.Second)
	s.coordinator.IncrementTimeBy(oneSecAfterReset.Sub(s.coordinator.CurrentTime))

	// Sending should succeed again
	_, err = s.AssertSend(true, s.MessageFromAToB(sdk.DefaultBondDenom, sdkmath.NewInt(1)))
	s.Assert().NoError(err)
}

// Test rate limiting on receives
func (s *MiddlewareTestSuite) fullRecvTest(native bool) {
	quotaPercentage := 4
	s.initializeEscrow()
	// Get the denom and amount to send
	sendDenom := sdk.DefaultBondDenom
	localDenom := sdk.DefaultBondDenom
	channel := "channel-0"
	if native {
		denomTrace := transfertypes.ParseDenomTrace(transfertypes.GetPrefixedDenom("transfer", "channel-0", localDenom))
		localDenom = denomTrace.IBCDenom()
	} else {
		denomTrace := transfertypes.ParseDenomTrace(transfertypes.GetPrefixedDenom("transfer", "channel-0", sendDenom))
		sendDenom = denomTrace.IBCDenom()
	}

	provenanceApp := s.chainA.GetProvenanceApp()

	channelValue := CalculateChannelValue(s.chainA.GetContext(), localDenom, provenanceApp.BankKeeper)

	// The amount to be sent is 2% (quota is 4%)
	quota := channelValue.QuoRaw(int64(100 / quotaPercentage))
	sendAmount := quota.QuoRaw(2)

	fmt.Printf("Testing recv rate limiting for denom=%s, channelValue=%s, quota=%s, sendAmount=%s\n", localDenom, channelValue, quota, sendAmount)

	// Setup contract
	s.chainA.StoreContractRateLimiterDirect(&s.Suite)
	quotas := s.BuildChannelQuota("weekly", channel, localDenom, 604800, 4, 4)
	initMsg := CreateRateLimiterInitMessage(s.chainA, quotas)
	addr := s.chainA.InstantiateContract(&s.Suite, initMsg, 1)
	s.chainA.RegisterRateLimiterContract(&s.Suite, addr)

	// receive 2.5% (quota is 5%)
	fmt.Printf("Sending %s from B to A. Represented in chain A as wrapped? %v\n", sendDenom, native)
	_, err := s.AssertReceive(true, s.MessageFromBToA(sendDenom, sendAmount))
	s.Assert().NoError(err)

	// receive 2.5% (quota is 5%)
	_, err = s.AssertReceive(true, s.MessageFromBToA(sendDenom, sendAmount))
	s.Assert().NoError(err)

	// Sending above the quota should fail. We send 2 instead of 1 to account for rounding errors
	_, err = s.AssertReceive(false, s.MessageFromBToA(sendDenom, sdkmath.NewInt(2)))
	s.Assert().NoError(err)
}

func (s *MiddlewareTestSuite) TestRecvTransferWithRateLimitingNative() {
	// Sends denom=stake from B->A.
	// Rate limit receives "stake" in the packet and should wrap it before calculating the value
	// types.ReceiverChainIsSource(packet.GetSourcePort(), packet.GetSourceChannel(), data.Denom) should return false => Wrap the token
	s.fullRecvTest(true)
}

func (s *MiddlewareTestSuite) TestRecvTransferWithRateLimitingNonNative() {
	// Sends denom=ibc/C053D637CCA2A2BA030E2C5EE1B28A16F71CCB0E45E8BE52766DC1B241B77878 from B->A.
	// Rate limit receives "transfer/channel-0/stake" in the packet and should turn it into "stake"
	// types.ReceiverChainIsSource(packet.GetSourcePort(), packet.GetSourceChannel(), data.Denom) should return true => unprefix. If unprefixed is not local, hash.
	s.fullRecvTest(false)
}

// Test no rate limiting occurs when the contract is set, but not quotas are condifured for the path
func (s *MiddlewareTestSuite) TestSendTransferNoQuota() {
	// Setup contract
	s.chainA.StoreContractRateLimiterDirect(&s.Suite)
	initMsg := CreateRateLimiterInitMessage(s.chainA, "")
	addr := s.chainA.InstantiateContract(&s.Suite, initMsg, 1)
	s.chainA.RegisterRateLimiterContract(&s.Suite, addr)

	// send 1 token.
	// If the contract doesn't have a quota for the current channel, all transfers are allowed
	_, err := s.AssertSend(true, s.MessageFromAToB(sdk.DefaultBondDenom, sdkmath.NewInt(1)))
	s.Assert().NoError(err)
}

// Test rate limits are reverted if a "send" fails
func (s *MiddlewareTestSuite) TestFailedSendTransfer() {
	s.initializeEscrow()
	// Setup contract
	s.chainA.StoreContractRateLimiterDirect(&s.Suite)
	quotas := s.BuildChannelQuota("weekly", "channel-0", sdk.DefaultBondDenom, 604800, 1, 1)
	initMsg := CreateRateLimiterInitMessage(s.chainA, quotas)
	addr := s.chainA.InstantiateContract(&s.Suite, initMsg, 1)
	s.chainA.RegisterRateLimiterContract(&s.Suite, addr)

	// Get the escrowed amount
	provenanceApp := s.chainA.GetProvenanceApp()
	// ToDo: This is what we eventually want here, but using the full supply temporarily for performance reasons. See calculateChannelValue
	// escrowAddress := transfertypes.GetEscrowAddress("transfer", "channel-0")
	// escrowed := provenanceApp.BankKeeper.GetBalance(s.chainA.GetContext(), escrowAddress, sdk.DefaultBondDenom)
	escrowed := provenanceApp.BankKeeper.GetSupply(s.chainA.GetContext(), sdk.DefaultBondDenom)
	quota := escrowed.Amount.QuoRaw(100) // 1% of the escrowed amount

	// Use the whole quota
	coins := sdk.NewCoin(sdk.DefaultBondDenom, quota)
	port := s.path.EndpointA.ChannelConfig.PortID
	channel := s.path.EndpointA.ChannelID
	accountFrom := s.chainA.SenderAccount.GetAddress().String()
	timeoutHeight := clienttypes.NewHeight(0, 100)
	memo := ""
	msg := transfertypes.NewMsgTransfer(port, channel, coins, accountFrom, "INVALID", timeoutHeight, 0, memo)

	// Sending the message manually because AssertSend updates both clients. We need to update the clients manually
	// for this test so that the failure to receive on chain B happens after the second packet is sent from chain A.
	// That way we validate that chain A is blocking as expected, but the flow is reverted after the receive failure is
	// acknowledged on chain A
	res, err := s.chainA.SendMsgsNoCheck(msg)
	s.Assert().NoError(err)

	// Sending again fails as the quota is filled
	_, err = s.AssertSend(false, s.MessageFromAToB(sdk.DefaultBondDenom, quota))
	s.Assert().Error(err)

	// Move forward one block
	s.chainA.NextBlock()
	err = s.chainA.SenderAccount.SetSequence(s.chainA.SenderAccount.GetSequence() + 1)
	s.Assert().NoError(err)
	s.chainA.Coordinator.IncrementTime()

	// Update both clients
	err = s.path.EndpointA.UpdateClient()
	s.Assert().NoError(err)
	err = s.path.EndpointB.UpdateClient()
	s.Assert().NoError(err)

	// Execute the acknowledgement from chain B in chain A

	// extract the sent packet
	packet, err := ibctesting.ParsePacketFromEvents(res.Events)
	s.Assert().NoError(err)

	// recv in chain b
	txRes, err := s.path.EndpointB.RecvPacketWithResult(packet)
	s.Assert().NoError(err)

	// get the ack from the chain b's response
	ack, err := ibctesting.ParseAckFromEvents(txRes.Events)
	s.Assert().NoError(err)

	// manually relay it to chain a
	err = s.path.EndpointA.AcknowledgePacket(packet, ack)
	s.Assert().NoError(err)

	// We should be able to send again because the packet that exceeded the quota failed and has been reverted
	_, err = s.AssertSend(true, s.MessageFromAToB(sdk.DefaultBondDenom, sdkmath.NewInt(1)))
	s.Assert().NoError(err)
}

func (s *MiddlewareTestSuite) TestUnsetRateLimitingContract() {
	// Setup contract
	s.chainA.StoreContractRateLimiterDirect(&s.Suite)
	msg := CreateRateLimiterInitMessage(s.chainA, "")
	addr := s.chainA.InstantiateContract(&s.Suite, msg, 1)
	s.chainA.RegisterRateLimiterContract(&s.Suite, addr)

	// Unset the contract param
	s.chainA.RegisterRateLimiterContract(&s.Suite, []byte(""))
	contractAddress := s.chainA.GetProvenanceApp().RateLimitingKeeper.GetContractAddress(s.chainA.GetContext())
	s.Assert().Equal("", contractAddress, "should unregister contract")

}

// FindEvent finds an event with a matching name.
func FindEvent(events []sdk.Event, name string) sdk.Event {
	index := slices.IndexFunc(events, func(e sdk.Event) bool { return e.Type == name })
	if index == -1 {
		return sdk.Event{}
	}
	return events[index]
}

// ExtractAttributes returns the event's attributes in a map.
func ExtractAttributes(event sdk.Event) map[string]string {
	attrs := make(map[string]string)
	if event.Attributes == nil {
		return attrs
	}
	for _, a := range event.Attributes {
		attrs[a.Key] = a.Value
	}
	return attrs
}

// CreateRateLimiterInitMessage creates a contract init message for the rate limiter using the supplied quotas.
func CreateRateLimiterInitMessage(chain *testutil.TestChain, quotas string) string {
	provenanceApp := chain.GetProvenanceApp()
	transferModule := provenanceApp.AccountKeeper.GetModuleAddress(transfertypes.ModuleName)
	govModule := provenanceApp.AccountKeeper.GetModuleAddress(govtypes.ModuleName)

	initMsg := fmt.Sprintf(`{
           "gov_module":  "%s",
           "ibc_module":"%s",
           "paths": [%s]
        }`,
		govModule, transferModule, quotas)
	return initMsg
}

// CalculateChannelValue returns the total number of denom on a channel.
func CalculateChannelValue(ctx sdk.Context, denom string, bankKeeper bankkeeper.Keeper) sdkmath.Int {
	return bankKeeper.GetSupply(ctx, denom).Amount

	// ToDo: The commented-out code bellow is what we want to happen, but we're temporarily
	//  using the whole supply for efficiency until there's a solution for
	//  https://github.com/cosmos/ibc-go/issues/2664

	// For non-native (ibc) tokens, return the supply if the token in osmosis
	//if strings.HasPrefix(denom, "ibc/") {
	//	return bankKeeper.GetSupplyWithOffset(ctx, denom).Amount
	//}
	//
	// For native tokens, obtain the balance held in escrow for all potential channels
	//channels := channelKeeper.GetAllChannels(ctx)
	//balance := osmomath.NewInt(0)
	//for _, channel := range channels {
	//	escrowAddress := transfertypes.GetEscrowAddress("transfer", channel.ChannelId)
	//	balance = balance.Add(bankKeeper.GetBalance(ctx, escrowAddress, denom).Amount)
	//
	//}
	//return balance
}

// NewTransferPath creates a new ibc transfer path for testing.
func NewTransferPath(chainA, chainB *testutil.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainA.TestChain, chainB.TestChain)
	path.EndpointA.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointB.ChannelConfig.PortID = ibctesting.TransferPort
	path.EndpointA.ChannelConfig.Version = transfertypes.Version
	path.EndpointB.ChannelConfig.Version = transfertypes.Version
	return path
}

// SkipIfWSL skips the test if it being ran on WSL.
func SkipIfWSL(t *testing.T) {
	t.Helper()
	skip := os.Getenv("SKIP_WASM_WSL_TESTS")
	if skip == "true" {
		t.Skip("Skipping Wasm tests")
	}
}

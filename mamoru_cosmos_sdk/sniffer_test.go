package mamoru_cosmos_sdk

import (
	log2 "cosmossdk.io/log"
	"os"
	"testing"

	abci "github.com/cometbft/cometbft/abci/types"
	tmprototypes "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gotest.tools/v3/assert"
)

// TestNewSniffer tests the NewSniffer function
func TestNewSniffer(t *testing.T) {
	snifferTest := NewSniffer(log2.NewLogger(os.Stdout, log2.TraceOption(true), log2.ColorOption(false)))
	if snifferTest == nil {
		t.Error("NewSniffer returned nil")
	}
}

// TestIsSnifferEnable tests the isSnifferEnable method
func TestIsSnifferEnable(t *testing.T) {

	// Set environment variable for testing
	t.Setenv("MAMORU_SNIFFER_ENABLE", "true")
	logger := log2.NewLogger(os.Stdout, log2.TraceOption(true), log2.ColorOption(false))
	_ = NewSniffer(logger)
	if !isSnifferEnabled() {
		t.Error("Expected sniffer to be enabled")
	}

	// Test with invalid value
	t.Setenv("MAMORU_SNIFFER_ENABLE", "not_a_bool")
	if isSnifferEnabled() {
		t.Error("Expected sniffer to be disabled with invalid env value")
	}
}

// smoke test for the sniffer
func TestSnifferSmoke(t *testing.T) {

	t.Setenv("MAMORU_SNIFFER_ENABLE", "true")
	t.Setenv("MAMORU_CHAIN_TYPE", "ETH_TESTNET")
	t.Setenv("MAMORU_CHAIN_ID", "validationchain")
	t.Setenv("MAMORU_STATISTICS_SEND_INTERVAL_SECS", "1")
	t.Setenv("MAMORU_ENDPOINT", "http://localhost:9090")
	t.Setenv("MAMORU_PRIVATE_KEY", "6Hi8mqAFp14m3pySNYDjXhUysZok0X6jaMWvwZGdd8=")
	//InitConnectFunc(func() (*cosmos.SnifferCosmos, error) {
	//	return nil, nil
	//})
	logger := log2.NewLogger(os.Stdout, log2.TraceOption(true), log2.ColorOption(false))
	sniffer := NewSniffer(logger)
	if sniffer == nil {
		t.Error("NewSniffer returned nil")
	}
	header := tmprototypes.Header{}
	ischeck := true
	ctx := sdk.NewContext(nil, header, ischeck, logger)

	streamingService := NewStreamingService(logger, sniffer)
	regBB := abci.RequestFinalizeBlock{}
	resBB := abci.ResponseFinalizeBlock{}
	err := streamingService.ListenFinalizeBlock(ctx, regBB, resBB)
	assert.NilError(t, err)

	resC := abci.ResponseCommit{}
	err = streamingService.ListenCommit(ctx, resC, nil)
	assert.NilError(t, err)
}

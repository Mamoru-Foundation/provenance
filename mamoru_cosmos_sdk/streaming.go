package mamoru_cosmos_sdk

import (
	"context"
	"cosmossdk.io/log"
	"encoding/hex"
	"strconv"
	"strings"

	"cosmossdk.io/store/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/bytes"
	types2 "github.com/cometbft/cometbft/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer/cosmos"
)

var _ types.ABCIListener = (*StreamingService)(nil)

type StreamingService struct {
	logger log.Logger

	blockMetadata      types.BlockMetadata
	currentBlockNumber int64

	sniffer *Sniffer
}

func NewStreamingService(logger log.Logger, sniffer *Sniffer) *StreamingService {
	logger.Info("Mamoru StreamingService start")

	return &StreamingService{
		sniffer: sniffer,
		logger:  logger,
	}
}

func (ss *StreamingService) ListenFinalizeBlock(ctx context.Context, req abci.RequestFinalizeBlock, res abci.ResponseFinalizeBlock) error {
	ss.blockMetadata = types.BlockMetadata{}
	ss.blockMetadata.RequestFinalizeBlock = &req
	ss.blockMetadata.ResponseFinalizeBlock = &res
	return nil
}

func (ss *StreamingService) ListenCommit(ctx context.Context, res abci.ResponseCommit, changeSet []*types.StoreKVPair) error {
	if ss.sniffer == nil || !ss.sniffer.CheckRequirements() {
		return nil
	}

	ss.blockMetadata.ResponseCommit = &res
	ss.logger.Info("Mamoru ListenCommit", "height", ss.currentBlockNumber)

	var eventCount uint64 = 0
	var txCount uint64 = 0
	var callTracesCount uint64 = 0
	builder := cosmos.NewCosmosCtxBuilder()

	blockHeight := uint64(ss.blockMetadata.RequestFinalizeBlock.Height)
	block := cosmos.Block{
		Seq:    blockHeight,
		Height: ss.blockMetadata.RequestFinalizeBlock.Height,
		Hash:   hex.EncodeToString(ss.blockMetadata.RequestFinalizeBlock.Hash),
		//VersionBlock:                  ss.blockMetadata.RequestBeginBlock.Header.Version.Block,
		//VersionApp:                    ss.blockMetadata.RequestBeginBlock.Header.Version.App,
		//ChainId:                       ss.blockMetadata.RequestBeginBlock.Header.ChainID,
		Time:            ss.blockMetadata.RequestFinalizeBlock.Time.Unix(),
		LastBlockIdHash: hex.EncodeToString(ss.blockMetadata.RequestFinalizeBlock.Hash),
		//LastBlockIdPartSetHeaderTotal: ss.blockMetadata.RequestBeginBlock.Header.LastBlockId.PartSetHeader.Total,
		//LastBlockIdPartSetHeaderHash:  hex.EncodeToString(ss.blockMetadata.RequestBeginBlock.Header.LastBlockId.PartSetHeader.Hash),
		//LastCommitHash:                hex.EncodeToString(ss.blockMetadata.RequestBeginBlock.Header.LastCommitHash),
		//DataHash:                      hex.EncodeToString(ss.blockMetadata.RequestBeginBlock.Header.DataHash),
		//ValidatorsHash:                hex.EncodeToString(ss.blockMetadata.RequestBeginBlock.Header.ValidatorsHash),
		NextValidatorsHash: hex.EncodeToString(ss.blockMetadata.RequestFinalizeBlock.NextValidatorsHash),
		//		ConsensusHash:                 hex.EncodeToString(ss.blockMetadata.RequestBeginBlock.Header.ConsensusHash),
		AppHash: hex.EncodeToString(ss.blockMetadata.ResponseFinalizeBlock.AppHash),
		//LastResultsHash:               hex.EncodeToString(ss.blockMetadata.RequestBeginBlock.Header.LastResultsHash),
		//EvidenceHash:                  hex.EncodeToString(ss.blockMetadata.RequestBeginBlock.Header.EvidenceHash),
		ProposerAddress: hex.EncodeToString(ss.blockMetadata.RequestFinalizeBlock.ProposerAddress),
		//LastCommitInfoRound:           ss.blockMetadata.RequestBeginBlock.LastCommitInfo.Round,
	}

	if ss.blockMetadata.ResponseFinalizeBlock.ConsensusParamUpdates != nil {
		block.ConsensusParamUpdatesBlockMaxBytes = ss.blockMetadata.ResponseFinalizeBlock.ConsensusParamUpdates.Block.MaxBytes
		block.ConsensusParamUpdatesBlockMaxGas = ss.blockMetadata.ResponseFinalizeBlock.ConsensusParamUpdates.Block.MaxGas
		block.ConsensusParamUpdatesEvidenceMaxAgeNumBlocks = ss.blockMetadata.ResponseFinalizeBlock.ConsensusParamUpdates.Evidence.MaxAgeNumBlocks
		block.ConsensusParamUpdatesEvidenceMaxAgeDuration = ss.blockMetadata.ResponseFinalizeBlock.ConsensusParamUpdates.Evidence.MaxAgeDuration.Milliseconds()
		block.ConsensusParamUpdatesEvidenceMaxBytes = ss.blockMetadata.ResponseFinalizeBlock.ConsensusParamUpdates.Evidence.MaxBytes
		block.ConsensusParamUpdatesValidatorPubKeyTypes = strings.Join(ss.blockMetadata.ResponseFinalizeBlock.ConsensusParamUpdates.Validator.PubKeyTypes[:], ",")
		block.ConsensusParamUpdatesVersionApp = ss.blockMetadata.ResponseFinalizeBlock.ConsensusParamUpdates.Version.GetApp()
	}

	builder.SetBlock(block)

	for _, beginBlock := range ss.blockMetadata.ResponseFinalizeBlock.Events {
		eventCount++
		builder.AppendEvents([]cosmos.Event{
			{
				Seq:       blockHeight,
				EventType: beginBlock.Type,
			},
		})
		for _, attribute := range beginBlock.Attributes {
			builder.AppendEventAttributes([]cosmos.EventAttribute{
				{
					Seq:      blockHeight,
					EventSeq: blockHeight,
					Key:      string(attribute.Key),
					Value:    string(attribute.Value),
					Index:    attribute.Index,
				},
			})
		}
	}

	for _, validatorUpdate := range ss.blockMetadata.ResponseFinalizeBlock.ValidatorUpdates {
		builder.AppendValidatorUpdates([]cosmos.ValidatorUpdate{
			{
				Seq:    blockHeight,
				PubKey: validatorUpdate.PubKey.GetEd25519(),
				Power:  validatorUpdate.Power,
			},
		})
	}

	for _, voteInfo := range ss.blockMetadata.RequestFinalizeBlock.DecidedLastCommit.Votes {
		builder.AppendVoteInfos([]cosmos.VoteInfo{
			{
				Seq:              blockHeight,
				BlockSeq:         blockHeight,
				ValidatorAddress: sdktypes.ValAddress(voteInfo.Validator.Address).String(),
				ValidatorPower:   voteInfo.Validator.Power,
				//SignedLastBlock:  voteInfo.SignedLastBlock,
			},
		})
	}

	for _, misbehavior := range ss.blockMetadata.RequestFinalizeBlock.Misbehavior {
		builder.AppendMisbehaviors([]cosmos.Misbehavior{
			{
				Seq:              blockHeight,
				BlockSeq:         blockHeight,
				Typ:              misbehavior.Type.String(),
				ValidatorPower:   misbehavior.Validator.Power,
				ValidatorAddress: sdktypes.ValAddress(misbehavior.Validator.Address).String(),
				Height:           misbehavior.Height,
				Time:             misbehavior.Time.Unix(),
				TotalVotingPower: misbehavior.TotalVotingPower,
			},
		})
	}

	for txIndex, tx := range ss.blockMetadata.ResponseFinalizeBlock.TxResults {
		txHash := bytes.HexBytes(types2.Tx(tx.Data).Hash()).String()
		builder.AppendTxs([]cosmos.Transaction{
			{
				Seq:       blockHeight,
				Tx:        []byte(txHash),
				TxHash:    txHash,
				TxIndex:   uint32(txIndex),
				Code:      tx.Code,
				Data:      tx.Data,
				Log:       tx.Log,
				Info:      tx.Info,
				GasWanted: tx.GasWanted,
				GasUsed:   tx.GasUsed,
				Codespace: tx.Codespace,
			},
		})

		for _, event := range tx.Events {
			eventCount++
			builder.AppendEvents([]cosmos.Event{
				{
					Seq:       blockHeight,
					EventType: event.Type,
				},
			})

			for _, attribute := range event.Attributes {
				builder.AppendEventAttributes([]cosmos.EventAttribute{
					{
						Seq:      blockHeight,
						EventSeq: blockHeight,
						Key:      string(attribute.Key),
						Value:    string(attribute.Value),
						Index:    attribute.Index,
					},
				})
			}
		}

		txCount++
	}

	for _, event := range ss.blockMetadata.ResponseFinalizeBlock.Events {
		eventCount++
		builder.AppendEvents([]cosmos.Event{
			{
				Seq:       blockHeight,
				EventType: event.Type,
			},
		})
		for _, attribute := range event.Attributes {
			builder.AppendEventAttributes([]cosmos.EventAttribute{
				{
					Seq:      blockHeight,
					EventSeq: blockHeight,
					Key:      string(attribute.Key),
					Value:    string(attribute.Value),
					Index:    attribute.Index,
				},
			})
		}
	}

	builder.SetBlockData(strconv.FormatUint(blockHeight, 10), hex.EncodeToString(ss.blockMetadata.RequestFinalizeBlock.Hash))

	statTxs := txCount
	statEvn := eventCount

	builder.SetStatistics(uint64(1), statTxs, statEvn, callTracesCount)

	cosmosCtx := builder.Finish()

	ss.logger.Info("Mamoru Send", "height", ss.currentBlockNumber, "txs", statTxs, "events", statEvn, "callTraces", callTracesCount)

	if client := ss.sniffer.Client(); client != nil {
		client.ObserveCosmosData(cosmosCtx)
	}

	return nil
}

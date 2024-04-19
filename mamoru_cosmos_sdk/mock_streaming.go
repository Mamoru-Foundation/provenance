package mamoru_cosmos_sdk

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/types"

	tmabci "github.com/tendermint/tendermint/abci/types"
	tmlog "github.com/tendermint/tendermint/libs/log"
	//"github.com/Mamoru-Foundation/mamoru-sniffer-go/mamoru_sniffer"
)

var _ types.ABCIListener = (*MockStreamingService)(nil)

// MockStreamingService mock streaming service
type MockStreamingService struct {
	logger             tmlog.Logger
	currentBlockNumber int64
}

func NewMockStreamingService(logger tmlog.Logger) *MockStreamingService {
	logger.Info("Mamoru MockStreamingService start")

	return &MockStreamingService{
		logger: logger,
	}
}

func (ss *MockStreamingService) ListenBeginBlock(ctx context.Context, req tmabci.RequestBeginBlock, res tmabci.ResponseBeginBlock) error {
	ss.currentBlockNumber = req.Header.Height
	ss.logger.Info("Mamoru Mock ListenBeginBlock", "height", ss.currentBlockNumber)

	return nil
}

func (ss *MockStreamingService) ListenDeliverTx(ctx context.Context, req tmabci.RequestDeliverTx, res tmabci.ResponseDeliverTx) error {
	ss.logger.Info("Mamoru Mock ListenDeliverTx", "height", ss.currentBlockNumber)

	return nil
}

func (ss *MockStreamingService) ListenEndBlock(ctx context.Context, req tmabci.RequestEndBlock, res tmabci.ResponseEndBlock) error {
	ss.logger.Info("Mamoru Mock ListenEndBlock", "height", ss.currentBlockNumber)

	return nil
}

func (ss *MockStreamingService) ListenCommit(ctx context.Context, res tmabci.ResponseCommit, changeSet []*types.StoreKVPair) error {
	ss.logger.Info("Mamoru Mock ListenCommit", "height", ss.currentBlockNumber)

	return nil
}

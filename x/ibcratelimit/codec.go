package ibcratelimit

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/gogoproto/proto"
)

// RegisterInterfaces registers implementations for the tx messages
func RegisterInterfaces(registry types.InterfaceRegistry) {
	messages := make([]proto.Message, len(AllRequestMsgs))
	for i, msg := range AllRequestMsgs {
		messages[i] = msg
	}
	registry.RegisterImplementations((*sdk.Msg)(nil), messages...)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

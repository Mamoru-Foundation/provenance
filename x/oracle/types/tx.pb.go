// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/oracle/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_CosmWasm_wasmd_x_wasm_types "github.com/CosmWasm/wasmd/x/wasm/types"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgSendQueryOracleRequest queries an oracle on another chain
type MsgSendQueryOracleRequest struct {
	// Query contains the query data passed to the oracle.
	Query github_com_CosmWasm_wasmd_x_wasm_types.RawContractMessage `protobuf:"bytes,1,opt,name=query,proto3,casttype=github.com/CosmWasm/wasmd/x/wasm/types.RawContractMessage" json:"query,omitempty"`
	// Channel is the channel to the oracle.
	Channel string `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	// The signing authority for the request
	Authority string `protobuf:"bytes,4,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (m *MsgSendQueryOracleRequest) Reset()         { *m = MsgSendQueryOracleRequest{} }
func (m *MsgSendQueryOracleRequest) String() string { return proto.CompactTextString(m) }
func (*MsgSendQueryOracleRequest) ProtoMessage()    {}
func (*MsgSendQueryOracleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66a39dda41c6a784, []int{0}
}
func (m *MsgSendQueryOracleRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendQueryOracleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendQueryOracleRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendQueryOracleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendQueryOracleRequest.Merge(m, src)
}
func (m *MsgSendQueryOracleRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendQueryOracleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendQueryOracleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendQueryOracleRequest proto.InternalMessageInfo

func (m *MsgSendQueryOracleRequest) GetQuery() github_com_CosmWasm_wasmd_x_wasm_types.RawContractMessage {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *MsgSendQueryOracleRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *MsgSendQueryOracleRequest) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

// MsgSendQueryOracleResponse contains the id of the oracle query.
type MsgSendQueryOracleResponse struct {
	// The sequence number that uniquely identifies the query.
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *MsgSendQueryOracleResponse) Reset()         { *m = MsgSendQueryOracleResponse{} }
func (m *MsgSendQueryOracleResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendQueryOracleResponse) ProtoMessage()    {}
func (*MsgSendQueryOracleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66a39dda41c6a784, []int{1}
}
func (m *MsgSendQueryOracleResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendQueryOracleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendQueryOracleResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendQueryOracleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendQueryOracleResponse.Merge(m, src)
}
func (m *MsgSendQueryOracleResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendQueryOracleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendQueryOracleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendQueryOracleResponse proto.InternalMessageInfo

func (m *MsgSendQueryOracleResponse) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

// MsgUpdateOracleRequest is the request type for updating an oracle's contract address
type MsgUpdateOracleRequest struct {
	// The address of the oracle's contract
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// The signing authorities for the request
	Authority string `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (m *MsgUpdateOracleRequest) Reset()         { *m = MsgUpdateOracleRequest{} }
func (m *MsgUpdateOracleRequest) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateOracleRequest) ProtoMessage()    {}
func (*MsgUpdateOracleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_66a39dda41c6a784, []int{2}
}
func (m *MsgUpdateOracleRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateOracleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateOracleRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateOracleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateOracleRequest.Merge(m, src)
}
func (m *MsgUpdateOracleRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateOracleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateOracleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateOracleRequest proto.InternalMessageInfo

func (m *MsgUpdateOracleRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgUpdateOracleRequest) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

// MsgUpdateOracleResponse is the response type for updating the oracle.
type MsgUpdateOracleResponse struct {
}

func (m *MsgUpdateOracleResponse) Reset()         { *m = MsgUpdateOracleResponse{} }
func (m *MsgUpdateOracleResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateOracleResponse) ProtoMessage()    {}
func (*MsgUpdateOracleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_66a39dda41c6a784, []int{3}
}
func (m *MsgUpdateOracleResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateOracleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateOracleResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateOracleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateOracleResponse.Merge(m, src)
}
func (m *MsgUpdateOracleResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateOracleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateOracleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateOracleResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSendQueryOracleRequest)(nil), "provenance.oracle.v1.MsgSendQueryOracleRequest")
	proto.RegisterType((*MsgSendQueryOracleResponse)(nil), "provenance.oracle.v1.MsgSendQueryOracleResponse")
	proto.RegisterType((*MsgUpdateOracleRequest)(nil), "provenance.oracle.v1.MsgUpdateOracleRequest")
	proto.RegisterType((*MsgUpdateOracleResponse)(nil), "provenance.oracle.v1.MsgUpdateOracleResponse")
}

func init() { proto.RegisterFile("provenance/oracle/v1/tx.proto", fileDescriptor_66a39dda41c6a784) }

var fileDescriptor_66a39dda41c6a784 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x6b, 0x36, 0xd8, 0x66, 0x4d, 0x42, 0xb2, 0x2a, 0x96, 0x46, 0x22, 0x9d, 0x72, 0xda,
	0x81, 0xc6, 0xac, 0x48, 0x08, 0x90, 0x38, 0xd0, 0x9d, 0x23, 0x20, 0x15, 0x42, 0xe2, 0x82, 0xbc,
	0xc4, 0x72, 0x23, 0x16, 0x3b, 0xf3, 0x73, 0xba, 0xf6, 0x4b, 0x20, 0x8e, 0x1c, 0xf7, 0x21, 0xf8,
	0x04, 0x9c, 0x38, 0x56, 0x9c, 0x38, 0x21, 0xd4, 0x5e, 0xfa, 0x19, 0x38, 0xa1, 0xc4, 0x2d, 0x6d,
	0x21, 0xa0, 0x9e, 0x92, 0xf7, 0xde, 0xff, 0x3d, 0xfb, 0x67, 0xfd, 0x1f, 0xbe, 0x9b, 0x6b, 0x35,
	0xe4, 0x92, 0xc9, 0x98, 0x53, 0xa5, 0x59, 0x7c, 0xc1, 0xe9, 0xf0, 0x94, 0x9a, 0x51, 0x90, 0x6b,
	0x65, 0x14, 0x69, 0xae, 0xca, 0x81, 0x2d, 0x07, 0xc3, 0x53, 0xb7, 0x15, 0x2b, 0xc8, 0x14, 0xbc,
	0xad, 0x34, 0xd4, 0x06, 0xb6, 0xc1, 0x6d, 0x0a, 0x25, 0x94, 0xcd, 0x97, 0x7f, 0x36, 0xeb, 0x7f,
	0x46, 0xb8, 0x15, 0x82, 0xe8, 0x73, 0x99, 0xbc, 0x2c, 0xb8, 0x1e, 0x3f, 0xaf, 0x46, 0x45, 0xfc,
	0xb2, 0xe0, 0x60, 0x48, 0x1f, 0xdf, 0xbc, 0x2c, 0xb3, 0x0e, 0x3a, 0x46, 0x27, 0x87, 0xbd, 0xa7,
	0x3f, 0xbf, 0xb7, 0x1f, 0x8b, 0xd4, 0x0c, 0x8a, 0xf3, 0x20, 0x56, 0x19, 0x3d, 0x53, 0x90, 0xbd,
	0x66, 0x90, 0xd1, 0x2b, 0x06, 0x59, 0x42, 0x47, 0xd5, 0x97, 0x9a, 0x71, 0xce, 0x21, 0x88, 0xd8,
	0xd5, 0x99, 0x92, 0x46, 0xb3, 0xd8, 0x84, 0x1c, 0x80, 0x09, 0x1e, 0xd9, 0x59, 0xc4, 0xc1, 0x7b,
	0xf1, 0x80, 0x49, 0xc9, 0x2f, 0x9c, 0x9d, 0x63, 0x74, 0x72, 0x10, 0x2d, 0x43, 0xf2, 0x10, 0x1f,
	0xb0, 0xc2, 0x0c, 0x94, 0x4e, 0xcd, 0xd8, 0xd9, 0x2d, 0x6b, 0x3d, 0xe7, 0xeb, 0xa7, 0x4e, 0x73,
	0xc1, 0xf1, 0x2c, 0x49, 0x34, 0x07, 0xe8, 0x1b, 0x9d, 0x4a, 0x11, 0xad, 0xa4, 0xfe, 0x23, 0xec,
	0xd6, 0x31, 0x40, 0xae, 0x24, 0x70, 0xe2, 0xe2, 0x7d, 0x28, 0x79, 0x64, 0xcc, 0x2b, 0x8e, 0xdd,
	0xe8, 0x77, 0xec, 0xbf, 0x47, 0xf8, 0x4e, 0x08, 0xe2, 0x55, 0x9e, 0x30, 0xc3, 0x37, 0xd9, 0xbb,
	0x78, 0x8f, 0xd9, 0x03, 0xab, 0xae, 0xff, 0x5d, 0x65, 0x29, 0xdc, 0x04, 0xb8, 0xb1, 0x35, 0xc0,
	0x93, 0xfd, 0x8f, 0xd7, 0x6d, 0x34, 0xbf, 0x6e, 0x23, 0xbf, 0x85, 0x8f, 0xfe, 0xba, 0x8f, 0xe5,
	0xe8, 0xce, 0x11, 0xde, 0x09, 0x41, 0x90, 0x77, 0xf8, 0x70, 0xbd, 0x4e, 0xee, 0x05, 0x75, 0x56,
	0x08, 0xea, 0xb1, 0xdc, 0xce, 0x96, 0xea, 0xc5, 0xe3, 0x19, 0x7c, 0xfb, 0x8f, 0x77, 0x25, 0xf4,
	0x9f, 0x13, 0xea, 0x5d, 0xe4, 0xde, 0xdf, 0xbe, 0xc1, 0x9e, 0xda, 0x13, 0x5f, 0xa6, 0x1e, 0x9a,
	0x4c, 0x3d, 0xf4, 0x63, 0xea, 0xa1, 0x0f, 0x33, 0xaf, 0x31, 0x99, 0x79, 0x8d, 0x6f, 0x33, 0xaf,
	0x81, 0x8f, 0x52, 0x55, 0x3b, 0xed, 0x05, 0x7a, 0xd3, 0x5d, 0x73, 0xe6, 0x4a, 0xd2, 0x49, 0xd5,
	0x5a, 0x44, 0x47, 0xcb, 0x5d, 0xaa, 0x5c, 0x7a, 0x7e, 0xab, 0xda, 0x82, 0x07, 0xbf, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x49, 0x33, 0xee, 0xe2, 0x6d, 0x03, 0x00, 0x00,
}

func (this *MsgUpdateOracleRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgUpdateOracleRequest)
	if !ok {
		that2, ok := that.(MsgUpdateOracleRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if this.Authority != that1.Authority {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateOracle is the RPC endpoint for updating the oracle
	UpdateOracle(ctx context.Context, in *MsgUpdateOracleRequest, opts ...grpc.CallOption) (*MsgUpdateOracleResponse, error)
	// SendQueryOracle sends a query to an oracle on another chain
	SendQueryOracle(ctx context.Context, in *MsgSendQueryOracleRequest, opts ...grpc.CallOption) (*MsgSendQueryOracleResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateOracle(ctx context.Context, in *MsgUpdateOracleRequest, opts ...grpc.CallOption) (*MsgUpdateOracleResponse, error) {
	out := new(MsgUpdateOracleResponse)
	err := c.cc.Invoke(ctx, "/provenance.oracle.v1.Msg/UpdateOracle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendQueryOracle(ctx context.Context, in *MsgSendQueryOracleRequest, opts ...grpc.CallOption) (*MsgSendQueryOracleResponse, error) {
	out := new(MsgSendQueryOracleResponse)
	err := c.cc.Invoke(ctx, "/provenance.oracle.v1.Msg/SendQueryOracle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateOracle is the RPC endpoint for updating the oracle
	UpdateOracle(context.Context, *MsgUpdateOracleRequest) (*MsgUpdateOracleResponse, error)
	// SendQueryOracle sends a query to an oracle on another chain
	SendQueryOracle(context.Context, *MsgSendQueryOracleRequest) (*MsgSendQueryOracleResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateOracle(ctx context.Context, req *MsgUpdateOracleRequest) (*MsgUpdateOracleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOracle not implemented")
}
func (*UnimplementedMsgServer) SendQueryOracle(ctx context.Context, req *MsgSendQueryOracleRequest) (*MsgSendQueryOracleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendQueryOracle not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateOracleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provenance.oracle.v1.Msg/UpdateOracle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateOracle(ctx, req.(*MsgUpdateOracleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendQueryOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendQueryOracleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendQueryOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/provenance.oracle.v1.Msg/SendQueryOracle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendQueryOracle(ctx, req.(*MsgSendQueryOracleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "provenance.oracle.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateOracle",
			Handler:    _Msg_UpdateOracle_Handler,
		},
		{
			MethodName: "SendQueryOracle",
			Handler:    _Msg_SendQueryOracle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provenance/oracle/v1/tx.proto",
}

func (m *MsgSendQueryOracleRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendQueryOracleRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendQueryOracleRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Query) > 0 {
		i -= len(m.Query)
		copy(dAtA[i:], m.Query)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Query)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendQueryOracleResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendQueryOracleResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendQueryOracleResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateOracleRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateOracleRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateOracleRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateOracleResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateOracleResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateOracleResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSendQueryOracleRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSendQueryOracleResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovTx(uint64(m.Sequence))
	}
	return n
}

func (m *MsgUpdateOracleRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateOracleResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSendQueryOracleRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSendQueryOracleRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendQueryOracleRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = append(m.Query[:0], dAtA[iNdEx:postIndex]...)
			if m.Query == nil {
				m.Query = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSendQueryOracleResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSendQueryOracleResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendQueryOracleResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateOracleRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateOracleRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateOracleRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateOracleResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateOracleResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateOracleResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)

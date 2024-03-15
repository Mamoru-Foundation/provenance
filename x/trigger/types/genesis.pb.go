// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/trigger/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the trigger module's genesis state.
type GenesisState struct {
	// Trigger id is the next auto incremented id to be assigned to the next created trigger
	TriggerId uint64 `protobuf:"varint,1,opt,name=trigger_id,json=triggerId,proto3" json:"trigger_id,omitempty"`
	// Queue start is the starting index of the queue.
	QueueStart uint64 `protobuf:"varint,2,opt,name=queue_start,json=queueStart,proto3" json:"queue_start,omitempty"`
	// Triggers to initially start with.
	Triggers []Trigger `protobuf:"bytes,3,rep,name=triggers,proto3" json:"triggers"`
	// Maximum amount of gas that the triggers can use.
	GasLimits []GasLimit `protobuf:"bytes,4,rep,name=gas_limits,json=gasLimits,proto3" json:"gas_limits"`
	// Triggers to initially start with in the queue.
	QueuedTriggers []QueuedTrigger `protobuf:"bytes,5,rep,name=queued_triggers,json=queuedTriggers,proto3" json:"queued_triggers"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e92f7d1706d41c9, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

// GasLimit defines the trigger module's grouping of a trigger and a gas limit
type GasLimit struct {
	// The identifier of the trigger this GasLimit belongs to.
	TriggerId uint64 `protobuf:"varint,1,opt,name=trigger_id,json=triggerId,proto3" json:"trigger_id,omitempty"`
	// The maximum amount of gas that the trigger can use.
	Amount uint64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *GasLimit) Reset()         { *m = GasLimit{} }
func (m *GasLimit) String() string { return proto.CompactTextString(m) }
func (*GasLimit) ProtoMessage()    {}
func (*GasLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e92f7d1706d41c9, []int{1}
}
func (m *GasLimit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GasLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GasLimit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GasLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasLimit.Merge(m, src)
}
func (m *GasLimit) XXX_Size() int {
	return m.Size()
}
func (m *GasLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_GasLimit.DiscardUnknown(m)
}

var xxx_messageInfo_GasLimit proto.InternalMessageInfo

func (m *GasLimit) GetTriggerId() uint64 {
	if m != nil {
		return m.TriggerId
	}
	return 0
}

func (m *GasLimit) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "provenance.trigger.v1.GenesisState")
	proto.RegisterType((*GasLimit)(nil), "provenance.trigger.v1.GasLimit")
}

func init() {
	proto.RegisterFile("provenance/trigger/v1/genesis.proto", fileDescriptor_5e92f7d1706d41c9)
}

var fileDescriptor_5e92f7d1706d41c9 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x28, 0xca, 0x2f,
	0x4b, 0xcd, 0x4b, 0xcc, 0x4b, 0x4e, 0xd5, 0x2f, 0x29, 0xca, 0x4c, 0x4f, 0x4f, 0x2d, 0xd2, 0x2f,
	0x33, 0xd4, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x45, 0x28, 0xd2, 0x83, 0x2a, 0xd2, 0x2b, 0x33, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0xab, 0xd0, 0x07, 0xb1, 0x20, 0x8a, 0xa5, 0x70, 0x98, 0x08, 0xd3, 0x07, 0x56, 0xa4, 0xb4, 0x95,
	0x89, 0x8b, 0xc7, 0x1d, 0x62, 0x47, 0x70, 0x49, 0x62, 0x49, 0xaa, 0x90, 0x2c, 0x17, 0x17, 0x54,
	0x45, 0x7c, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x27, 0x54, 0xc4, 0x33, 0x45,
	0x48, 0x9e, 0x8b, 0xbb, 0xb0, 0x34, 0xb5, 0x34, 0x35, 0xbe, 0xb8, 0x24, 0xb1, 0xa8, 0x44, 0x82,
	0x09, 0x2c, 0xcf, 0x05, 0x16, 0x0a, 0x06, 0x89, 0x08, 0x39, 0x70, 0x71, 0x40, 0x55, 0x17, 0x4b,
	0x30, 0x2b, 0x30, 0x6b, 0x70, 0x1b, 0xc9, 0xe9, 0x61, 0x75, 0xb5, 0x5e, 0x08, 0x84, 0xe9, 0xc4,
	0x72, 0xe2, 0x9e, 0x3c, 0x43, 0x10, 0x5c, 0x97, 0x90, 0x0b, 0x17, 0x57, 0x7a, 0x62, 0x71, 0x7c,
	0x4e, 0x66, 0x6e, 0x66, 0x49, 0xb1, 0x04, 0x0b, 0xd8, 0x0c, 0x79, 0x1c, 0x66, 0xb8, 0x27, 0x16,
	0xfb, 0x80, 0xd4, 0x41, 0x0d, 0xe1, 0x4c, 0x87, 0xf2, 0x8b, 0x85, 0x82, 0xb9, 0xf8, 0xc1, 0xae,
	0x4a, 0x89, 0x87, 0x3b, 0x87, 0x15, 0x6c, 0x94, 0x0a, 0x0e, 0xa3, 0x02, 0xc1, 0xaa, 0x51, 0x1d,
	0xc5, 0x57, 0x88, 0x2c, 0x58, 0x6c, 0xc5, 0xd1, 0xb1, 0x40, 0x9e, 0xe1, 0xc5, 0x02, 0x79, 0x06,
	0x25, 0x47, 0x2e, 0x0e, 0x98, 0xdd, 0x84, 0x82, 0x4c, 0x8c, 0x8b, 0x2d, 0x31, 0x37, 0xbf, 0x34,
	0x0f, 0x16, 0x5a, 0x50, 0x9e, 0x53, 0xe6, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e,
	0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31,
	0x70, 0x49, 0x64, 0xe6, 0x63, 0x77, 0x64, 0x00, 0x63, 0x94, 0x71, 0x7a, 0x66, 0x49, 0x46, 0x69,
	0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x42, 0x8d, 0x6e, 0x66, 0x3e, 0x12, 0x4f, 0xbf, 0x02, 0x1e,
	0xe1, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0xc8, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x99, 0x56, 0xc9, 0x58, 0x65, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.QueuedTriggers) > 0 {
		for iNdEx := len(m.QueuedTriggers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.QueuedTriggers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.GasLimits) > 0 {
		for iNdEx := len(m.GasLimits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasLimits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Triggers) > 0 {
		for iNdEx := len(m.Triggers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Triggers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.QueueStart != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.QueueStart))
		i--
		dAtA[i] = 0x10
	}
	if m.TriggerId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TriggerId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GasLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasLimit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasLimit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if m.TriggerId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TriggerId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TriggerId != 0 {
		n += 1 + sovGenesis(uint64(m.TriggerId))
	}
	if m.QueueStart != 0 {
		n += 1 + sovGenesis(uint64(m.QueueStart))
	}
	if len(m.Triggers) > 0 {
		for _, e := range m.Triggers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GasLimits) > 0 {
		for _, e := range m.GasLimits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.QueuedTriggers) > 0 {
		for _, e := range m.QueuedTriggers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GasLimit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TriggerId != 0 {
		n += 1 + sovGenesis(uint64(m.TriggerId))
	}
	if m.Amount != 0 {
		n += 1 + sovGenesis(uint64(m.Amount))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TriggerId", wireType)
			}
			m.TriggerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TriggerId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueueStart", wireType)
			}
			m.QueueStart = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QueueStart |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Triggers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Triggers = append(m.Triggers, Trigger{})
			if err := m.Triggers[len(m.Triggers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasLimits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasLimits = append(m.GasLimits, GasLimit{})
			if err := m.GasLimits[len(m.GasLimits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueuedTriggers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueuedTriggers = append(m.QueuedTriggers, QueuedTrigger{})
			if err := m.QueuedTriggers[len(m.QueuedTriggers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GasLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GasLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TriggerId", wireType)
			}
			m.TriggerId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TriggerId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)

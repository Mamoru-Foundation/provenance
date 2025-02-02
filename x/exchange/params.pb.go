// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/exchange/v1/params.proto

package exchange

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Params is a representation of the exchange module parameters.
type Params struct {
	// default_split is the default proportion of fees the exchange receives in basis points.
	// It is used if there isn't an applicable denom-specific split defined.
	// E.g. 100 = 1%. Min = 0, Max = 10000.
	DefaultSplit uint32 `protobuf:"varint,1,opt,name=default_split,json=defaultSplit,proto3" json:"default_split,omitempty"`
	// denom_splits are the denom-specific amounts the exchange receives.
	DenomSplits []DenomSplit `protobuf:"bytes,2,rep,name=denom_splits,json=denomSplits,proto3" json:"denom_splits"`
	// fee_create_payment_flat is the flat fee options for creating a payment.
	// If the source amount is not zero then one of these fee entries is required to create the payment.
	// This field is currently limited to zero or one entries.
	FeeCreatePaymentFlat []types.Coin `protobuf:"bytes,3,rep,name=fee_create_payment_flat,json=feeCreatePaymentFlat,proto3" json:"fee_create_payment_flat"`
	// fee_accept_payment_flat is the flat fee options for accepting a payment.
	// If the target amount is not zero then one of these fee entries is required to accept the payment.
	// This field is currently limited to zero or one entries.
	FeeAcceptPaymentFlat []types.Coin `protobuf:"bytes,4,rep,name=fee_accept_payment_flat,json=feeAcceptPaymentFlat,proto3" json:"fee_accept_payment_flat"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d689cfc7a7422f1, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetDefaultSplit() uint32 {
	if m != nil {
		return m.DefaultSplit
	}
	return 0
}

func (m *Params) GetDenomSplits() []DenomSplit {
	if m != nil {
		return m.DenomSplits
	}
	return nil
}

func (m *Params) GetFeeCreatePaymentFlat() []types.Coin {
	if m != nil {
		return m.FeeCreatePaymentFlat
	}
	return nil
}

func (m *Params) GetFeeAcceptPaymentFlat() []types.Coin {
	if m != nil {
		return m.FeeAcceptPaymentFlat
	}
	return nil
}

// DenomSplit associates a coin denomination with an amount the exchange receives for that denom.
type DenomSplit struct {
	// denom is the coin denomination this split applies to.
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// split is the proportion of fees the exchange receives for this denom in basis points.
	// E.g. 100 = 1%. Min = 0, Max = 10000.
	Split uint32 `protobuf:"varint,2,opt,name=split,proto3" json:"split,omitempty"`
}

func (m *DenomSplit) Reset()         { *m = DenomSplit{} }
func (m *DenomSplit) String() string { return proto.CompactTextString(m) }
func (*DenomSplit) ProtoMessage()    {}
func (*DenomSplit) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d689cfc7a7422f1, []int{1}
}
func (m *DenomSplit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomSplit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomSplit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomSplit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomSplit.Merge(m, src)
}
func (m *DenomSplit) XXX_Size() int {
	return m.Size()
}
func (m *DenomSplit) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomSplit.DiscardUnknown(m)
}

var xxx_messageInfo_DenomSplit proto.InternalMessageInfo

func (m *DenomSplit) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *DenomSplit) GetSplit() uint32 {
	if m != nil {
		return m.Split
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "provenance.exchange.v1.Params")
	proto.RegisterType((*DenomSplit)(nil), "provenance.exchange.v1.DenomSplit")
}

func init() {
	proto.RegisterFile("provenance/exchange/v1/params.proto", fileDescriptor_5d689cfc7a7422f1)
}

var fileDescriptor_5d689cfc7a7422f1 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x6e, 0xea, 0x30,
	0x14, 0xc6, 0x13, 0xe0, 0x22, 0x5d, 0x03, 0x4b, 0x84, 0xee, 0x05, 0x06, 0x5f, 0x04, 0x0b, 0xcb,
	0xb5, 0x45, 0xbb, 0x74, 0x2d, 0x54, 0x5d, 0xba, 0x20, 0x2a, 0x75, 0xe8, 0x12, 0x39, 0xe6, 0x24,
	0x44, 0x4a, 0xec, 0x28, 0x36, 0x88, 0xbe, 0x45, 0x9f, 0xa3, 0x4f, 0xc2, 0xc8, 0xd8, 0xa9, 0xaa,
	0xe0, 0x45, 0xaa, 0xd8, 0xfc, 0xab, 0xd4, 0x4a, 0xdd, 0x72, 0x4e, 0xbe, 0xef, 0xa7, 0xef, 0x7c,
	0x09, 0xea, 0x67, 0xb9, 0x5c, 0x82, 0x60, 0x82, 0x03, 0x85, 0x15, 0x9f, 0x33, 0x11, 0x01, 0x5d,
	0x0e, 0x69, 0xc6, 0x72, 0x96, 0x2a, 0x92, 0xe5, 0x52, 0x4b, 0xef, 0xcf, 0x49, 0x44, 0x0e, 0x22,
	0xb2, 0x1c, 0x76, 0x9a, 0x91, 0x8c, 0xa4, 0x91, 0xd0, 0xe2, 0xc9, 0xaa, 0x3b, 0x98, 0x4b, 0x95,
	0x4a, 0x45, 0x03, 0xa6, 0x0a, 0x54, 0x00, 0x9a, 0x0d, 0x29, 0x97, 0xb1, 0xb0, 0xef, 0x7b, 0x2f,
	0x25, 0x54, 0x9d, 0x18, 0xbc, 0xd7, 0x47, 0x8d, 0x19, 0x84, 0x6c, 0x91, 0x68, 0x5f, 0x65, 0x49,
	0xac, 0x5b, 0x6e, 0xd7, 0x1d, 0x34, 0xa6, 0xf5, 0xfd, 0xf2, 0xbe, 0xd8, 0x79, 0x77, 0xa8, 0x3e,
	0x03, 0x21, 0x53, 0x2b, 0x51, 0xad, 0x52, 0xb7, 0x3c, 0xa8, 0x5d, 0xf4, 0xc8, 0xd7, 0xa1, 0xc8,
	0x4d, 0xa1, 0x35, 0xce, 0x51, 0x65, 0xfd, 0xf6, 0xcf, 0x99, 0xd6, 0x66, 0xc7, 0x8d, 0xf2, 0x1e,
	0xd0, 0xdf, 0x10, 0xc0, 0xe7, 0x39, 0x30, 0x0d, 0x7e, 0xc6, 0x9e, 0x52, 0x10, 0xda, 0x0f, 0x13,
	0xa6, 0x5b, 0x65, 0xc3, 0x6d, 0x13, 0x1b, 0x9f, 0x14, 0xf1, 0xc9, 0x3e, 0x3e, 0x19, 0xcb, 0x58,
	0xec, 0x71, 0xcd, 0x10, 0x60, 0x6c, 0xec, 0x13, 0xeb, 0xbe, 0x4d, 0x98, 0x3e, 0x70, 0x19, 0xe7,
	0x90, 0xe9, 0xcf, 0xdc, 0xca, 0xcf, 0xb9, 0xd7, 0xc6, 0x7e, 0xc6, 0xed, 0x5d, 0x21, 0x74, 0x3a,
	0xc8, 0x6b, 0xa2, 0x5f, 0xe6, 0x18, 0xd3, 0xd3, 0xef, 0xa9, 0x1d, 0x8a, 0xad, 0x6d, 0xaf, 0x64,
	0xda, 0xb3, 0xc3, 0x08, 0xd6, 0x5b, 0xec, 0x6e, 0xb6, 0xd8, 0x7d, 0xdf, 0x62, 0xf7, 0x79, 0x87,
	0x9d, 0xcd, 0x0e, 0x3b, 0xaf, 0x3b, 0xec, 0xa0, 0x76, 0x2c, 0xbf, 0x29, 0x6f, 0xe2, 0x3e, 0x92,
	0x28, 0xd6, 0xf3, 0x45, 0x40, 0xb8, 0x4c, 0xe9, 0x49, 0xf4, 0x3f, 0x96, 0x67, 0x13, 0x5d, 0x1d,
	0xff, 0x95, 0xa0, 0x6a, 0x3e, 0xea, 0xe5, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0x5f, 0x8e,
	0x59, 0x49, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeeAcceptPaymentFlat) > 0 {
		for iNdEx := len(m.FeeAcceptPaymentFlat) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeAcceptPaymentFlat[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.FeeCreatePaymentFlat) > 0 {
		for iNdEx := len(m.FeeCreatePaymentFlat) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeCreatePaymentFlat[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DenomSplits) > 0 {
		for iNdEx := len(m.DenomSplits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomSplits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.DefaultSplit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DefaultSplit))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DenomSplit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomSplit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomSplit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Split != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.Split))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DefaultSplit != 0 {
		n += 1 + sovParams(uint64(m.DefaultSplit))
	}
	if len(m.DenomSplits) > 0 {
		for _, e := range m.DenomSplits {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.FeeCreatePaymentFlat) > 0 {
		for _, e := range m.FeeCreatePaymentFlat {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if len(m.FeeAcceptPaymentFlat) > 0 {
		for _, e := range m.FeeAcceptPaymentFlat {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func (m *DenomSplit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.Split != 0 {
		n += 1 + sovParams(uint64(m.Split))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultSplit", wireType)
			}
			m.DefaultSplit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DefaultSplit |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomSplits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomSplits = append(m.DenomSplits, DenomSplit{})
			if err := m.DenomSplits[len(m.DenomSplits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeCreatePaymentFlat", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeCreatePaymentFlat = append(m.FeeCreatePaymentFlat, types.Coin{})
			if err := m.FeeCreatePaymentFlat[len(m.FeeCreatePaymentFlat)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeAcceptPaymentFlat", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeAcceptPaymentFlat = append(m.FeeAcceptPaymentFlat, types.Coin{})
			if err := m.FeeAcceptPaymentFlat[len(m.FeeAcceptPaymentFlat)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *DenomSplit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: DenomSplit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomSplit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Split", wireType)
			}
			m.Split = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Split |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)

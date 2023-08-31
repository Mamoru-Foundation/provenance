// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/hold/v1/hold.proto

package hold

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// AccountHold associates an address with an amount on hold for that address.
type AccountHold struct {
	// address is the account address that holds the funds on hold.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// amount is the balances that are on hold for the address.
	Amount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *AccountHold) Reset()         { *m = AccountHold{} }
func (m *AccountHold) String() string { return proto.CompactTextString(m) }
func (*AccountHold) ProtoMessage()    {}
func (*AccountHold) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc6e4f15dd47e2b, []int{0}
}
func (m *AccountHold) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountHold) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountHold.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountHold) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountHold.Merge(m, src)
}
func (m *AccountHold) XXX_Size() int {
	return m.Size()
}
func (m *AccountHold) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountHold.DiscardUnknown(m)
}

var xxx_messageInfo_AccountHold proto.InternalMessageInfo

func (m *AccountHold) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AccountHold) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountHold)(nil), "provenance.hold.v1.AccountHold")
}

func init() { proto.RegisterFile("provenance/hold/v1/hold.proto", fileDescriptor_cfc6e4f15dd47e2b) }

var fileDescriptor_cfc6e4f15dd47e2b = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x28, 0xca, 0x2f,
	0x4b, 0xcd, 0x4b, 0xcc, 0x4b, 0x4e, 0xd5, 0xcf, 0xc8, 0xcf, 0x49, 0xd1, 0x2f, 0x33, 0x04, 0xd3,
	0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x42, 0x08, 0x69, 0x3d, 0xb0, 0x70, 0x99, 0xa1, 0x94,
	0x5c, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0xb1, 0x7e, 0x52, 0x62, 0x71, 0xaa, 0x7e, 0x99, 0x61, 0x52,
	0x6a, 0x49, 0xa2, 0xa1, 0x7e, 0x72, 0x7e, 0x66, 0x1e, 0x44, 0x8f, 0x94, 0x48, 0x7a, 0x7e, 0x7a,
	0x3e, 0x98, 0xa9, 0x0f, 0x62, 0x41, 0x44, 0x95, 0x7a, 0x18, 0xb9, 0xb8, 0x1d, 0x93, 0x93, 0xf3,
	0x4b, 0xf3, 0x4a, 0x3c, 0xf2, 0x73, 0x52, 0x84, 0x24, 0xb8, 0xd8, 0x13, 0x53, 0x52, 0x8a, 0x52,
	0x8b, 0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0xa1, 0x64, 0x2e, 0xb6, 0xc4,
	0x5c, 0x90, 0x3a, 0x09, 0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x49, 0x3d, 0x88, 0x85, 0x7a, 0x20,
	0x0b, 0xf5, 0xa0, 0x16, 0xea, 0x39, 0xe7, 0x67, 0xe6, 0x39, 0x19, 0x9c, 0xb8, 0x27, 0xcf, 0xb0,
	0xea, 0xbe, 0xbc, 0x46, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xd4,
	0x75, 0x10, 0x4a, 0xb7, 0x38, 0x25, 0x5b, 0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x18, 0xac, 0xa1, 0x38,
	0x08, 0x6a, 0xb4, 0x53, 0xec, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x70, 0x89,
	0x66, 0x82, 0x9d, 0x8c, 0xe6, 0xeb, 0x00, 0xc6, 0x28, 0x2d, 0x24, 0x4b, 0x10, 0x0a, 0x74, 0x33,
	0xf3, 0x91, 0x78, 0xfa, 0x15, 0xe0, 0xd0, 0x4b, 0x62, 0x03, 0x7b, 0xda, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x55, 0x00, 0xe4, 0x28, 0x5f, 0x01, 0x00, 0x00,
}

func (m *AccountHold) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountHold) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountHold) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHold(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintHold(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHold(dAtA []byte, offset int, v uint64) int {
	offset -= sovHold(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccountHold) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovHold(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovHold(uint64(l))
		}
	}
	return n
}

func sovHold(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHold(x uint64) (n int) {
	return sovHold(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountHold) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHold
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
			return fmt.Errorf("proto: AccountHold: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountHold: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHold
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
				return ErrInvalidLengthHold
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHold
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHold
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
				return ErrInvalidLengthHold
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHold
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHold(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHold
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
func skipHold(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHold
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
					return 0, ErrIntOverflowHold
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
					return 0, ErrIntOverflowHold
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
				return 0, ErrInvalidLengthHold
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHold
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHold
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHold        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHold          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHold = fmt.Errorf("proto: unexpected end of group")
)

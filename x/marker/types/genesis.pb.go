// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/marker/v1/genesis.proto

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

// GenesisState defines the account module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// A collection of marker accounts to create on start
	Markers []MarkerAccount `protobuf:"bytes,2,rep,name=markers,proto3" json:"markers"`
	// list of marker net asset values
	NetAssetValues []MarkerNetAssetValues `protobuf:"bytes,3,rep,name=net_asset_values,json=netAssetValues,proto3" json:"net_asset_values"`
	// list of denom based denied send addresses
	DenySendAddresses []DenySendAddress `protobuf:"bytes,4,rep,name=deny_send_addresses,json=denySendAddresses,proto3" json:"deny_send_addresses"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dcc4ab7c9d2f78f, []int{0}
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

// DenySendAddress defines addresses that are denied sends for marker denom
type DenySendAddress struct {
	// marker_address is the marker's address for denied address
	MarkerAddress string `protobuf:"bytes,1,opt,name=marker_address,json=markerAddress,proto3" json:"marker_address,omitempty"`
	// deny_address defines all wallet addresses that are denied sends for the marker
	DenyAddress string `protobuf:"bytes,2,opt,name=deny_address,json=denyAddress,proto3" json:"deny_address,omitempty"`
}

func (m *DenySendAddress) Reset()         { *m = DenySendAddress{} }
func (m *DenySendAddress) String() string { return proto.CompactTextString(m) }
func (*DenySendAddress) ProtoMessage()    {}
func (*DenySendAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dcc4ab7c9d2f78f, []int{1}
}
func (m *DenySendAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenySendAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenySendAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenySendAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenySendAddress.Merge(m, src)
}
func (m *DenySendAddress) XXX_Size() int {
	return m.Size()
}
func (m *DenySendAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_DenySendAddress.DiscardUnknown(m)
}

var xxx_messageInfo_DenySendAddress proto.InternalMessageInfo

// MarkerNetAssetValues defines the net asset values for a marker
type MarkerNetAssetValues struct {
	// address defines the marker address
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// net_asset_values that are assigned to marker
	NetAssetValues []NetAssetValue `protobuf:"bytes,2,rep,name=net_asset_values,json=netAssetValues,proto3" json:"net_asset_values"`
}

func (m *MarkerNetAssetValues) Reset()         { *m = MarkerNetAssetValues{} }
func (m *MarkerNetAssetValues) String() string { return proto.CompactTextString(m) }
func (*MarkerNetAssetValues) ProtoMessage()    {}
func (*MarkerNetAssetValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dcc4ab7c9d2f78f, []int{2}
}
func (m *MarkerNetAssetValues) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MarkerNetAssetValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MarkerNetAssetValues.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MarkerNetAssetValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkerNetAssetValues.Merge(m, src)
}
func (m *MarkerNetAssetValues) XXX_Size() int {
	return m.Size()
}
func (m *MarkerNetAssetValues) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkerNetAssetValues.DiscardUnknown(m)
}

var xxx_messageInfo_MarkerNetAssetValues proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "provenance.marker.v1.GenesisState")
	proto.RegisterType((*DenySendAddress)(nil), "provenance.marker.v1.DenySendAddress")
	proto.RegisterType((*MarkerNetAssetValues)(nil), "provenance.marker.v1.MarkerNetAssetValues")
}

func init() {
	proto.RegisterFile("provenance/marker/v1/genesis.proto", fileDescriptor_5dcc4ab7c9d2f78f)
}

var fileDescriptor_5dcc4ab7c9d2f78f = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0xaa, 0xda, 0x40,
	0x18, 0x86, 0x13, 0x15, 0x6d, 0x47, 0x6b, 0xdb, 0xa9, 0xd0, 0x20, 0x25, 0xfe, 0x14, 0x41, 0x0a,
	0x4d, 0xd0, 0xee, 0xdc, 0x69, 0x0b, 0x5d, 0xb5, 0x88, 0x42, 0x17, 0x76, 0x11, 0xc6, 0xe4, 0x23,
	0x0d, 0xad, 0x33, 0x21, 0x33, 0x86, 0x7a, 0x07, 0xdd, 0xb5, 0x97, 0xe0, 0xe5, 0xb8, 0x94, 0xae,
	0xce, 0xea, 0x70, 0xd0, 0xcd, 0xb9, 0x8c, 0x83, 0x93, 0x04, 0xcd, 0x61, 0x70, 0x37, 0xf3, 0xf1,
	0xbc, 0xcf, 0x3b, 0x99, 0x0c, 0xea, 0x86, 0x11, 0x8b, 0x81, 0x12, 0xea, 0x82, 0xbd, 0x22, 0xd1,
	0x4f, 0x88, 0xec, 0x78, 0x60, 0xfb, 0x40, 0x81, 0x07, 0xdc, 0x0a, 0x23, 0x26, 0x18, 0x6e, 0x9c,
	0x19, 0x2b, 0x61, 0xac, 0x78, 0xd0, 0x6c, 0xf8, 0xcc, 0x67, 0x12, 0xb0, 0x4f, 0xab, 0x84, 0x6d,
	0x76, 0x94, 0xbe, 0x34, 0x25, 0x91, 0xee, 0xff, 0x02, 0xaa, 0x7d, 0x4e, 0x0a, 0xe6, 0x82, 0x08,
	0xc0, 0x23, 0x54, 0x0e, 0x49, 0x44, 0x56, 0xdc, 0xd0, 0xdb, 0x7a, 0xbf, 0x3a, 0x7c, 0x63, 0xa9,
	0x0a, 0xad, 0xa9, 0x64, 0x26, 0xa5, 0xdd, 0x6d, 0x4b, 0x9b, 0xa5, 0x09, 0xfc, 0x11, 0x55, 0x12,
	0x82, 0x1b, 0x85, 0x76, 0xb1, 0x5f, 0x1d, 0xbe, 0x55, 0x87, 0xbf, 0xc8, 0xd5, 0xd8, 0x75, 0xd9,
	0x9a, 0x8a, 0xd4, 0x91, 0x25, 0xf1, 0x02, 0xbd, 0xa0, 0x20, 0x1c, 0xc2, 0x39, 0x08, 0x27, 0x26,
	0xbf, 0xd6, 0xc0, 0x8d, 0xa2, 0xb4, 0xbd, 0xbb, 0x66, 0xfb, 0x0a, 0x62, 0x7c, 0x8a, 0x7c, 0x93,
	0x89, 0x54, 0x5a, 0xa7, 0xb9, 0x29, 0xfe, 0x8e, 0x5e, 0x79, 0x40, 0x37, 0x0e, 0x07, 0xea, 0x39,
	0xc4, 0xf3, 0x22, 0xe0, 0x1c, 0xb8, 0x51, 0x92, 0xfa, 0x9e, 0x5a, 0xff, 0x09, 0xe8, 0x66, 0x0e,
	0xd4, 0x1b, 0x27, 0x78, 0x6a, 0x7e, 0xe9, 0xe5, 0xc7, 0xc0, 0x47, 0x4f, 0xfe, 0x6c, 0x5b, 0xda,
	0xfd, 0xb6, 0xa5, 0x75, 0x01, 0x3d, 0x7f, 0x94, 0xc2, 0x3d, 0x54, 0x4f, 0x94, 0x59, 0xad, 0xbc,
	0xde, 0xa7, 0xb3, 0x67, 0xc9, 0x34, 0xc3, 0x3a, 0xa8, 0x26, 0x0f, 0x98, 0x41, 0x05, 0x09, 0x55,
	0x4f, 0xb3, 0x14, 0xb9, 0xa8, 0xf9, 0xab, 0xa3, 0x86, 0xea, 0xe3, 0xb1, 0x81, 0x2a, 0xf9, 0x96,
	0x6c, 0x8b, 0xe7, 0x8a, 0xcb, 0xbd, 0xfa, 0xab, 0x72, 0x66, 0xf5, 0xad, 0x9e, 0x4f, 0x34, 0xf1,
	0x77, 0x07, 0x53, 0xdf, 0x1f, 0x4c, 0xfd, 0xee, 0x60, 0xea, 0xff, 0x8e, 0xa6, 0xb6, 0x3f, 0x9a,
	0xda, 0xcd, 0xd1, 0xd4, 0xd0, 0xeb, 0x80, 0x29, 0x0b, 0xa6, 0xfa, 0x62, 0xe8, 0x07, 0xe2, 0xc7,
	0x7a, 0x69, 0xb9, 0x6c, 0x65, 0x9f, 0x91, 0xf7, 0x01, 0xbb, 0xd8, 0xd9, 0xbf, 0xb3, 0x07, 0x2c,
	0x36, 0x21, 0xf0, 0x65, 0x59, 0xbe, 0xde, 0x0f, 0x0f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x2d,
	0x05, 0xd5, 0x32, 0x03, 0x00, 0x00,
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
	if len(m.DenySendAddresses) > 0 {
		for iNdEx := len(m.DenySendAddresses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenySendAddresses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.NetAssetValues) > 0 {
		for iNdEx := len(m.NetAssetValues) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NetAssetValues[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Markers) > 0 {
		for iNdEx := len(m.Markers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Markers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *DenySendAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenySendAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenySendAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DenyAddress) > 0 {
		i -= len(m.DenyAddress)
		copy(dAtA[i:], m.DenyAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DenyAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MarkerAddress) > 0 {
		i -= len(m.MarkerAddress)
		copy(dAtA[i:], m.MarkerAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.MarkerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MarkerNetAssetValues) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MarkerNetAssetValues) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MarkerNetAssetValues) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NetAssetValues) > 0 {
		for iNdEx := len(m.NetAssetValues) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NetAssetValues[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Markers) > 0 {
		for _, e := range m.Markers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NetAssetValues) > 0 {
		for _, e := range m.NetAssetValues {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DenySendAddresses) > 0 {
		for _, e := range m.DenySendAddresses {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *DenySendAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MarkerAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.DenyAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *MarkerNetAssetValues) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.NetAssetValues) > 0 {
		for _, e := range m.NetAssetValues {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Markers", wireType)
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
			m.Markers = append(m.Markers, MarkerAccount{})
			if err := m.Markers[len(m.Markers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetAssetValues", wireType)
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
			m.NetAssetValues = append(m.NetAssetValues, MarkerNetAssetValues{})
			if err := m.NetAssetValues[len(m.NetAssetValues)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenySendAddresses", wireType)
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
			m.DenySendAddresses = append(m.DenySendAddresses, DenySendAddress{})
			if err := m.DenySendAddresses[len(m.DenySendAddresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *DenySendAddress) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DenySendAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenySendAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarkerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MarkerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenyAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenyAddress = string(dAtA[iNdEx:postIndex])
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
func (m *MarkerNetAssetValues) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MarkerNetAssetValues: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MarkerNetAssetValues: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetAssetValues", wireType)
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
			m.NetAssetValues = append(m.NetAssetValues, NetAssetValue{})
			if err := m.NetAssetValues[len(m.NetAssetValues)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

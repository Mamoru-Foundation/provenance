// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/metadata/v1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the account module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// A collection of metadata scopes and specs to create on start
	Scopes                 []Scope                 `protobuf:"bytes,2,rep,name=scopes,proto3" json:"scopes"`
	Sessions               []Session               `protobuf:"bytes,3,rep,name=sessions,proto3" json:"sessions"`
	Records                []Record                `protobuf:"bytes,4,rep,name=records,proto3" json:"records"`
	ScopeSpecifications    []ScopeSpecification    `protobuf:"bytes,5,rep,name=scope_specifications,json=scopeSpecifications,proto3" json:"scope_specifications"`
	ContractSpecifications []ContractSpecification `protobuf:"bytes,6,rep,name=contract_specifications,json=contractSpecifications,proto3" json:"contract_specifications"`
	RecordSpecifications   []RecordSpecification   `protobuf:"bytes,7,rep,name=record_specifications,json=recordSpecifications,proto3" json:"record_specifications"`
	OSLocatorParams        OSLocatorParams         `protobuf:"bytes,8,opt,name=o_s_locator_params,json=oSLocatorParams,proto3" json:"o_s_locator_params"`
	ObjectStoreLocators    []ObjectStoreLocator    `protobuf:"bytes,9,rep,name=object_store_locators,json=objectStoreLocators,proto3" json:"object_store_locators"`
	// Net asset values assigned to scopes
	NetAssetValues []MarkerNetAssetValues `protobuf:"bytes,10,rep,name=net_asset_values,json=netAssetValues,proto3" json:"net_asset_values"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a835c20198efc302, []int{0}
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

// MarkerNetAssetValues defines the net asset values for a scope
type MarkerNetAssetValues struct {
	// address defines the scope address
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// net_asset_values that are assigned to scope
	NetAssetValues []NetAssetValue `protobuf:"bytes,2,rep,name=net_asset_values,json=netAssetValues,proto3" json:"net_asset_values"`
}

func (m *MarkerNetAssetValues) Reset()         { *m = MarkerNetAssetValues{} }
func (m *MarkerNetAssetValues) String() string { return proto.CompactTextString(m) }
func (*MarkerNetAssetValues) ProtoMessage()    {}
func (*MarkerNetAssetValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_a835c20198efc302, []int{1}
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
	proto.RegisterType((*GenesisState)(nil), "provenance.metadata.v1.GenesisState")
	proto.RegisterType((*MarkerNetAssetValues)(nil), "provenance.metadata.v1.MarkerNetAssetValues")
}

func init() {
	proto.RegisterFile("provenance/metadata/v1/genesis.proto", fileDescriptor_a835c20198efc302)
}

var fileDescriptor_a835c20198efc302 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x6d, 0x5a, 0x92, 0x74, 0x8b, 0x00, 0x2d, 0x69, 0x31, 0x95, 0x70, 0xaa, 0x88, 0x8a,
	0xa8, 0x50, 0x5b, 0x2d, 0x9c, 0x00, 0x21, 0xb5, 0x1c, 0xb8, 0x00, 0xad, 0x12, 0xc1, 0xa1, 0x42,
	0xb2, 0x36, 0x9b, 0x6d, 0x30, 0x4d, 0x3c, 0xd6, 0xce, 0x36, 0x82, 0x37, 0xe0, 0x08, 0x6f, 0xd0,
	0xc7, 0xe9, 0xb1, 0x47, 0x4e, 0x08, 0x25, 0x17, 0xae, 0xbc, 0x01, 0xca, 0x7a, 0xdd, 0x34, 0x8d,
	0x37, 0xdc, 0x6c, 0xcf, 0xff, 0xfd, 0xff, 0x8c, 0x3d, 0x5e, 0xf2, 0x20, 0x95, 0x30, 0x10, 0x09,
	0x4b, 0xb8, 0x08, 0xfb, 0x42, 0xb1, 0x0e, 0x53, 0x2c, 0x1c, 0x6c, 0x87, 0x5d, 0x91, 0x08, 0x8c,
	0x31, 0x48, 0x25, 0x28, 0xa0, 0xab, 0x13, 0x55, 0x90, 0xab, 0x82, 0xc1, 0xf6, 0x5a, 0xb5, 0x0b,
	0x5d, 0xd0, 0x92, 0x70, 0x7c, 0x95, 0xa9, 0xd7, 0x36, 0x2c, 0x9e, 0x17, 0x64, 0x26, 0xab, 0x5b,
	0x64, 0xc8, 0x21, 0x15, 0x46, 0xb3, 0x69, 0xd3, 0xa4, 0x82, 0xc7, 0x47, 0x31, 0x67, 0x2a, 0x86,
	0xc4, 0x68, 0x1b, 0x16, 0x2d, 0xb4, 0x3f, 0x0b, 0xae, 0x50, 0x81, 0x34, 0xae, 0xf5, 0xbf, 0x25,
	0x72, 0xe3, 0x75, 0x36, 0x60, 0x4b, 0x31, 0x25, 0xe8, 0x0b, 0x52, 0x4a, 0x99, 0x64, 0x7d, 0xf4,
	0xdc, 0x75, 0xb7, 0xb1, 0xbc, 0xe3, 0x07, 0xc5, 0x03, 0x07, 0x07, 0x5a, 0xb5, 0xb7, 0x78, 0xf6,
	0xab, 0xe6, 0x34, 0x0d, 0x43, 0x9f, 0x93, 0x92, 0xee, 0x19, 0xbd, 0x6b, 0xeb, 0x0b, 0x8d, 0xe5,
	0x9d, 0xfb, 0x36, 0xba, 0x35, 0x56, 0xe5, 0x70, 0x86, 0xd0, 0x5d, 0x52, 0x41, 0x81, 0x18, 0x43,
	0x82, 0xde, 0x82, 0xc6, 0x6b, 0x56, 0x3c, 0xd3, 0x19, 0x83, 0x0b, 0x8c, 0xbe, 0x24, 0x65, 0x29,
	0x38, 0xc8, 0x0e, 0x7a, 0x8b, 0xda, 0xc1, 0xda, 0x7e, 0x53, 0xcb, 0x8c, 0x41, 0x0e, 0x51, 0x4e,
	0xaa, 0xba, 0x99, 0x68, 0xea, 0xad, 0xa2, 0x77, 0x5d, 0x9b, 0x6d, 0xce, 0x9d, 0xa6, 0x75, 0x19,
	0x31, 0xc6, 0x77, 0x70, 0xa6, 0x82, 0xb4, 0x47, 0xee, 0x72, 0x48, 0x94, 0x64, 0x5c, 0x5d, 0xcd,
	0x29, 0xe9, 0x9c, 0x2d, 0x5b, 0xce, 0x2b, 0x83, 0x15, 0x45, 0xad, 0xf2, 0xa2, 0x22, 0xd2, 0x23,
	0xb2, 0x92, 0x4d, 0x77, 0x35, 0xab, 0xac, 0xb3, 0x1e, 0xcd, 0x7f, 0x41, 0x45, 0x49, 0x55, 0x39,
	0x5b, 0x42, 0x7a, 0x48, 0x28, 0x44, 0x18, 0xf5, 0x80, 0x33, 0x05, 0x32, 0x32, 0x4b, 0x54, 0xd1,
	0x4b, 0xf4, 0xd0, 0x16, 0xb2, 0xdf, 0x7a, 0x93, 0xe9, 0xa7, 0xb6, 0xe9, 0x16, 0x4c, 0x3f, 0xa6,
	0x1d, 0xb2, 0x92, 0xad, 0x6e, 0xa4, 0x77, 0x37, 0x0f, 0x41, 0x6f, 0x69, 0xfe, 0x77, 0xd9, 0xd7,
	0x50, 0x6b, 0xcc, 0x18, 0xc3, 0xfc, 0xbb, 0xc0, 0x4c, 0x05, 0xe9, 0x47, 0x72, 0x3b, 0x11, 0x2a,
	0x62, 0x88, 0x42, 0x45, 0x03, 0xd6, 0x3b, 0x11, 0xe8, 0x11, 0x1d, 0xf0, 0xd8, 0x16, 0xf0, 0x96,
	0xc9, 0x63, 0x21, 0xdf, 0x09, 0xb5, 0x3b, 0x86, 0x3e, 0x68, 0xc6, 0x44, 0xdc, 0x4c, 0xa6, 0x9e,
	0x3e, 0xab, 0x7c, 0x3b, 0xad, 0x39, 0x7f, 0x4e, 0x6b, 0x4e, 0xfd, 0x87, 0x4b, 0xaa, 0x45, 0x20,
	0xf5, 0x48, 0x99, 0x75, 0x3a, 0x52, 0x60, 0xf6, 0xf3, 0x2d, 0x35, 0xf3, 0x5b, 0xfa, 0xbe, 0xa0,
	0xb5, 0xec, 0x0f, 0xdb, 0xb0, 0xb5, 0x36, 0xe5, 0xfd, 0xbf, 0x9e, 0xf6, 0x8e, 0xcf, 0x86, 0xbe,
	0x7b, 0x3e, 0xf4, 0xdd, 0xdf, 0x43, 0xdf, 0xfd, 0x3e, 0xf2, 0x9d, 0xf3, 0x91, 0xef, 0xfc, 0x1c,
	0xf9, 0x0e, 0xb9, 0x17, 0x83, 0x25, 0xe2, 0xc0, 0x3d, 0x7c, 0xda, 0x8d, 0xd5, 0xa7, 0x93, 0x76,
	0xc0, 0xa1, 0x1f, 0x4e, 0x44, 0x5b, 0x31, 0x5c, 0xba, 0x0b, 0xbf, 0x4c, 0xce, 0x20, 0xf5, 0x35,
	0x15, 0xd8, 0x2e, 0xe9, 0xb3, 0xe7, 0xc9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x60, 0x12,
	0x92, 0x72, 0x05, 0x00, 0x00,
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
			dAtA[i] = 0x52
		}
	}
	if len(m.ObjectStoreLocators) > 0 {
		for iNdEx := len(m.ObjectStoreLocators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ObjectStoreLocators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	{
		size, err := m.OSLocatorParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.RecordSpecifications) > 0 {
		for iNdEx := len(m.RecordSpecifications) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RecordSpecifications[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.ContractSpecifications) > 0 {
		for iNdEx := len(m.ContractSpecifications) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContractSpecifications[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ScopeSpecifications) > 0 {
		for iNdEx := len(m.ScopeSpecifications) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ScopeSpecifications[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Sessions) > 0 {
		for iNdEx := len(m.Sessions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Sessions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Scopes) > 0 {
		for iNdEx := len(m.Scopes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Scopes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Scopes) > 0 {
		for _, e := range m.Scopes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Sessions) > 0 {
		for _, e := range m.Sessions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ScopeSpecifications) > 0 {
		for _, e := range m.ScopeSpecifications {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ContractSpecifications) > 0 {
		for _, e := range m.ContractSpecifications {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RecordSpecifications) > 0 {
		for _, e := range m.RecordSpecifications {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.OSLocatorParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ObjectStoreLocators) > 0 {
		for _, e := range m.ObjectStoreLocators {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Scopes", wireType)
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
			m.Scopes = append(m.Scopes, Scope{})
			if err := m.Scopes[len(m.Scopes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sessions", wireType)
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
			m.Sessions = append(m.Sessions, Session{})
			if err := m.Sessions[len(m.Sessions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
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
			m.Records = append(m.Records, Record{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScopeSpecifications", wireType)
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
			m.ScopeSpecifications = append(m.ScopeSpecifications, ScopeSpecification{})
			if err := m.ScopeSpecifications[len(m.ScopeSpecifications)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractSpecifications", wireType)
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
			m.ContractSpecifications = append(m.ContractSpecifications, ContractSpecification{})
			if err := m.ContractSpecifications[len(m.ContractSpecifications)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecordSpecifications", wireType)
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
			m.RecordSpecifications = append(m.RecordSpecifications, RecordSpecification{})
			if err := m.RecordSpecifications[len(m.RecordSpecifications)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OSLocatorParams", wireType)
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
			if err := m.OSLocatorParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectStoreLocators", wireType)
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
			m.ObjectStoreLocators = append(m.ObjectStoreLocators, ObjectStoreLocator{})
			if err := m.ObjectStoreLocators[len(m.ObjectStoreLocators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
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

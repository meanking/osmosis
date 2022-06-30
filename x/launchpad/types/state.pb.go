// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/launchpad/v1/state.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Sale struct {
	// Destination for the earned token_in
	Treasury string `protobuf:"bytes,1,opt,name=treasury,proto3" json:"treasury,omitempty" yaml:"address"`
	Id       uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// token_out is a token denom to be bootstraped. May be referred as base
	// currency, or a sale token.
	TokenOut string `protobuf:"bytes,3,opt,name=token_out,json=tokenOut,proto3" json:"token_out,omitempty"`
	// token_in is a token denom used to buy sale tokens (`token_out`). May be
	// referred as quote_currency or payment token.
	TokenIn string `protobuf:"bytes,4,opt,name=token_in,json=tokenIn,proto3" json:"token_in,omitempty"`
	// total number of `tokens_out` to be sold during the continous sale.
	TokenOutSupply github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=token_out_supply,json=tokenOutSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"token_out_supply"`
	// start time when the token emission starts.
	StartTime time.Time `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	// end time when the token emission ends. Can't be bigger than start +
	// 139years (to avoid round overflow)
	EndTime time.Time `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time"`
	// Round number when the sale was last time updated.
	Round int64 `protobuf:"varint,8,opt,name=round,proto3" json:"round,omitempty"`
	// Last round of the Sale;
	EndRound     int64                                  `protobuf:"varint,9,opt,name=end_round,json=endRound,proto3" json:"end_round,omitempty"`
	OutRemaining github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=out_remaining,json=outRemaining,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"out_remaining"`
	OutSold      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=out_sold,json=outSold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"out_sold"`
	// out token per share
	OutPerShare github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,12,opt,name=out_per_share,json=outPerShare,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"out_per_share"`
	// total amount of currently staked coins (token_in) but not spent coins.
	Staked github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,13,opt,name=staked,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"staked"`
	// total amount of earned coins (token_in)
	Income github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,14,opt,name=income,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"income"`
	// total amount of shares
	Shares github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,15,opt,name=shares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"shares"`
}

func (m *Sale) Reset()         { *m = Sale{} }
func (m *Sale) String() string { return proto.CompactTextString(m) }
func (*Sale) ProtoMessage()    {}
func (*Sale) Descriptor() ([]byte, []int) {
	return fileDescriptor_271fd8d9436658cb, []int{0}
}
func (m *Sale) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sale) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sale.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Sale) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sale.Merge(m, src)
}
func (m *Sale) XXX_Size() int {
	return m.Size()
}
func (m *Sale) XXX_DiscardUnknown() {
	xxx_messageInfo_Sale.DiscardUnknown(m)
}

var xxx_messageInfo_Sale proto.InternalMessageInfo

// UserPosition represents user account in a sale
type UserPosition struct {
	Shares github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=shares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"shares"`
	// total number of currently staked tokens
	Staked github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=staked,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"staked"`
	// last token/share ratio
	OutPerShare github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=out_per_share,json=outPerShare,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"out_per_share"`
	// amount of token_in spent
	Spent github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=spent,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"spent"`
	// Amount of accumulated, not withdrawn, purchased tokens (token_out)
	Purchased github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=purchased,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"purchased"`
}

func (m *UserPosition) Reset()         { *m = UserPosition{} }
func (m *UserPosition) String() string { return proto.CompactTextString(m) }
func (*UserPosition) ProtoMessage()    {}
func (*UserPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_271fd8d9436658cb, []int{1}
}
func (m *UserPosition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserPosition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPosition.Merge(m, src)
}
func (m *UserPosition) XXX_Size() int {
	return m.Size()
}
func (m *UserPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPosition.DiscardUnknown(m)
}

var xxx_messageInfo_UserPosition proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Sale)(nil), "osmosis.launchpad.v1.Sale")
	proto.RegisterType((*UserPosition)(nil), "osmosis.launchpad.v1.UserPosition")
}

func init() { proto.RegisterFile("osmosis/launchpad/v1/state.proto", fileDescriptor_271fd8d9436658cb) }

var fileDescriptor_271fd8d9436658cb = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0xb3, 0x49, 0x9b, 0x6c, 0xa6, 0x6d, 0x94, 0xa1, 0x87, 0x69, 0x85, 0x4d, 0xc8, 0x41,
	0x72, 0xe9, 0x2e, 0x55, 0x41, 0xf0, 0x22, 0x54, 0x11, 0x0a, 0x82, 0x65, 0xa3, 0x20, 0x5e, 0x96,
	0x49, 0x66, 0xdc, 0x2c, 0xdd, 0x9d, 0x59, 0xe6, 0x47, 0x31, 0xff, 0x80, 0xe7, 0xfe, 0x59, 0x3d,
	0xf6, 0x28, 0x82, 0x55, 0x93, 0xff, 0xc0, 0xbf, 0x40, 0x66, 0x66, 0x13, 0x03, 0x9e, 0xba, 0x39,
	0x65, 0xde, 0xbc, 0xf7, 0xfd, 0xbc, 0x97, 0xef, 0xb0, 0x0f, 0x0c, 0xb8, 0x2c, 0xb8, 0xcc, 0x64,
	0x94, 0x63, 0xcd, 0xa6, 0xb3, 0x12, 0x93, 0xe8, 0xea, 0x34, 0x92, 0x0a, 0x2b, 0x1a, 0x96, 0x82,
	0x2b, 0x0e, 0x0f, 0xab, 0x8a, 0x70, 0x5d, 0x11, 0x5e, 0x9d, 0x1e, 0xf7, 0x53, 0xce, 0xd3, 0x9c,
	0x46, 0xb6, 0x66, 0xa2, 0x3f, 0x47, 0x2a, 0x2b, 0xa8, 0x54, 0xb8, 0x28, 0x9d, 0xec, 0xf8, 0x68,
	0x6a, 0x75, 0x89, 0x8d, 0x22, 0x17, 0x54, 0xa9, 0xc3, 0x94, 0xa7, 0xdc, 0xdd, 0x9b, 0x93, 0xbb,
	0x1d, 0xfe, 0x68, 0x83, 0x9d, 0x31, 0xce, 0x29, 0x0c, 0x81, 0xaf, 0x04, 0xc5, 0x52, 0x8b, 0x39,
	0xf2, 0x06, 0xde, 0xa8, 0x7b, 0x06, 0xff, 0xdc, 0xf5, 0x7b, 0x73, 0x5c, 0xe4, 0x2f, 0x86, 0x98,
	0x10, 0x41, 0xa5, 0x1c, 0xc6, 0xeb, 0x1a, 0xd8, 0x03, 0xcd, 0x8c, 0xa0, 0xe6, 0xc0, 0x1b, 0xed,
	0xc4, 0xcd, 0x8c, 0xc0, 0x47, 0xa0, 0xab, 0xf8, 0x25, 0x65, 0x09, 0xd7, 0x0a, 0xb5, 0x0c, 0x20,
	0xf6, 0xed, 0xc5, 0x3b, 0xad, 0xe0, 0x11, 0x70, 0xe7, 0x24, 0x63, 0x68, 0xc7, 0xe6, 0x3a, 0x36,
	0x3e, 0x67, 0xf0, 0x23, 0x78, 0xb8, 0xd6, 0x25, 0x52, 0x97, 0x65, 0x3e, 0x47, 0xbb, 0xb6, 0x7f,
	0x78, 0x73, 0xd7, 0x6f, 0x7c, 0xbf, 0xeb, 0x3f, 0x4e, 0x33, 0x35, 0xd3, 0x93, 0x70, 0xca, 0x8b,
	0xea, 0x1f, 0x55, 0x3f, 0x27, 0x92, 0x5c, 0x46, 0x6a, 0x5e, 0x52, 0x19, 0x9e, 0x33, 0x15, 0xf7,
	0x56, 0xed, 0xc6, 0x96, 0x02, 0x5f, 0x01, 0x20, 0x15, 0x16, 0x2a, 0x31, 0x26, 0xa1, 0xf6, 0xc0,
	0x1b, 0xed, 0x3d, 0x39, 0x0e, 0x9d, 0x83, 0xe1, 0xca, 0xc1, 0xf0, 0xfd, 0xca, 0xc1, 0x33, 0xdf,
	0xf4, 0xbb, 0xfe, 0xd9, 0xf7, 0xe2, 0xae, 0xd5, 0x99, 0x0c, 0x7c, 0x09, 0x7c, 0xca, 0x88, 0x43,
	0x74, 0xee, 0x81, 0xe8, 0x50, 0x46, 0x2c, 0xe0, 0x10, 0xec, 0x0a, 0xae, 0x19, 0x41, 0xfe, 0xc0,
	0x1b, 0xb5, 0x62, 0x17, 0x18, 0xb7, 0x0c, 0xd6, 0x65, 0xba, 0x36, 0x63, 0xfa, 0xc4, 0x36, 0x39,
	0x06, 0x07, 0xc6, 0x0c, 0x41, 0x0b, 0x9c, 0xb1, 0x8c, 0xa5, 0x08, 0xd4, 0xf2, 0x63, 0x9f, 0x6b,
	0x15, 0xaf, 0x18, 0xf0, 0x1c, 0xf8, 0xd6, 0x61, 0x9e, 0x13, 0xb4, 0x57, 0x8b, 0xd7, 0xe1, 0x5a,
	0x8d, 0x79, 0x4e, 0x60, 0xec, 0xe6, 0x2b, 0xa9, 0x48, 0xe4, 0x0c, 0x0b, 0x8a, 0xf6, 0x6b, 0xf1,
	0xf6, 0xb8, 0x56, 0x17, 0x54, 0x8c, 0x0d, 0x02, 0xbe, 0x01, 0x6d, 0xa9, 0xf0, 0x25, 0x25, 0xe8,
	0xa0, 0x16, 0xac, 0x52, 0x1b, 0x4e, 0xc6, 0xa6, 0xbc, 0xa0, 0xa8, 0x57, 0x8f, 0xe3, 0xd4, 0x76,
	0x1e, 0x33, 0x98, 0x44, 0x0f, 0x6a, 0xce, 0x63, 0xd5, 0xc3, 0xaf, 0x2d, 0xb0, 0xff, 0x41, 0x52,
	0x71, 0xc1, 0x65, 0xa6, 0x32, 0xce, 0x36, 0xc0, 0xde, 0x36, 0xe0, 0x0d, 0xc3, 0x9a, 0x5b, 0x19,
	0xf6, 0xdf, 0x63, 0xb6, 0xb6, 0x7f, 0xcc, 0xd7, 0x60, 0x57, 0x96, 0x94, 0x29, 0xf7, 0xad, 0xdf,
	0x9b, 0xe5, 0xc4, 0xf0, 0x2d, 0xe8, 0x96, 0x5a, 0x4c, 0x67, 0x58, 0x52, 0x52, 0x73, 0x25, 0xfc,
	0x03, 0x9c, 0xc5, 0x37, 0xbf, 0x83, 0xc6, 0xcd, 0x22, 0xf0, 0x6e, 0x17, 0x81, 0xf7, 0x6b, 0x11,
	0x78, 0xd7, 0xcb, 0xa0, 0x71, 0xbb, 0x0c, 0x1a, 0xdf, 0x96, 0x41, 0xe3, 0xd3, 0xb3, 0x0d, 0x60,
	0xb5, 0x79, 0x4f, 0x72, 0x3c, 0x91, 0xab, 0x20, 0xba, 0x7a, 0x1e, 0x7d, 0xd9, 0xd8, 0xd6, 0xb6,
	0xc5, 0xa4, 0x6d, 0x57, 0xc0, 0xd3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x2e, 0x42, 0x63,
	0xcf, 0x05, 0x00, 0x00,
}

func (m *Sale) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sale) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Sale) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Shares.Size()
		i -= size
		if _, err := m.Shares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x7a
	{
		size := m.Income.Size()
		i -= size
		if _, err := m.Income.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x72
	{
		size := m.Staked.Size()
		i -= size
		if _, err := m.Staked.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size := m.OutPerShare.Size()
		i -= size
		if _, err := m.OutPerShare.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.OutSold.Size()
		i -= size
		if _, err := m.OutSold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.OutRemaining.Size()
		i -= size
		if _, err := m.OutRemaining.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.EndRound != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.EndRound))
		i--
		dAtA[i] = 0x48
	}
	if m.Round != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x40
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintState(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x3a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintState(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x32
	{
		size := m.TokenOutSupply.Size()
		i -= size
		if _, err := m.TokenOutSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.TokenIn) > 0 {
		i -= len(m.TokenIn)
		copy(dAtA[i:], m.TokenIn)
		i = encodeVarintState(dAtA, i, uint64(len(m.TokenIn)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TokenOut) > 0 {
		i -= len(m.TokenOut)
		copy(dAtA[i:], m.TokenOut)
		i = encodeVarintState(dAtA, i, uint64(len(m.TokenOut)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Treasury) > 0 {
		i -= len(m.Treasury)
		copy(dAtA[i:], m.Treasury)
		i = encodeVarintState(dAtA, i, uint64(len(m.Treasury)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserPosition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserPosition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserPosition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Purchased.Size()
		i -= size
		if _, err := m.Purchased.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Spent.Size()
		i -= size
		if _, err := m.Spent.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.OutPerShare.Size()
		i -= size
		if _, err := m.OutPerShare.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Staked.Size()
		i -= size
		if _, err := m.Staked.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Shares.Size()
		i -= size
		if _, err := m.Shares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Sale) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Treasury)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovState(uint64(m.Id))
	}
	l = len(m.TokenOut)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.TokenIn)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = m.TokenOutSupply.Size()
	n += 1 + l + sovState(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovState(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovState(uint64(l))
	if m.Round != 0 {
		n += 1 + sovState(uint64(m.Round))
	}
	if m.EndRound != 0 {
		n += 1 + sovState(uint64(m.EndRound))
	}
	l = m.OutRemaining.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.OutSold.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.OutPerShare.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.Staked.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.Income.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.Shares.Size()
	n += 1 + l + sovState(uint64(l))
	return n
}

func (m *UserPosition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Shares.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.Staked.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.OutPerShare.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.Spent.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.Purchased.Size()
	n += 1 + l + sovState(uint64(l))
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Sale) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: Sale: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sale: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Treasury", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Treasury = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenOut = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenIn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOutSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenOutSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndRound", wireType)
			}
			m.EndRound = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndRound |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutRemaining", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OutRemaining.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutSold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OutSold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutPerShare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OutPerShare.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staked", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Staked.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Income", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Income.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Shares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *UserPosition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: UserPosition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserPosition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Shares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staked", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Staked.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutPerShare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OutPerShare.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Purchased", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Purchased.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)

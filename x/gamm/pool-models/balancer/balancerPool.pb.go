// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/gamm/pool-models/balancer/balancerPool.proto

// this is a temporary package setup, needs migration logic

package balancer

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types "github.com/osmosis-labs/osmosis/v7/x/gamm/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Parameters for changing the weights in a balancer pool smoothly from
// a start weight and end weight over a period of time.
// Currently, the only smooth change supported is linear changing between
// the two weights, but more types may be added in the future.
// When these parameters are set, the weight w(t) for pool time `t` is the
// following:
//   t <= start_time: w(t) = initial_pool_weights
//   start_time < t <= start_time + duration:
//     w(t) = initial_pool_weights + (t - start_time) *
//       (target_pool_weights - initial_pool_weights) / (duration)
//   t > start_time + duration: w(t) = target_pool_weights
type SmoothWeightChangeParams struct {
	// The start time for beginning the weight change.
	// If a parameter change / pool instantiation leaves this blank,
	// it should be generated by the state_machine as the current time.
	StartTime time.Time `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	// Duration for the weights to change over
	Duration time.Duration `protobuf:"bytes,2,opt,name=duration,proto3,stdduration" json:"duration,omitempty" yaml:"duration"`
	// The initial pool weights. These are copied from the pool's settings
	// at the time of weight change instantiation.
	// The amount PoolAsset.token.amount field is ignored if present,
	// future type refactorings should just have a type with the denom & weight
	// here.
	InitialPoolWeights []types.PoolAsset `protobuf:"bytes,3,rep,name=initialPoolWeights,proto3" json:"initialPoolWeights" yaml:"initial_pool_weights"`
	// The target pool weights. The pool weights will change linearly with respect
	// to time between start_time, and start_time + duration. The amount
	// PoolAsset.token.amount field is ignored if present, future type
	// refactorings should just have a type with the denom & weight here.
	TargetPoolWeights []types.PoolAsset `protobuf:"bytes,4,rep,name=targetPoolWeights,proto3" json:"targetPoolWeights" yaml:"target_pool_weights"`
}

func (m *SmoothWeightChangeParams) Reset()         { *m = SmoothWeightChangeParams{} }
func (m *SmoothWeightChangeParams) String() string { return proto.CompactTextString(m) }
func (*SmoothWeightChangeParams) ProtoMessage()    {}
func (*SmoothWeightChangeParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e991f749f68c2a4, []int{0}
}
func (m *SmoothWeightChangeParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SmoothWeightChangeParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SmoothWeightChangeParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SmoothWeightChangeParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmoothWeightChangeParams.Merge(m, src)
}
func (m *SmoothWeightChangeParams) XXX_Size() int {
	return m.Size()
}
func (m *SmoothWeightChangeParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SmoothWeightChangeParams.DiscardUnknown(m)
}

var xxx_messageInfo_SmoothWeightChangeParams proto.InternalMessageInfo

func (m *SmoothWeightChangeParams) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *SmoothWeightChangeParams) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *SmoothWeightChangeParams) GetInitialPoolWeights() []types.PoolAsset {
	if m != nil {
		return m.InitialPoolWeights
	}
	return nil
}

func (m *SmoothWeightChangeParams) GetTargetPoolWeights() []types.PoolAsset {
	if m != nil {
		return m.TargetPoolWeights
	}
	return nil
}

// PoolParams defined the parameters that will be managed by the pool
// governance in the future. This params are not managed by the chain
// governance. Instead they will be managed by the token holders of the pool.
// The pool's token holders are specified in future_pool_governor.
type PoolParams struct {
	SwapFee                  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=swapFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swapFee" yaml:"swap_fee"`
	ExitFee                  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=exitFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"exitFee" yaml:"exit_fee"`
	SmoothWeightChangeParams *SmoothWeightChangeParams              `protobuf:"bytes,3,opt,name=smoothWeightChangeParams,proto3" json:"smoothWeightChangeParams,omitempty" yaml:"smooth_weight_change_params"`
}

func (m *PoolParams) Reset()         { *m = PoolParams{} }
func (m *PoolParams) String() string { return proto.CompactTextString(m) }
func (*PoolParams) ProtoMessage()    {}
func (*PoolParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e991f749f68c2a4, []int{1}
}
func (m *PoolParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolParams.Merge(m, src)
}
func (m *PoolParams) XXX_Size() int {
	return m.Size()
}
func (m *PoolParams) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolParams.DiscardUnknown(m)
}

var xxx_messageInfo_PoolParams proto.InternalMessageInfo

func (m *PoolParams) GetSmoothWeightChangeParams() *SmoothWeightChangeParams {
	if m != nil {
		return m.SmoothWeightChangeParams
	}
	return nil
}

type Pool struct {
	Address    string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Id         uint64     `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	PoolParams PoolParams `protobuf:"bytes,3,opt,name=poolParams,proto3" json:"poolParams" yaml:"balancer_pool_params"`
	// This string specifies who will govern the pool in the future.
	// Valid forms of this are:
	// {token name},{duration}
	// {duration}
	// where {token name} if specified is the token which determines the
	// governor, and if not specified is the LP token for this pool.duration is
	// a time specified as 0w,1w,2w, etc. which specifies how long the token
	// would need to be locked up to count in governance. 0w means no lockup.
	// TODO: Further improve these docs
	FuturePoolGovernor string `protobuf:"bytes,4,opt,name=future_pool_governor,json=futurePoolGovernor,proto3" json:"future_pool_governor,omitempty" yaml:"future_pool_governor"`
	// sum of all LP tokens sent out
	TotalShares types1.Coin `protobuf:"bytes,5,opt,name=totalShares,proto3" json:"totalShares" yaml:"total_shares"`
	// These are assumed to be sorted by denomiation.
	// They contain the pool asset and the information about the weight
	PoolAssets []types.PoolAsset `protobuf:"bytes,6,rep,name=poolAssets,proto3" json:"poolAssets" yaml:"pool_assets"`
	// sum of all non-normalized pool weights
	TotalWeight github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=totalWeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"totalWeight" yaml:"total_weight"`
}

func (pa *Pool) Reset()     { *pa = Pool{} }
func (*Pool) ProtoMessage() {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e991f749f68c2a4, []int{2}
}
func (pa *Pool) XXX_Unmarshal(b []byte) error {
	return pa.Unmarshal(b)
}
func (pa *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, pa, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := pa.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (pa *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(pa, src)
}
func (pa *Pool) XXX_Size() int {
	return pa.Size()
}
func (pa *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(pa)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SmoothWeightChangeParams)(nil), "osmosis.gamm.v1beta1.SmoothWeightChangeParams")
	proto.RegisterType((*PoolParams)(nil), "osmosis.gamm.v1beta1.PoolParams")
	proto.RegisterType((*Pool)(nil), "osmosis.gamm.v1beta1.Pool")
}

func init() {
	proto.RegisterFile("osmosis/gamm/pool-models/balancer/balancerPool.proto", fileDescriptor_7e991f749f68c2a4)
}

var fileDescriptor_7e991f749f68c2a4 = []byte{
	// 775 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0xc7, 0xe3, 0x24, 0xc0, 0x65, 0x90, 0xb8, 0x62, 0x2e, 0x0b, 0x13, 0x74, 0x63, 0xe4, 0x4a,
	0x15, 0xaa, 0x88, 0x2d, 0x68, 0xa5, 0x4a, 0x2c, 0x2a, 0x11, 0x68, 0x2b, 0x76, 0xd4, 0x54, 0x2a,
	0x2a, 0x0b, 0x6b, 0x92, 0x0c, 0x8e, 0x55, 0x3b, 0xe3, 0x7a, 0x26, 0x7c, 0xbc, 0x41, 0x97, 0xac,
	0x2a, 0x96, 0xbc, 0x42, 0xa5, 0xbe, 0x43, 0x59, 0xa2, 0xae, 0xaa, 0x2e, 0x4c, 0x05, 0xbb, 0x2e,
	0xf3, 0x04, 0xd5, 0xcc, 0x1c, 0x87, 0x34, 0x24, 0x12, 0x55, 0x57, 0xf1, 0xcc, 0x9c, 0xf3, 0x3b,
	0xff, 0xf3, 0x31, 0x13, 0xf4, 0x84, 0xf1, 0x98, 0xf1, 0x90, 0xbb, 0x01, 0x89, 0x63, 0x37, 0x61,
	0x2c, 0xaa, 0xc5, 0xac, 0x45, 0x23, 0xee, 0x36, 0x48, 0x44, 0x3a, 0x4d, 0x9a, 0xf6, 0x3f, 0x76,
	0x18, 0x8b, 0x9c, 0x24, 0x65, 0x82, 0xe1, 0x79, 0xf0, 0x72, 0xa4, 0x97, 0x73, 0xb8, 0xda, 0xa0,
	0x82, 0xac, 0x56, 0x16, 0x9a, 0x6a, 0xdb, 0x57, 0x36, 0xae, 0x5e, 0x68, 0x87, 0xca, 0x7c, 0xc0,
	0x02, 0xa6, 0xf7, 0xe5, 0x17, 0xec, 0x56, 0x03, 0xc6, 0x82, 0x88, 0xba, 0x6a, 0xd5, 0xe8, 0x1e,
	0xb8, 0xad, 0x6e, 0x4a, 0x44, 0xc8, 0x3a, 0x70, 0x6e, 0x0d, 0x9f, 0x8b, 0x30, 0xa6, 0x5c, 0x90,
	0x38, 0xc9, 0x01, 0x3a, 0x88, 0x4b, 0xba, 0xa2, 0xed, 0x82, 0x0c, 0xb5, 0x18, 0x3a, 0x6f, 0x10,
	0x4e, 0xfb, 0xe7, 0x4d, 0x16, 0xf6, 0x03, 0xfc, 0x96, 0x7d, 0x6e, 0x90, 0xf4, 0x13, 0xb5, 0xbf,
	0x94, 0x90, 0xb9, 0x1b, 0x33, 0x26, 0xda, 0x6f, 0x68, 0x18, 0xb4, 0xc5, 0x66, 0x9b, 0x74, 0x02,
	0xba, 0x43, 0x52, 0x12, 0x73, 0xbc, 0x87, 0x10, 0x17, 0x24, 0x15, 0xbe, 0x94, 0x65, 0x1a, 0x4b,
	0xc6, 0xf2, 0xcc, 0x5a, 0xc5, 0xd1, 0x9a, 0x9d, 0x5c, 0xb3, 0xf3, 0x3a, 0xd7, 0x5c, 0xff, 0xff,
	0x22, 0xb3, 0x0a, 0xbd, 0xcc, 0x9a, 0x3b, 0x21, 0x71, 0xb4, 0x6e, 0xdf, 0xfa, 0xda, 0xa7, 0x57,
	0x96, 0xe1, 0x4d, 0xab, 0x0d, 0x69, 0x8e, 0xdb, 0xe8, 0x9f, 0xbc, 0x14, 0x66, 0x51, 0x71, 0x17,
	0xee, 0x70, 0xb7, 0xc0, 0xa0, 0xbe, 0x2a, 0xb1, 0x3f, 0x33, 0x0b, 0xe7, 0x2e, 0x2b, 0x2c, 0x0e,
	0x05, 0x8d, 0x13, 0x71, 0xd2, 0xcb, 0xac, 0x7f, 0x75, 0xb0, 0xfc, 0xcc, 0x3e, 0x93, 0xa1, 0xfa,
	0x74, 0x2c, 0x10, 0x0e, 0x3b, 0xa1, 0x08, 0x49, 0x24, 0xdb, 0xab, 0x93, 0xe4, 0x66, 0x69, 0xa9,
	0xb4, 0x3c, 0xb3, 0x66, 0x39, 0xa3, 0xda, 0xec, 0x48, 0xc3, 0x0d, 0xce, 0xa9, 0xa8, 0x3f, 0x80,
	0x84, 0x16, 0x75, 0x0c, 0x00, 0xf9, 0xb2, 0x7e, 0xfe, 0x91, 0x46, 0xd9, 0xde, 0x08, 0x3e, 0x7e,
	0x8f, 0xe6, 0x04, 0x49, 0x03, 0x2a, 0x06, 0x83, 0x96, 0xef, 0x17, 0xd4, 0x86, 0xa0, 0x15, 0x1d,
	0x54, 0x73, 0x86, 0x62, 0xde, 0xa5, 0xdb, 0x57, 0x45, 0x84, 0xe4, 0x1a, 0x7a, 0xb7, 0x8f, 0xa6,
	0xf8, 0x11, 0x49, 0x5e, 0x50, 0xdd, 0xb8, 0xe9, 0xfa, 0x86, 0xc4, 0x7e, 0xcf, 0xac, 0x87, 0x41,
	0x28, 0xda, 0xdd, 0x86, 0xd3, 0x64, 0x31, 0x8c, 0x30, 0xfc, 0xd4, 0x78, 0xeb, 0x9d, 0x2b, 0x4e,
	0x12, 0xca, 0x9d, 0x2d, 0xda, 0xbc, 0xad, 0xac, 0xc4, 0xf8, 0x07, 0x94, 0xda, 0x5e, 0x4e, 0x94,
	0x70, 0x7a, 0x1c, 0x0a, 0x09, 0x2f, 0xfe, 0x1d, 0x5c, 0x62, 0x00, 0x0e, 0x44, 0xfc, 0xd1, 0x40,
	0x26, 0x1f, 0x33, 0x92, 0x66, 0x49, 0x0d, 0x8b, 0x33, 0xba, 0x86, 0xe3, 0x06, 0xb9, 0xfe, 0xe8,
	0x22, 0xb3, 0x8c, 0x5e, 0x66, 0xd9, 0x90, 0x91, 0xb2, 0x83, 0x6a, 0xfa, 0x4d, 0x65, 0xe9, 0x27,
	0xca, 0xd4, 0xf6, 0xc6, 0xc6, 0xb6, 0x3f, 0x95, 0x51, 0x59, 0x56, 0x18, 0xaf, 0xa0, 0x29, 0xd2,
	0x6a, 0xa5, 0x94, 0x73, 0xa8, 0x2d, 0xee, 0x65, 0xd6, 0xac, 0x66, 0xc3, 0x81, 0xed, 0xe5, 0x26,
	0x78, 0x16, 0x15, 0xc3, 0x96, 0xaa, 0x53, 0xd9, 0x2b, 0x86, 0x2d, 0x4c, 0x11, 0x4a, 0xfa, 0x7d,
	0x82, 0x84, 0x96, 0xc6, 0x0f, 0x05, 0xa4, 0x30, 0x34, 0x8a, 0xf9, 0x9b, 0xa5, 0xe7, 0x22, 0xd7,
	0x3e, 0x00, 0xc6, 0xaf, 0xd0, 0xfc, 0x41, 0x57, 0x74, 0x53, 0xaa, 0x4d, 0x02, 0x76, 0x48, 0xd3,
	0x0e, 0x4b, 0xcd, 0xb2, 0x52, 0x6c, 0xdd, 0xa2, 0x46, 0x59, 0xd9, 0x1e, 0xd6, 0xdb, 0x52, 0xc1,
	0x4b, 0xd8, 0xc4, 0x7b, 0x68, 0x46, 0x30, 0x41, 0xa2, 0xdd, 0x36, 0x49, 0x29, 0x37, 0x27, 0xe0,
	0xe2, 0xc2, 0x43, 0x28, 0xdf, 0xa0, 0xbe, 0xf2, 0x4d, 0x16, 0x76, 0xea, 0x8b, 0xa0, 0xf9, 0x3f,
	0x98, 0x64, 0xe9, 0xeb, 0x73, 0xe5, 0x6c, 0x7b, 0x83, 0x28, 0xbc, 0xaf, 0x6b, 0xa2, 0x2e, 0x00,
	0x37, 0x27, 0xef, 0x77, 0x51, 0x2a, 0x80, 0xc7, 0x1a, 0xaf, 0x12, 0x20, 0x8a, 0x00, 0x95, 0xd0,
	0x38, 0x1c, 0x80, 0x6c, 0xdd, 0x52, 0x73, 0x4a, 0x15, 0xe0, 0xf9, 0x1f, 0x4c, 0xec, 0x76, 0x47,
	0x0c, 0x67, 0xa1, 0x67, 0x27, 0xcf, 0x42, 0x93, 0xd7, 0xe7, 0x3e, 0x9c, 0x5b, 0x85, 0xb3, 0x73,
	0xab, 0xf0, 0xf5, 0x73, 0x6d, 0x42, 0xea, 0xdc, 0xae, 0xef, 0x5d, 0x5c, 0x57, 0x8d, 0xcb, 0xeb,
	0xaa, 0xf1, 0xe3, 0xba, 0x6a, 0x9c, 0xde, 0x54, 0x0b, 0x97, 0x37, 0xd5, 0xc2, 0xb7, 0x9b, 0x6a,
	0xe1, 0xed, 0xb3, 0x81, 0xc0, 0x90, 0x68, 0x2d, 0x22, 0x0d, 0x9e, 0x2f, 0xdc, 0xc3, 0xa7, 0xee,
	0xf1, 0xf8, 0x7f, 0xad, 0xc6, 0xa4, 0x7a, 0x28, 0x1f, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x8d,
	0xd7, 0xef, 0xcf, 0xe1, 0x06, 0x00, 0x00,
}

func (m *SmoothWeightChangeParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SmoothWeightChangeParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SmoothWeightChangeParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TargetPoolWeights) > 0 {
		for iNdEx := len(m.TargetPoolWeights) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TargetPoolWeights[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBalancerPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.InitialPoolWeights) > 0 {
		for iNdEx := len(m.InitialPoolWeights) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialPoolWeights[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBalancerPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintBalancerPool(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintBalancerPool(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PoolParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SmoothWeightChangeParams != nil {
		{
			size, err := m.SmoothWeightChangeParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBalancerPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.ExitFee.Size()
		i -= size
		if _, err := m.ExitFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBalancerPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.SwapFee.Size()
		i -= size
		if _, err := m.SwapFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBalancerPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (pa *Pool) Marshal() (dAtA []byte, err error) {
	size := pa.Size()
	dAtA = make([]byte, size)
	n, err := pa.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (pa *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := pa.Size()
	return pa.MarshalToSizedBuffer(dAtA[:size])
}

func (pa *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := pa.TotalWeight.Size()
		i -= size
		if _, err := pa.TotalWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBalancerPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(pa.PoolAssets) > 0 {
		for iNdEx := len(pa.PoolAssets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := pa.PoolAssets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBalancerPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size, err := pa.TotalShares.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBalancerPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(pa.FuturePoolGovernor) > 0 {
		i -= len(pa.FuturePoolGovernor)
		copy(dAtA[i:], pa.FuturePoolGovernor)
		i = encodeVarintBalancerPool(dAtA, i, uint64(len(pa.FuturePoolGovernor)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := pa.PoolParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBalancerPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if pa.Id != 0 {
		i = encodeVarintBalancerPool(dAtA, i, uint64(pa.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(pa.Address) > 0 {
		i -= len(pa.Address)
		copy(dAtA[i:], pa.Address)
		i = encodeVarintBalancerPool(dAtA, i, uint64(len(pa.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBalancerPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovBalancerPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SmoothWeightChangeParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovBalancerPool(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovBalancerPool(uint64(l))
	if len(m.InitialPoolWeights) > 0 {
		for _, e := range m.InitialPoolWeights {
			l = e.Size()
			n += 1 + l + sovBalancerPool(uint64(l))
		}
	}
	if len(m.TargetPoolWeights) > 0 {
		for _, e := range m.TargetPoolWeights {
			l = e.Size()
			n += 1 + l + sovBalancerPool(uint64(l))
		}
	}
	return n
}

func (m *PoolParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SwapFee.Size()
	n += 1 + l + sovBalancerPool(uint64(l))
	l = m.ExitFee.Size()
	n += 1 + l + sovBalancerPool(uint64(l))
	if m.SmoothWeightChangeParams != nil {
		l = m.SmoothWeightChangeParams.Size()
		n += 1 + l + sovBalancerPool(uint64(l))
	}
	return n
}

func (pa *Pool) Size() (n int) {
	if pa == nil {
		return 0
	}
	var l int
	_ = l
	l = len(pa.Address)
	if l > 0 {
		n += 1 + l + sovBalancerPool(uint64(l))
	}
	if pa.Id != 0 {
		n += 1 + sovBalancerPool(uint64(pa.Id))
	}
	l = pa.PoolParams.Size()
	n += 1 + l + sovBalancerPool(uint64(l))
	l = len(pa.FuturePoolGovernor)
	if l > 0 {
		n += 1 + l + sovBalancerPool(uint64(l))
	}
	l = pa.TotalShares.Size()
	n += 1 + l + sovBalancerPool(uint64(l))
	if len(pa.PoolAssets) > 0 {
		for _, e := range pa.PoolAssets {
			l = e.Size()
			n += 1 + l + sovBalancerPool(uint64(l))
		}
	}
	l = pa.TotalWeight.Size()
	n += 1 + l + sovBalancerPool(uint64(l))
	return n
}

func sovBalancerPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBalancerPool(x uint64) (n int) {
	return sovBalancerPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SmoothWeightChangeParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBalancerPool
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
			return fmt.Errorf("proto: SmoothWeightChangeParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SmoothWeightChangeParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialPoolWeights", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialPoolWeights = append(m.InitialPoolWeights, types.PoolAsset{})
			if err := m.InitialPoolWeights[len(m.InitialPoolWeights)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetPoolWeights", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TargetPoolWeights = append(m.TargetPoolWeights, types.PoolAsset{})
			if err := m.TargetPoolWeights[len(m.TargetPoolWeights)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBalancerPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBalancerPool
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
func (m *PoolParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBalancerPool
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
			return fmt.Errorf("proto: PoolParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExitFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SmoothWeightChangeParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SmoothWeightChangeParams == nil {
				m.SmoothWeightChangeParams = &SmoothWeightChangeParams{}
			}
			if err := m.SmoothWeightChangeParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBalancerPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBalancerPool
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
func (pa *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBalancerPool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			pa.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			pa.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				pa.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := pa.PoolParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FuturePoolGovernor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			pa.FuturePoolGovernor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := pa.TotalShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAssets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			pa.PoolAssets = append(pa.PoolAssets, types.PoolAsset{})
			if err := pa.PoolAssets[len(pa.PoolAssets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalancerPool
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
				return ErrInvalidLengthBalancerPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBalancerPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := pa.TotalWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBalancerPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBalancerPool
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
func skipBalancerPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBalancerPool
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
					return 0, ErrIntOverflowBalancerPool
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
					return 0, ErrIntOverflowBalancerPool
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
				return 0, ErrInvalidLengthBalancerPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBalancerPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBalancerPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBalancerPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBalancerPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBalancerPool = fmt.Errorf("proto: unexpected end of group")
)

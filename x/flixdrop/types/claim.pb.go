// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: omniflix/flixdrop/v1beta1/claim.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

type WeightedAction struct {
	Index     uint64                                 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty" yaml:"index"`
	Name      string                                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
	StartTime time.Time                              `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	Weight    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"weight" yaml:"weight"`
}

func (m *WeightedAction) Reset()         { *m = WeightedAction{} }
func (m *WeightedAction) String() string { return proto.CompactTextString(m) }
func (*WeightedAction) ProtoMessage()    {}
func (*WeightedAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_97edaf08aaa6941d, []int{0}
}
func (m *WeightedAction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WeightedAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WeightedAction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WeightedAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedAction.Merge(m, src)
}
func (m *WeightedAction) XXX_Size() int {
	return m.Size()
}
func (m *WeightedAction) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedAction.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedAction proto.InternalMessageInfo

func (m *WeightedAction) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *WeightedAction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WeightedAction) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

// A Claim Records is the metadata of claim data per address
type ClaimRecord struct {
	// address of claim user
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// total initial claimable amount for the user
	InitialClaimableAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=initial_claimable_amount,json=initialClaimableAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"initial_claimable_amount" yaml:"initial_claimable_amount"`
	// true if action is claimed
	// index of bool in array refers to action index #
	ActionsClaimed []bool `protobuf:"varint,3,rep,packed,name=actions_claimed,json=actionsClaimed,proto3" json:"actions_claimed,omitempty" yaml:"actions_claimed"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_97edaf08aaa6941d, []int{1}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func (m *ClaimRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClaimRecord) GetInitialClaimableAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.InitialClaimableAmount
	}
	return nil
}

func (m *ClaimRecord) GetActionsClaimed() []bool {
	if m != nil {
		return m.ActionsClaimed
	}
	return nil
}

func init() {
	proto.RegisterType((*WeightedAction)(nil), "omniflix.flixdrop.v1beta1.WeightedAction")
	proto.RegisterType((*ClaimRecord)(nil), "omniflix.flixdrop.v1beta1.ClaimRecord")
}

func init() {
	proto.RegisterFile("omniflix/flixdrop/v1beta1/claim.proto", fileDescriptor_97edaf08aaa6941d)
}

var fileDescriptor_97edaf08aaa6941d = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x93, 0x50, 0xc8, 0x06, 0x52, 0xb0, 0x50, 0xe5, 0x46, 0xc2, 0x1b, 0x19, 0x51, 0xf9,
	0x00, 0xbb, 0xb4, 0xdc, 0xb8, 0xa0, 0x3a, 0x88, 0x03, 0x17, 0x24, 0x83, 0x54, 0xc4, 0x25, 0x5a,
	0xdb, 0x5b, 0x67, 0x85, 0xd7, 0x1b, 0x65, 0x37, 0xd0, 0xfe, 0x45, 0x3e, 0x03, 0xc1, 0x8f, 0xf4,
	0xd8, 0x23, 0xe2, 0xe0, 0xa2, 0xe4, 0x0f, 0xfc, 0x05, 0x68, 0x77, 0x6d, 0xb5, 0x48, 0x1c, 0xb8,
	0x24, 0xf6, 0xcc, 0x9b, 0x37, 0x6f, 0xe6, 0x8d, 0xc1, 0x13, 0xc1, 0x4b, 0x76, 0x5a, 0xb0, 0x33,
	0xac, 0x7f, 0xb2, 0xa5, 0x58, 0xe0, 0x2f, 0x87, 0x09, 0x55, 0xe4, 0x10, 0xa7, 0x05, 0x61, 0x1c,
	0x2d, 0x96, 0x42, 0x09, 0x77, 0xbf, 0x85, 0xa1, 0x16, 0x86, 0x1a, 0xd8, 0xf8, 0x61, 0x2e, 0x72,
	0x61, 0x50, 0x58, 0x3f, 0xd9, 0x82, 0xb1, 0x9f, 0x0a, 0xc9, 0x85, 0xc4, 0x09, 0x91, 0xf4, 0x9a,
	0x51, 0xb0, 0xb2, 0xc9, 0xc3, 0x5c, 0x88, 0xbc, 0xa0, 0xd8, 0xbc, 0x25, 0xab, 0x53, 0xac, 0x18,
	0xa7, 0x52, 0x11, 0xbe, 0xb0, 0x80, 0x60, 0xdd, 0x05, 0xa3, 0x13, 0xca, 0xf2, 0xb9, 0xa2, 0xd9,
	0x71, 0xaa, 0x98, 0x28, 0xdd, 0x03, 0x70, 0x8b, 0x95, 0x19, 0x3d, 0xf3, 0x9c, 0x89, 0x13, 0xf6,
	0xa3, 0xfb, 0x75, 0x05, 0xef, 0x9e, 0x13, 0x5e, 0xbc, 0x0c, 0x4c, 0x38, 0x88, 0x6d, 0xda, 0x7d,
	0x0c, 0xfa, 0x25, 0xe1, 0xd4, 0xeb, 0x4e, 0x9c, 0x70, 0x10, 0xed, 0xd6, 0x15, 0x1c, 0x5a, 0x98,
	0x8e, 0x06, 0xb1, 0x49, 0xba, 0x1f, 0x01, 0x90, 0x8a, 0x2c, 0xd5, 0x4c, 0x37, 0xf6, 0x7a, 0x13,
	0x27, 0x1c, 0x1e, 0x8d, 0x91, 0x55, 0x85, 0x5a, 0x55, 0xe8, 0x43, 0xab, 0x2a, 0x7a, 0x74, 0x51,
	0xc1, 0x4e, 0x5d, 0xc1, 0x07, 0x96, 0xea, 0xba, 0x36, 0x58, 0x5f, 0x41, 0x27, 0x1e, 0x98, 0x80,
	0x86, 0xbb, 0x27, 0x60, 0xe7, 0xab, 0x11, 0xee, 0xf5, 0x8d, 0x80, 0x57, 0xba, 0xf2, 0x57, 0x05,
	0x0f, 0x72, 0xa6, 0xe6, 0xab, 0x04, 0xa5, 0x82, 0xe3, 0x66, 0x3b, 0xf6, 0xef, 0x99, 0xcc, 0x3e,
	0x63, 0x75, 0xbe, 0xa0, 0x12, 0xbd, 0xa6, 0x69, 0x5d, 0xc1, 0x7b, 0xb6, 0x87, 0x65, 0x09, 0xe2,
	0x86, 0x2e, 0xf8, 0xd1, 0x05, 0xc3, 0xa9, 0x36, 0x25, 0xa6, 0xa9, 0x58, 0x66, 0xee, 0x53, 0x70,
	0x9b, 0x64, 0xd9, 0x92, 0x4a, 0x69, 0x36, 0x32, 0x88, 0xdc, 0xba, 0x82, 0x23, 0x5b, 0xdb, 0x24,
	0x82, 0xb8, 0x85, 0xb8, 0xdf, 0x1c, 0xe0, 0xb1, 0x92, 0x29, 0x46, 0x8a, 0x99, 0xb1, 0x96, 0x24,
	0x05, 0x9d, 0x11, 0x2e, 0x56, 0xa5, 0xf2, 0xba, 0x93, 0x5e, 0x38, 0x3c, 0xda, 0x47, 0x56, 0x10,
	0xd2, 0xae, 0xb5, 0x06, 0xa3, 0xa9, 0x60, 0x65, 0xf4, 0xbe, 0x19, 0x1f, 0xb6, 0x0b, 0xff, 0x37,
	0x51, 0xf0, 0xfd, 0x0a, 0x86, 0xff, 0x31, 0xa7, 0xe6, 0x94, 0xf1, 0x5e, 0x43, 0x33, 0x6d, 0x59,
	0x8e, 0x0d, 0x89, 0x3b, 0x05, 0xbb, 0xc4, 0x58, 0x2e, 0x6d, 0x03, 0x9a, 0x79, 0xbd, 0x49, 0x2f,
	0xbc, 0x13, 0x8d, 0xeb, 0x0a, 0xee, 0x35, 0x03, 0xfe, 0x0d, 0x08, 0xe2, 0x51, 0x13, 0x99, 0xda,
	0x40, 0xf4, 0xf6, 0x62, 0xe3, 0x3b, 0x97, 0x1b, 0xdf, 0xf9, 0xbd, 0xf1, 0x9d, 0xf5, 0xd6, 0xef,
	0x5c, 0x6e, 0xfd, 0xce, 0xcf, 0xad, 0xdf, 0xf9, 0xf4, 0xfc, 0x86, 0xc0, 0x77, 0xbc, 0x64, 0x6f,
	0xf4, 0xf9, 0xb7, 0x07, 0x3e, 0x5f, 0x25, 0xf8, 0xc6, 0xc7, 0x60, 0xe4, 0x26, 0x3b, 0xe6, 0x20,
	0x5e, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x08, 0xa1, 0x7c, 0x2e, 0x03, 0x00, 0x00,
}

func (m *WeightedAction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WeightedAction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WeightedAction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClaim(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintClaim(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Index != 0 {
		i = encodeVarintClaim(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionsClaimed) > 0 {
		for iNdEx := len(m.ActionsClaimed) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.ActionsClaimed[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintClaim(dAtA, i, uint64(len(m.ActionsClaimed)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InitialClaimableAmount) > 0 {
		for iNdEx := len(m.InitialClaimableAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialClaimableAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaim(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaim(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WeightedAction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovClaim(uint64(m.Index))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovClaim(uint64(l))
	l = m.Weight.Size()
	n += 1 + l + sovClaim(uint64(l))
	return n
}

func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	if len(m.InitialClaimableAmount) > 0 {
		for _, e := range m.InitialClaimableAmount {
			l = e.Size()
			n += 1 + l + sovClaim(uint64(l))
		}
	}
	if len(m.ActionsClaimed) > 0 {
		n += 1 + sovClaim(uint64(len(m.ActionsClaimed))) + len(m.ActionsClaimed)*1
	}
	return n
}

func sovClaim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaim(x uint64) (n int) {
	return sovClaim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WeightedAction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
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
			return fmt.Errorf("proto: WeightedAction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WeightedAction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
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
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
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
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClaimableAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialClaimableAmount = append(m.InitialClaimableAmount, types.Coin{})
			if err := m.InitialClaimableAmount[len(m.InitialClaimableAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ActionsClaimed = append(m.ActionsClaimed, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaim
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaim
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.ActionsClaimed) == 0 {
					m.ActionsClaimed = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaim
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ActionsClaimed = append(m.ActionsClaimed, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionsClaimed", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
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
func skipClaim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
				return 0, ErrInvalidLengthClaim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaim = fmt.Errorf("proto: unexpected end of group")
)

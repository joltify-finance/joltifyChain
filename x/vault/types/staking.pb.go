// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vault/staking.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Params struct {
	BlockChurnInterval int64                                  `protobuf:"varint,1,opt,name=block_churn_interval,json=blockChurnInterval,proto3" json:"block_churn_interval,omitempty" yaml:"block_churn_interval"`
	Power              int64                                  `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty" yaml:"power"`
	Step               int64                                  `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty" yaml:"step"`
	CandidateRatio     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=candidate_ratio,json=candidateRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"candidate_ratio" yaml:"candidate_ratio"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_f70c1fe68a07899e, []int{0}
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

func (m *Params) GetBlockChurnInterval() int64 {
	if m != nil {
		return m.BlockChurnInterval
	}
	return 0
}

func (m *Params) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

func (m *Params) GetStep() int64 {
	if m != nil {
		return m.Step
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "joltify.joltifychain.vault.Params")
}

func init() { proto.RegisterFile("vault/staking.proto", fileDescriptor_f70c1fe68a07899e) }

var fileDescriptor_f70c1fe68a07899e = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4e, 0xeb, 0x30,
	0x18, 0x85, 0xe3, 0xb6, 0xb7, 0xba, 0x04, 0x44, 0x51, 0xa8, 0x50, 0x54, 0xa4, 0xb8, 0x0a, 0x52,
	0xd5, 0xa5, 0xc9, 0xc0, 0xd6, 0xb1, 0x45, 0x02, 0x36, 0xc8, 0xc8, 0x52, 0xb9, 0x6e, 0x48, 0x4d,
	0x93, 0x38, 0xc4, 0x6e, 0xa1, 0x6f, 0xc1, 0xc8, 0xd8, 0x85, 0x77, 0xe9, 0xd8, 0x11, 0x31, 0x58,
	0xa8, 0x59, 0x98, 0xf3, 0x04, 0x28, 0x76, 0x40, 0x08, 0x31, 0x1d, 0xeb, 0xfb, 0x8f, 0x8f, 0xad,
	0xf3, 0xeb, 0x87, 0x0b, 0x34, 0x0f, 0xb9, 0xcb, 0x38, 0x9a, 0x91, 0x38, 0x70, 0x92, 0x94, 0x72,
	0x6a, 0xb4, 0xee, 0x68, 0xc8, 0xc9, 0xed, 0xd2, 0x29, 0x15, 0x4f, 0x11, 0x89, 0x1d, 0xe9, 0x6c,
	0x35, 0x03, 0x1a, 0x50, 0x69, 0x73, 0x8b, 0x93, 0xba, 0x61, 0xbf, 0x54, 0xf4, 0xfa, 0x15, 0x4a,
	0x51, 0xc4, 0x8c, 0x6b, 0xbd, 0x39, 0x0e, 0x29, 0x9e, 0x8d, 0xf0, 0x74, 0x9e, 0xc6, 0x23, 0x12,
	0x73, 0x3f, 0x5d, 0xa0, 0xd0, 0x04, 0x6d, 0xd0, 0xad, 0x0e, 0x60, 0x2e, 0xe0, 0xf1, 0x12, 0x45,
	0x61, 0xdf, 0xfe, 0xcb, 0x65, 0x7b, 0x86, 0xc4, 0xc3, 0x82, 0x5e, 0x96, 0xd0, 0xe8, 0xe8, 0xff,
	0x12, 0xfa, 0xe0, 0xa7, 0x66, 0x45, 0x66, 0x1c, 0xe4, 0x02, 0xee, 0xa9, 0x0c, 0x89, 0x6d, 0x4f,
	0x8d, 0x8d, 0x13, 0xbd, 0xc6, 0xb8, 0x9f, 0x98, 0x55, 0x69, 0x6b, 0xe4, 0x02, 0xee, 0x2a, 0x5b,
	0x41, 0x6d, 0x4f, 0x0e, 0x8d, 0x7b, 0xbd, 0x81, 0x51, 0x3c, 0x21, 0x13, 0xc4, 0xfd, 0x51, 0x8a,
	0x38, 0xa1, 0x66, 0xad, 0x0d, 0xba, 0x3b, 0x83, 0x8b, 0xb5, 0x80, 0xda, 0x9b, 0x80, 0x9d, 0x80,
	0xf0, 0xe9, 0x7c, 0xec, 0x60, 0x1a, 0xb9, 0x98, 0xb2, 0x88, 0xb2, 0x52, 0x7a, 0x6c, 0x32, 0x73,
	0xf9, 0x32, 0xf1, 0x99, 0x73, 0xe6, 0xe3, 0x5c, 0xc0, 0x23, 0x95, 0xfe, 0x2b, 0xce, 0xf6, 0xf6,
	0xbf, 0x89, 0x57, 0x80, 0xfe, 0xff, 0xe7, 0x15, 0xd4, 0x3e, 0x56, 0x10, 0x0c, 0xce, 0xd7, 0x5b,
	0x0b, 0x6c, 0xb6, 0x16, 0x78, 0xdf, 0x5a, 0xe0, 0x29, 0xb3, 0xb4, 0x4d, 0x66, 0x69, 0xaf, 0x99,
	0xa5, 0xdd, 0xf4, 0x7e, 0xbc, 0x5a, 0xd6, 0xfe, 0xa5, 0xc3, 0xa2, 0x7e, 0xf7, 0xd1, 0x55, 0xab,
	0x92, 0x1f, 0x18, 0xd7, 0x65, 0xef, 0xa7, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xc9, 0x6a,
	0x5c, 0xc0, 0x01, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.BlockChurnInterval != that1.BlockChurnInterval {
		return false
	}
	if this.Power != that1.Power {
		return false
	}
	if this.Step != that1.Step {
		return false
	}
	if !this.CandidateRatio.Equal(that1.CandidateRatio) {
		return false
	}
	return true
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
	{
		size := m.CandidateRatio.Size()
		i -= size
		if _, err := m.CandidateRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Step != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.Step))
		i--
		dAtA[i] = 0x18
	}
	if m.Power != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x10
	}
	if m.BlockChurnInterval != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.BlockChurnInterval))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStaking(dAtA []byte, offset int, v uint64) int {
	offset -= sovStaking(v)
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
	if m.BlockChurnInterval != 0 {
		n += 1 + sovStaking(uint64(m.BlockChurnInterval))
	}
	if m.Power != 0 {
		n += 1 + sovStaking(uint64(m.Power))
	}
	if m.Step != 0 {
		n += 1 + sovStaking(uint64(m.Step))
	}
	l = m.CandidateRatio.Size()
	n += 1 + l + sovStaking(uint64(l))
	return n
}

func sovStaking(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStaking(x uint64) (n int) {
	return sovStaking(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
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
				return fmt.Errorf("proto: wrong wireType = %d for field BlockChurnInterval", wireType)
			}
			m.BlockChurnInterval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockChurnInterval |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			m.Step = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Step |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CandidateRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CandidateRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStaking
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
func skipStaking(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStaking
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
					return 0, ErrIntOverflowStaking
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
					return 0, ErrIntOverflowStaking
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
				return 0, ErrInvalidLengthStaking
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStaking
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStaking
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStaking        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStaking          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStaking = fmt.Errorf("proto: unexpected end of group")
)

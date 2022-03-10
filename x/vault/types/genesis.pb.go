// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vault/genesis.proto

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

// GenesisState defines the vault module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of related to deposit.
	Params         Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	OutboundTxList []OutboundTx `protobuf:"bytes,5,rep,name=outboundTxList,proto3" json:"outboundTxList"`
	// this line is used by starport scaffolding # genesis/proto/state
	IssueTokenList []*IssueToken `protobuf:"bytes,2,rep,name=issueTokenList,proto3" json:"issueTokenList,omitempty"`
	CreatePoolList []*CreatePool `protobuf:"bytes,3,rep,name=createPoolList,proto3" json:"createPoolList,omitempty"`
	Exported       bool          `protobuf:"varint,4,opt,name=exported,proto3" json:"exported,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c6e91fd07fc0e40, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetOutboundTxList() []OutboundTx {
	if m != nil {
		return m.OutboundTxList
	}
	return nil
}

func (m *GenesisState) GetIssueTokenList() []*IssueToken {
	if m != nil {
		return m.IssueTokenList
	}
	return nil
}

func (m *GenesisState) GetCreatePoolList() []*CreatePool {
	if m != nil {
		return m.CreatePoolList
	}
	return nil
}

func (m *GenesisState) GetExported() bool {
	if m != nil {
		return m.Exported
	}
	return false
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "joltify.joltifychain.vault.GenesisState")
}

func init() { proto.RegisterFile("vault/genesis.proto", fileDescriptor_3c6e91fd07fc0e40) }

var fileDescriptor_3c6e91fd07fc0e40 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x40, 0x5b, 0x40, 0x42, 0x8a, 0xe1, 0x50, 0x4d, 0x24, 0x1c, 0x2a, 0xe1, 0x60, 0xb8, 0xb8,
	0x9b, 0xe0, 0x0f, 0x18, 0x3c, 0x10, 0x13, 0xa3, 0xa4, 0x72, 0xf2, 0x42, 0x16, 0x58, 0xeb, 0x4a,
	0xe9, 0x34, 0xdd, 0xa9, 0x29, 0x7f, 0xe1, 0x67, 0x71, 0xe4, 0xe8, 0xc9, 0x98, 0xf2, 0x23, 0x86,
	0xdd, 0xa5, 0x51, 0x12, 0xe5, 0xd4, 0xed, 0xce, 0xbc, 0x37, 0xb3, 0x33, 0xce, 0xc9, 0x1b, 0x4b,
	0x43, 0xa4, 0x01, 0x8f, 0xb8, 0x14, 0x92, 0xc4, 0x09, 0x20, 0xb8, 0xad, 0x57, 0x08, 0x51, 0x3c,
	0x2f, 0x89, 0xf9, 0x4e, 0x5f, 0x98, 0x88, 0x88, 0xca, 0x6c, 0x9d, 0x69, 0x00, 0x52, 0x9c, 0x40,
	0x1a, 0xcd, 0xc6, 0x98, 0x69, 0x68, 0x17, 0x10, 0x52, 0xa6, 0x7c, 0x8c, 0x30, 0xe7, 0xd1, 0xef,
	0xc0, 0x34, 0xe1, 0x0c, 0xf9, 0x38, 0x06, 0x08, 0x4d, 0xc0, 0xd4, 0x96, 0xc8, 0xe6, 0x22, 0x0a,
	0xcc, 0xe5, 0x69, 0x00, 0x01, 0xa8, 0x23, 0xdd, 0x9e, 0xf4, 0x6d, 0x27, 0x2f, 0x39, 0xc7, 0x03,
	0xdd, 0xe3, 0x23, 0x32, 0xe4, 0xee, 0xb5, 0x53, 0x8d, 0x59, 0xc2, 0x16, 0xb2, 0x69, 0xb7, 0xed,
	0x6e, 0xbd, 0xd7, 0x21, 0x7f, 0xf7, 0x4c, 0x86, 0x2a, 0xb3, 0x5f, 0x59, 0x7d, 0x9e, 0x5b, 0xbe,
	0xe1, 0xdc, 0x91, 0xd3, 0xd8, 0x3d, 0x62, 0x94, 0xdd, 0x09, 0x89, 0xcd, 0xa3, 0x76, 0xb9, 0x5b,
	0xef, 0x5d, 0xfc, 0x67, 0x7a, 0x28, 0x08, 0x63, 0xdb, 0x73, 0xb8, 0xf7, 0x4e, 0x43, 0x4d, 0x60,
	0xb4, 0x1d, 0x80, 0xb2, 0x96, 0x0e, 0x5b, 0x6f, 0x0b, 0xc2, 0xdf, 0xa3, 0xb7, 0x3e, 0x3d, 0xb8,
	0x21, 0x40, 0xa8, 0x7c, 0xe5, 0xc3, 0xbe, 0x9b, 0x82, 0xf0, 0xf7, 0x68, 0xb7, 0xe5, 0xd4, 0x78,
	0x16, 0x43, 0x82, 0x7c, 0xd6, 0xac, 0xb4, 0xed, 0x6e, 0xcd, 0x2f, 0xfe, 0xfb, 0x83, 0x55, 0xee,
	0xd9, 0xeb, 0xdc, 0xb3, 0xbf, 0x72, 0xcf, 0x7e, 0xdf, 0x78, 0xd6, 0x7a, 0xe3, 0x59, 0x1f, 0x1b,
	0xcf, 0x7a, 0xba, 0x0c, 0x04, 0x86, 0x6c, 0x42, 0xa6, 0xb0, 0xa0, 0xa6, 0x1e, 0xfd, 0x59, 0x97,
	0x66, 0x54, 0xef, 0x12, 0x97, 0x31, 0x97, 0x93, 0xaa, 0x5a, 0xda, 0xd5, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd3, 0x92, 0x45, 0xde, 0x5d, 0x02, 0x00, 0x00,
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
	if len(m.OutboundTxList) > 0 {
		for iNdEx := len(m.OutboundTxList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OutboundTxList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.Exported {
		i--
		if m.Exported {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.CreatePoolList) > 0 {
		for iNdEx := len(m.CreatePoolList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CreatePoolList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.IssueTokenList) > 0 {
		for iNdEx := len(m.IssueTokenList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IssueTokenList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.IssueTokenList) > 0 {
		for _, e := range m.IssueTokenList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CreatePoolList) > 0 {
		for _, e := range m.CreatePoolList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.Exported {
		n += 2
	}
	if len(m.OutboundTxList) > 0 {
		for _, e := range m.OutboundTxList {
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
				return fmt.Errorf("proto: wrong wireType = %d for field IssueTokenList", wireType)
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
			m.IssueTokenList = append(m.IssueTokenList, &IssueToken{})
			if err := m.IssueTokenList[len(m.IssueTokenList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatePoolList", wireType)
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
			m.CreatePoolList = append(m.CreatePoolList, &CreatePool{})
			if err := m.CreatePoolList[len(m.CreatePoolList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exported", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.Exported = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutboundTxList", wireType)
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
			m.OutboundTxList = append(m.OutboundTxList, OutboundTx{})
			if err := m.OutboundTxList[len(m.OutboundTxList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

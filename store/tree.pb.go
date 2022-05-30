// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: store/v1beta1/tree.proto

package store

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
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

type Node struct {
	Children []*Child `protobuf:"bytes,1,rep,name=children,proto3" json:"children,omitempty"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_a24990997136ffd0, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Node.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return m.Size()
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetChildren() []*Child {
	if m != nil {
		return m.Children
	}
	return nil
}

type Child struct {
	Index        []byte                                 `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Accumulation github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=accumulation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"accumulation"`
}

func (m *Child) Reset()         { *m = Child{} }
func (m *Child) String() string { return proto.CompactTextString(m) }
func (*Child) ProtoMessage()    {}
func (*Child) Descriptor() ([]byte, []int) {
	return fileDescriptor_a24990997136ffd0, []int{1}
}
func (m *Child) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Child) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Child.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Child) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Child.Merge(m, src)
}
func (m *Child) XXX_Size() int {
	return m.Size()
}
func (m *Child) XXX_DiscardUnknown() {
	xxx_messageInfo_Child.DiscardUnknown(m)
}

var xxx_messageInfo_Child proto.InternalMessageInfo

func (m *Child) GetIndex() []byte {
	if m != nil {
		return m.Index
	}
	return nil
}

type Leaf struct {
	Leaf *Child `protobuf:"bytes,1,opt,name=leaf,proto3" json:"leaf,omitempty"`
}

func (m *Leaf) Reset()         { *m = Leaf{} }
func (m *Leaf) String() string { return proto.CompactTextString(m) }
func (*Leaf) ProtoMessage()    {}
func (*Leaf) Descriptor() ([]byte, []int) {
	return fileDescriptor_a24990997136ffd0, []int{2}
}
func (m *Leaf) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Leaf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Leaf.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Leaf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Leaf.Merge(m, src)
}
func (m *Leaf) XXX_Size() int {
	return m.Size()
}
func (m *Leaf) XXX_DiscardUnknown() {
	xxx_messageInfo_Leaf.DiscardUnknown(m)
}

var xxx_messageInfo_Leaf proto.InternalMessageInfo

func (m *Leaf) GetLeaf() *Child {
	if m != nil {
		return m.Leaf
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "joltify.joltifychain.store.v1beta1.Node")
	proto.RegisterType((*Child)(nil), "joltify.joltifychain.store.v1beta1.Child")
	proto.RegisterType((*Leaf)(nil), "joltify.joltifychain.store.v1beta1.Leaf")
}

func init() { proto.RegisterFile("store/v1beta1/tree.proto", fileDescriptor_a24990997136ffd0) }

var fileDescriptor_a24990997136ffd0 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x63, 0x68, 0x11, 0x98, 0x4e, 0x51, 0x87, 0xa8, 0x83, 0x5b, 0x45, 0x02, 0x95, 0x01,
	0x5b, 0x85, 0x19, 0x21, 0x15, 0x75, 0x40, 0x02, 0x86, 0x8c, 0x6c, 0x8e, 0xe3, 0xa6, 0x86, 0xd4,
	0xb7, 0xc4, 0x0e, 0xa2, 0x6f, 0xc1, 0x63, 0x75, 0xec, 0x88, 0x18, 0x2a, 0x94, 0xbc, 0x08, 0xaa,
	0x13, 0xf1, 0x33, 0x21, 0xa6, 0xeb, 0x9f, 0x73, 0xbe, 0x73, 0x74, 0x71, 0x60, 0x2c, 0xe4, 0x92,
	0x3d, 0x8f, 0x62, 0x69, 0xf9, 0x88, 0xd9, 0x5c, 0x4a, 0xba, 0xc8, 0xc1, 0x82, 0x1f, 0x3e, 0x40,
	0x66, 0xd5, 0x74, 0x49, 0x9b, 0x29, 0x66, 0x5c, 0x69, 0xea, 0xe4, 0xb4, 0x91, 0xf7, 0xba, 0x29,
	0xa4, 0xe0, 0xe4, 0x6c, 0x7b, 0xaa, 0x9d, 0x3d, 0x22, 0xc0, 0xcc, 0xc1, 0xb0, 0x98, 0x9b, 0x6f,
	0xb2, 0x00, 0xa5, 0xeb, 0xff, 0xf0, 0x16, 0xb7, 0xee, 0x20, 0x91, 0xfe, 0x04, 0xef, 0x8b, 0x99,
	0xca, 0x92, 0x5c, 0xea, 0x00, 0x0d, 0x76, 0x87, 0x87, 0x67, 0x27, 0xf4, 0xef, 0x50, 0x7a, 0xb5,
	0xf5, 0x44, 0x5f, 0xd6, 0xf0, 0x09, 0xb7, 0xdd, 0x93, 0xdf, 0xc5, 0x6d, 0xa5, 0x13, 0xf9, 0x12,
	0xa0, 0x01, 0x1a, 0x76, 0xa2, 0xfa, 0xe2, 0x47, 0xb8, 0xc3, 0x85, 0x28, 0xe6, 0x45, 0xc6, 0xad,
	0x02, 0x1d, 0xec, 0x0c, 0xd0, 0xf0, 0x60, 0x4c, 0x57, 0x9b, 0xbe, 0xf7, 0xbe, 0xe9, 0x1f, 0xa7,
	0xca, 0xce, 0x8a, 0x98, 0x0a, 0x98, 0xb3, 0xa6, 0x76, 0x3d, 0x4e, 0x4d, 0xf2, 0xc8, 0xec, 0x72,
	0x21, 0x0d, 0xbd, 0xd6, 0x36, 0xfa, 0xc5, 0x08, 0x27, 0xb8, 0x75, 0x23, 0xf9, 0xd4, 0xbf, 0xc0,
	0xad, 0x4c, 0xf2, 0xa9, 0x0b, 0xfc, 0x57, 0x7b, 0x67, 0x1b, 0x5f, 0xae, 0x4a, 0x82, 0xd6, 0x25,
	0x41, 0x1f, 0x25, 0x41, 0xaf, 0x15, 0xf1, 0xd6, 0x15, 0xf1, 0xde, 0x2a, 0xe2, 0xdd, 0x1f, 0xa5,
	0xca, 0x66, 0xbc, 0xae, 0xd5, 0xc0, 0xd8, 0x4f, 0x28, 0x73, 0xd0, 0x78, 0xcf, 0x2d, 0xf4, 0xfc,
	0x33, 0x00, 0x00, 0xff, 0xff, 0x35, 0xa8, 0x04, 0x46, 0xc6, 0x01, 0x00, 0x00,
}

func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Node) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Children) > 0 {
		for iNdEx := len(m.Children) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Children[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTree(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Child) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Child) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Child) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Accumulation.Size()
		i -= size
		if _, err := m.Accumulation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTree(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTree(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Leaf) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Leaf) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Leaf) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Leaf != nil {
		{
			size, err := m.Leaf.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTree(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTree(dAtA []byte, offset int, v uint64) int {
	offset -= sovTree(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Node) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Children) > 0 {
		for _, e := range m.Children {
			l = e.Size()
			n += 1 + l + sovTree(uint64(l))
		}
	}
	return n
}

func (m *Child) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTree(uint64(l))
	}
	l = m.Accumulation.Size()
	n += 1 + l + sovTree(uint64(l))
	return n
}

func (m *Leaf) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Leaf != nil {
		l = m.Leaf.Size()
		n += 1 + l + sovTree(uint64(l))
	}
	return n
}

func sovTree(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTree(x uint64) (n int) {
	return sovTree(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTree
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
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Children", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTree
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
				return ErrInvalidLengthTree
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Children = append(m.Children, &Child{})
			if err := m.Children[len(m.Children)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTree
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
func (m *Child) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTree
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
			return fmt.Errorf("proto: Child: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Child: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTree
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTree
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = append(m.Index[:0], dAtA[iNdEx:postIndex]...)
			if m.Index == nil {
				m.Index = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accumulation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTree
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
				return ErrInvalidLengthTree
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Accumulation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTree
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
func (m *Leaf) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTree
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
			return fmt.Errorf("proto: Leaf: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Leaf: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leaf", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTree
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
				return ErrInvalidLengthTree
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTree
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Leaf == nil {
				m.Leaf = &Child{}
			}
			if err := m.Leaf.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTree(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTree
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
func skipTree(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTree
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
					return 0, ErrIntOverflowTree
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
					return 0, ErrIntOverflowTree
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
				return 0, ErrInvalidLengthTree
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTree
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTree
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTree        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTree          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTree = fmt.Errorf("proto: unexpected end of group")
)

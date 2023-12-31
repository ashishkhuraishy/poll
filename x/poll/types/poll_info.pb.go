// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poll/poll_info.proto

package types

import (
	fmt "fmt"
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

type PollInfo struct {
	NexiId uint64 `protobuf:"varint,1,opt,name=nexiId,proto3" json:"nexiId,omitempty"`
}

func (m *PollInfo) Reset()         { *m = PollInfo{} }
func (m *PollInfo) String() string { return proto.CompactTextString(m) }
func (*PollInfo) ProtoMessage()    {}
func (*PollInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_49d30bde0d850159, []int{0}
}
func (m *PollInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PollInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PollInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PollInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollInfo.Merge(m, src)
}
func (m *PollInfo) XXX_Size() int {
	return m.Size()
}
func (m *PollInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PollInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PollInfo proto.InternalMessageInfo

func (m *PollInfo) GetNexiId() uint64 {
	if m != nil {
		return m.NexiId
	}
	return 0
}

func init() {
	proto.RegisterType((*PollInfo)(nil), "ashishkhuraishy.poll.poll.PollInfo")
}

func init() { proto.RegisterFile("poll/poll_info.proto", fileDescriptor_49d30bde0d850159) }

var fileDescriptor_49d30bde0d850159 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0xc8, 0xcf, 0xc9,
	0xd1, 0x07, 0x11, 0xf1, 0x99, 0x79, 0x69, 0xf9, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x92,
	0x89, 0xc5, 0x19, 0x99, 0xc5, 0x19, 0xd9, 0x19, 0xa5, 0x45, 0x89, 0x99, 0xc5, 0x19, 0x95, 0x7a,
	0x20, 0x05, 0x60, 0x42, 0x49, 0x89, 0x8b, 0x23, 0x20, 0x3f, 0x27, 0xc7, 0x33, 0x2f, 0x2d, 0x5f,
	0x48, 0x8c, 0x8b, 0x2d, 0x2f, 0xb5, 0x22, 0xd3, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x25,
	0x08, 0xca, 0x73, 0x72, 0x3b, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4,
	0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x9d,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x34, 0x3b, 0xc0, 0x8e, 0xd0,
	0xaf, 0x80, 0x50, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xd7, 0x18, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x8f, 0xdb, 0x05, 0x48, 0xa5, 0x00, 0x00, 0x00,
}

func (m *PollInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PollInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PollInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NexiId != 0 {
		i = encodeVarintPollInfo(dAtA, i, uint64(m.NexiId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPollInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovPollInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PollInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NexiId != 0 {
		n += 1 + sovPollInfo(uint64(m.NexiId))
	}
	return n
}

func sovPollInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPollInfo(x uint64) (n int) {
	return sovPollInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PollInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPollInfo
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
			return fmt.Errorf("proto: PollInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PollInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NexiId", wireType)
			}
			m.NexiId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPollInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NexiId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPollInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPollInfo
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
func skipPollInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPollInfo
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
					return 0, ErrIntOverflowPollInfo
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
					return 0, ErrIntOverflowPollInfo
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
				return 0, ErrInvalidLengthPollInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPollInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPollInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPollInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPollInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPollInfo = fmt.Errorf("proto: unexpected end of group")
)

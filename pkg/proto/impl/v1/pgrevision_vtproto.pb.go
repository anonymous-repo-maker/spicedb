// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.5.1-0.20231212170721-e7d721933795
// source: impl/v1/pgrevision.proto

package implv1

import (
	fmt "fmt"
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *PostgresRevision) CloneVT() *PostgresRevision {
	if m == nil {
		return (*PostgresRevision)(nil)
	}
	r := new(PostgresRevision)
	r.Xmin = m.Xmin
	r.RelativeXmax = m.RelativeXmax
	r.OptionalTxid = m.OptionalTxid
	r.OptionalTimestamp = m.OptionalTimestamp
	if rhs := m.RelativeXips; rhs != nil {
		tmpContainer := make([]int64, len(rhs))
		copy(tmpContainer, rhs)
		r.RelativeXips = tmpContainer
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *PostgresRevision) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *PostgresRevision) EqualVT(that *PostgresRevision) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if this.Xmin != that.Xmin {
		return false
	}
	if this.RelativeXmax != that.RelativeXmax {
		return false
	}
	if len(this.RelativeXips) != len(that.RelativeXips) {
		return false
	}
	for i, vx := range this.RelativeXips {
		vy := that.RelativeXips[i]
		if vx != vy {
			return false
		}
	}
	if this.OptionalTxid != that.OptionalTxid {
		return false
	}
	if this.OptionalTimestamp != that.OptionalTimestamp {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *PostgresRevision) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*PostgresRevision)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *PostgresRevision) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostgresRevision) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *PostgresRevision) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.OptionalTimestamp != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.OptionalTimestamp))
		i--
		dAtA[i] = 0x28
	}
	if m.OptionalTxid != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.OptionalTxid))
		i--
		dAtA[i] = 0x20
	}
	if len(m.RelativeXips) > 0 {
		var pksize2 int
		for _, num := range m.RelativeXips {
			pksize2 += protohelpers.SizeOfVarint(uint64(num))
		}
		i -= pksize2
		j1 := i
		for _, num1 := range m.RelativeXips {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA[j1] = uint8(num)
			j1++
		}
		i = protohelpers.EncodeVarint(dAtA, i, uint64(pksize2))
		i--
		dAtA[i] = 0x1a
	}
	if m.RelativeXmax != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.RelativeXmax))
		i--
		dAtA[i] = 0x10
	}
	if m.Xmin != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Xmin))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PostgresRevision) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Xmin != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Xmin))
	}
	if m.RelativeXmax != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.RelativeXmax))
	}
	if len(m.RelativeXips) > 0 {
		l = 0
		for _, e := range m.RelativeXips {
			l += protohelpers.SizeOfVarint(uint64(e))
		}
		n += 1 + protohelpers.SizeOfVarint(uint64(l)) + l
	}
	if m.OptionalTxid != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.OptionalTxid))
	}
	if m.OptionalTimestamp != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.OptionalTimestamp))
	}
	n += len(m.unknownFields)
	return n
}

func (m *PostgresRevision) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
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
			return fmt.Errorf("proto: PostgresRevision: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostgresRevision: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Xmin", wireType)
			}
			m.Xmin = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Xmin |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelativeXmax", wireType)
			}
			m.RelativeXmax = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RelativeXmax |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protohelpers.ErrIntOverflow
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.RelativeXips = append(m.RelativeXips, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protohelpers.ErrIntOverflow
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
					return protohelpers.ErrInvalidLength
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return protohelpers.ErrInvalidLength
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.RelativeXips) == 0 {
					m.RelativeXips = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return protohelpers.ErrIntOverflow
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.RelativeXips = append(m.RelativeXips, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field RelativeXips", wireType)
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalTxid", wireType)
			}
			m.OptionalTxid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OptionalTxid |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptionalTimestamp", wireType)
			}
			m.OptionalTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OptionalTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

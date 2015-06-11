// Code generated by protoc-gen-gogo.
// source: cockroach/proto/log.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "gogoproto/gogo.pb"

import io "io"
import fmt "fmt"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = math.Inf

// Log represents a cockroach structured log entry.
type LogEntry struct {
	// Log message severity.
	Severity int32 `protobuf:"varint,1,opt,name=severity" json:"severity"`
	// Time, measured in nanoseconds since the epoch.
	Time int64 `protobuf:"varint,2,opt,name=time" json:"time"`
	// Thread id of logging routine.
	ThreadID int32 `protobuf:"varint,3,opt,name=thread_id" json:"thread_id"`
	// File which generated log statement.
	File string `protobuf:"bytes,4,opt,name=file" json:"file"`
	// Line in file which generated log statement.
	Line int32 `protobuf:"varint,5,opt,name=line" json:"line"`
	// Log format message.
	Format string         `protobuf:"bytes,6,opt,name=format" json:"format"`
	Args   []LogEntry_Arg `protobuf:"bytes,7,rep,name=args" json:"args"`
	// Optional parameters which may be set with log entry.
	NodeID  *NodeID  `protobuf:"varint,8,opt,name=node_id,casttype=NodeID" json:"node_id,omitempty"`
	StoreID *StoreID `protobuf:"varint,9,opt,name=store_id,casttype=StoreID" json:"store_id,omitempty"`
	RaftID  *RaftID  `protobuf:"varint,10,opt,name=raft_id,casttype=RaftID" json:"raft_id,omitempty"`
	Method  *Method  `protobuf:"varint,11,opt,name=method,casttype=Method" json:"method,omitempty"`
	Key     Key      `protobuf:"bytes,12,opt,name=key,casttype=Key" json:"key,omitempty"`
	// Stack traces if requested.
	Stacks           []byte `protobuf:"bytes,13,opt,name=stacks" json:"stacks"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto1.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}

func (m *LogEntry) GetSeverity() int32 {
	if m != nil {
		return m.Severity
	}
	return 0
}

func (m *LogEntry) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *LogEntry) GetThreadID() int32 {
	if m != nil {
		return m.ThreadID
	}
	return 0
}

func (m *LogEntry) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *LogEntry) GetLine() int32 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *LogEntry) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *LogEntry) GetArgs() []LogEntry_Arg {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *LogEntry) GetStacks() []byte {
	if m != nil {
		return m.Stacks
	}
	return nil
}

// Log format arguments.
type LogEntry_Arg struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type"`
	Str  string `protobuf:"bytes,2,opt,name=str" json:"str"`
	// Optional json representation.
	Json             []byte `protobuf:"bytes,3,opt,name=json" json:"json"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *LogEntry_Arg) Reset()         { *m = LogEntry_Arg{} }
func (m *LogEntry_Arg) String() string { return proto1.CompactTextString(m) }
func (*LogEntry_Arg) ProtoMessage()    {}

func (m *LogEntry_Arg) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *LogEntry_Arg) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *LogEntry_Arg) GetJson() []byte {
	if m != nil {
		return m.Json
	}
	return nil
}

func init() {
}
func (m *LogEntry) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Severity", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.Severity |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.Time |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThreadID", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.ThreadID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field File", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.File = string(data[index:postIndex])
			index = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Line", wireType)
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.Line |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Format", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Format = string(data[index:postIndex])
			index = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Args", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Args = append(m.Args, LogEntry_Arg{})
			if err := m.Args[len(m.Args)-1].Unmarshal(data[index:postIndex]); err != nil {
				return err
			}
			index = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			var v NodeID
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (NodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NodeID = &v
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreID", wireType)
			}
			var v StoreID
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (StoreID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.StoreID = &v
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftID", wireType)
			}
			var v RaftID
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (RaftID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RaftID = &v
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var v Method
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (Method(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Method = &v
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append([]byte{}, data[index:postIndex]...)
			index = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stacks", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Stacks = append([]byte{}, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}

	return nil
}
func (m *LogEntry_Arg) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(data[index:postIndex])
			index = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Str", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Str = string(data[index:postIndex])
			index = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Json", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Json = append([]byte{}, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}

	return nil
}
func (m *LogEntry) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovLog(uint64(m.Severity))
	n += 1 + sovLog(uint64(m.Time))
	n += 1 + sovLog(uint64(m.ThreadID))
	l = len(m.File)
	n += 1 + l + sovLog(uint64(l))
	n += 1 + sovLog(uint64(m.Line))
	l = len(m.Format)
	n += 1 + l + sovLog(uint64(l))
	if len(m.Args) > 0 {
		for _, e := range m.Args {
			l = e.Size()
			n += 1 + l + sovLog(uint64(l))
		}
	}
	if m.NodeID != nil {
		n += 1 + sovLog(uint64(*m.NodeID))
	}
	if m.StoreID != nil {
		n += 1 + sovLog(uint64(*m.StoreID))
	}
	if m.RaftID != nil {
		n += 1 + sovLog(uint64(*m.RaftID))
	}
	if m.Method != nil {
		n += 1 + sovLog(uint64(*m.Method))
	}
	if m.Key != nil {
		l = len(m.Key)
		n += 1 + l + sovLog(uint64(l))
	}
	if m.Stacks != nil {
		l = len(m.Stacks)
		n += 1 + l + sovLog(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LogEntry_Arg) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	n += 1 + l + sovLog(uint64(l))
	l = len(m.Str)
	n += 1 + l + sovLog(uint64(l))
	if m.Json != nil {
		l = len(m.Json)
		n += 1 + l + sovLog(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLog(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLog(x uint64) (n int) {
	return sovLog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LogEntry) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *LogEntry) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintLog(data, i, uint64(m.Severity))
	data[i] = 0x10
	i++
	i = encodeVarintLog(data, i, uint64(m.Time))
	data[i] = 0x18
	i++
	i = encodeVarintLog(data, i, uint64(m.ThreadID))
	data[i] = 0x22
	i++
	i = encodeVarintLog(data, i, uint64(len(m.File)))
	i += copy(data[i:], m.File)
	data[i] = 0x28
	i++
	i = encodeVarintLog(data, i, uint64(m.Line))
	data[i] = 0x32
	i++
	i = encodeVarintLog(data, i, uint64(len(m.Format)))
	i += copy(data[i:], m.Format)
	if len(m.Args) > 0 {
		for _, msg := range m.Args {
			data[i] = 0x3a
			i++
			i = encodeVarintLog(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.NodeID != nil {
		data[i] = 0x40
		i++
		i = encodeVarintLog(data, i, uint64(*m.NodeID))
	}
	if m.StoreID != nil {
		data[i] = 0x48
		i++
		i = encodeVarintLog(data, i, uint64(*m.StoreID))
	}
	if m.RaftID != nil {
		data[i] = 0x50
		i++
		i = encodeVarintLog(data, i, uint64(*m.RaftID))
	}
	if m.Method != nil {
		data[i] = 0x58
		i++
		i = encodeVarintLog(data, i, uint64(*m.Method))
	}
	if m.Key != nil {
		data[i] = 0x62
		i++
		i = encodeVarintLog(data, i, uint64(len(m.Key)))
		i += copy(data[i:], m.Key)
	}
	if m.Stacks != nil {
		data[i] = 0x6a
		i++
		i = encodeVarintLog(data, i, uint64(len(m.Stacks)))
		i += copy(data[i:], m.Stacks)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *LogEntry_Arg) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *LogEntry_Arg) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintLog(data, i, uint64(len(m.Type)))
	i += copy(data[i:], m.Type)
	data[i] = 0x12
	i++
	i = encodeVarintLog(data, i, uint64(len(m.Str)))
	i += copy(data[i:], m.Str)
	if m.Json != nil {
		data[i] = 0x1a
		i++
		i = encodeVarintLog(data, i, uint64(len(m.Json)))
		i += copy(data[i:], m.Json)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Log(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Log(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintLog(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}

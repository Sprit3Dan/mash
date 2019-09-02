// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event/event.proto

package event

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FOO int32

const (
	FOO_X FOO = 17
)

var FOO_name = map[int32]string{
	17: "X",
}

var FOO_value = map[string]int32{
	"X": 17,
}

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}

func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}

func (x *FOO) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FOO_value, data, "FOO")
	if err != nil {
		return err
	}
	*x = FOO(value)
	return nil
}

func (FOO) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_510ac585eaa75cc0, []int{0}
}

type Event struct {
	Label                *string  `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
	Type                 *int32   `protobuf:"varint,2,opt,name=type,def=77" json:"type,omitempty"`
	Reps                 []int64  `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_510ac585eaa75cc0, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

const Default_Event_Type int32 = 77

func (m *Event) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Event) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Event_Type
}

func (m *Event) GetReps() []int64 {
	if m != nil {
		return m.Reps
	}
	return nil
}

func init() {
	proto.RegisterEnum("mash.FOO", FOO_name, FOO_value)
	proto.RegisterType((*Event)(nil), "mash.Event")
}

func init() { proto.RegisterFile("event/event.proto", fileDescriptor_510ac585eaa75cc0) }

var fileDescriptor_510ac585eaa75cc0 = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2d, 0x4b, 0xcd,
	0x2b, 0xd1, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0xb9, 0x89, 0xc5, 0x19,
	0x4a, 0x9e, 0x5c, 0xac, 0xae, 0x20, 0x41, 0x21, 0x11, 0x2e, 0xd6, 0x9c, 0xc4, 0xa4, 0xd4, 0x1c,
	0x09, 0x46, 0x05, 0x26, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x8c, 0x8b, 0xa5, 0xa4, 0xb2, 0x20,
	0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0xd5, 0x8a, 0xc9, 0xdc, 0x3c, 0x08, 0xcc, 0x17, 0x12, 0xe2,
	0x62, 0x29, 0x4a, 0x2d, 0x28, 0x96, 0x60, 0x56, 0x60, 0xd6, 0x60, 0x0e, 0x02, 0xb3, 0xb5, 0x78,
	0xb8, 0x98, 0xdd, 0xfc, 0xfd, 0x85, 0x58, 0xb9, 0x18, 0x23, 0x04, 0x04, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xdc, 0x87, 0xf5, 0xe3, 0x72, 0x00, 0x00, 0x00,
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v2/service.proto

package v2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	v1 "v1"
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

type Entity_Type int32

const (
	Entity_Board  Entity_Type = 0
	Entity_Column Entity_Type = 1
	Entity_Card   Entity_Type = 2
)

var Entity_Type_name = map[int32]string{
	0: "Board",
	1: "Column",
	2: "Card",
}

var Entity_Type_value = map[string]int32{
	"Board":  0,
	"Column": 1,
	"Card":   2,
}

func (x Entity_Type) String() string {
	return proto.EnumName(Entity_Type_name, int32(x))
}

func (Entity_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_561abc64c574c56c, []int{0, 0}
}

type Entity struct {
	// Types that are valid to be assigned to Entity:
	//	*Entity_Board
	//	*Entity_Column
	//	*Entity_Card
	Entity               isEntity_Entity `protobuf_oneof:"entity"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_561abc64c574c56c, []int{0}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

type isEntity_Entity interface {
	isEntity_Entity()
}

type Entity_Board struct {
	Board *v1.Board `protobuf:"bytes,1,opt,name=board,proto3,oneof"`
}

type Entity_Column struct {
	Column *v1.Column `protobuf:"bytes,2,opt,name=column,proto3,oneof"`
}

type Entity_Card struct {
	Card *v1.Card `protobuf:"bytes,3,opt,name=card,proto3,oneof"`
}

func (*Entity_Board) isEntity_Entity() {}

func (*Entity_Column) isEntity_Entity() {}

func (*Entity_Card) isEntity_Entity() {}

func (m *Entity) GetEntity() isEntity_Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *Entity) GetBoard() *v1.Board {
	if x, ok := m.GetEntity().(*Entity_Board); ok {
		return x.Board
	}
	return nil
}

func (m *Entity) GetColumn() *v1.Column {
	if x, ok := m.GetEntity().(*Entity_Column); ok {
		return x.Column
	}
	return nil
}

func (m *Entity) GetCard() *v1.Card {
	if x, ok := m.GetEntity().(*Entity_Card); ok {
		return x.Card
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Entity) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Entity_Board)(nil),
		(*Entity_Column)(nil),
		(*Entity_Card)(nil),
	}
}

type Entity_URI struct {
	Id                   *v1.URI     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 Entity_Type `protobuf:"varint,2,opt,name=type,proto3,enum=octolab.api.tablo.v2.Entity_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Entity_URI) Reset()         { *m = Entity_URI{} }
func (m *Entity_URI) String() string { return proto.CompactTextString(m) }
func (*Entity_URI) ProtoMessage()    {}
func (*Entity_URI) Descriptor() ([]byte, []int) {
	return fileDescriptor_561abc64c574c56c, []int{0, 0}
}

func (m *Entity_URI) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity_URI.Unmarshal(m, b)
}
func (m *Entity_URI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity_URI.Marshal(b, m, deterministic)
}
func (m *Entity_URI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity_URI.Merge(m, src)
}
func (m *Entity_URI) XXX_Size() int {
	return xxx_messageInfo_Entity_URI.Size(m)
}
func (m *Entity_URI) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity_URI.DiscardUnknown(m)
}

var xxx_messageInfo_Entity_URI proto.InternalMessageInfo

func (m *Entity_URI) GetId() *v1.URI {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Entity_URI) GetType() Entity_Type {
	if m != nil {
		return m.Type
	}
	return Entity_Board
}

func init() {
	proto.RegisterEnum("octolab.api.tablo.v2.Entity_Type", Entity_Type_name, Entity_Type_value)
	proto.RegisterType((*Entity)(nil), "octolab.api.tablo.v2.Entity")
	proto.RegisterType((*Entity_URI)(nil), "octolab.api.tablo.v2.Entity.URI")
}

func init() { proto.RegisterFile("v2/service.proto", fileDescriptor_561abc64c574c56c) }

var fileDescriptor_561abc64c574c56c = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x4e, 0xc2, 0x40,
	0x14, 0x85, 0x61, 0x28, 0x13, 0xbc, 0x46, 0x6d, 0x6e, 0x5c, 0x60, 0x65, 0x81, 0x6c, 0xd4, 0xcd,
	0x20, 0x43, 0x74, 0x6f, 0x11, 0x02, 0xdb, 0xd1, 0xba, 0x70, 0xd7, 0x9f, 0x89, 0x99, 0xa4, 0x30,
	0x4d, 0x1d, 0x9b, 0xf0, 0x42, 0xbe, 0x88, 0x2f, 0x66, 0x3a, 0xd5, 0x85, 0x49, 0x5b, 0x5d, 0x36,
	0xf9, 0xbe, 0x7b, 0xce, 0x69, 0x0b, 0x6e, 0xc1, 0xa7, 0x6f, 0x32, 0x2f, 0x54, 0x2c, 0x59, 0x96,
	0x6b, 0xa3, 0xf1, 0x54, 0xc7, 0x46, 0xa7, 0x61, 0xc4, 0xc2, 0x4c, 0x31, 0x13, 0x46, 0xa9, 0x66,
	0x05, 0xf7, 0x4e, 0x8a, 0xd9, 0x34, 0xd6, 0xdb, 0xad, 0xde, 0x55, 0x98, 0xe7, 0x16, 0xb3, 0xdf,
	0xe2, 0xe4, 0x93, 0x00, 0x5d, 0xee, 0x8c, 0x32, 0x7b, 0x9c, 0x43, 0x3f, 0xd2, 0x61, 0x9e, 0x0c,
	0xbb, 0xe3, 0xee, 0xd5, 0x21, 0x3f, 0x67, 0x35, 0x37, 0x67, 0xcc, 0x2f, 0x91, 0x75, 0x47, 0x54,
	0x2c, 0xde, 0x01, 0x8d, 0x75, 0xfa, 0xbe, 0xdd, 0x0d, 0x89, 0xb5, 0x46, 0xf5, 0xd6, 0xc2, 0x32,
	0xeb, 0x8e, 0xf8, 0xa6, 0xf1, 0x06, 0x9c, 0xb8, 0xcc, 0xea, 0x59, 0xcb, 0x6b, 0xb0, 0xaa, 0x28,
	0x4b, 0x7a, 0xaf, 0xd0, 0x0b, 0xc4, 0x06, 0xaf, 0x81, 0xa8, 0x9f, 0x8a, 0x67, 0xf5, 0x5a, 0x20,
	0x36, 0x82, 0xa8, 0x04, 0x6f, 0xc1, 0x31, 0xfb, 0x4c, 0xda, 0x66, 0xc7, 0xfc, 0xa2, 0x0e, 0xe6,
	0xac, 0x1a, 0xcf, 0x9e, 0xf6, 0x99, 0x14, 0x16, 0x9f, 0x5c, 0x82, 0x53, 0x3e, 0xe1, 0x01, 0xf4,
	0xed, 0x58, 0xb7, 0x83, 0x00, 0xb4, 0x5a, 0xe0, 0x76, 0x71, 0x00, 0x4e, 0xd9, 0xcb, 0x25, 0xfe,
	0x00, 0xa8, 0xb4, 0x36, 0xff, 0x20, 0x70, 0xb4, 0x4c, 0x94, 0xd1, 0xf9, 0x63, 0xf5, 0x76, 0xf1,
	0x1e, 0xe8, 0x22, 0x97, 0xa1, 0x91, 0x38, 0x6a, 0xcb, 0xf5, 0x9a, 0x27, 0xe0, 0x0a, 0x1c, 0x21,
	0xc3, 0x04, 0xc7, 0xad, 0xc5, 0x03, 0xb1, 0xf1, 0x5a, 0x23, 0xd0, 0x07, 0x1a, 0x64, 0xc9, 0xdf,
	0x55, 0x1a, 0x3e, 0xc2, 0xb3, 0x56, 0x09, 0xae, 0x80, 0x3e, 0xc8, 0x54, 0x1a, 0xf9, 0x8f, 0x36,
	0x2d, 0x77, 0x7c, 0xe7, 0x85, 0x14, 0x3c, 0xa2, 0xf6, 0xdf, 0x9b, 0x7f, 0x05, 0x00, 0x00, 0xff,
	0xff, 0xbb, 0x4a, 0x28, 0xe5, 0xc8, 0x02, 0x00, 0x00,
}

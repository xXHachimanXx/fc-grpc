// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UserResultStream struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResultStream) Reset()         { *m = UserResultStream{} }
func (m *UserResultStream) String() string { return proto.CompactTextString(m) }
func (*UserResultStream) ProtoMessage()    {}
func (*UserResultStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserResultStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResultStream.Unmarshal(m, b)
}
func (m *UserResultStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResultStream.Marshal(b, m, deterministic)
}
func (m *UserResultStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResultStream.Merge(m, src)
}
func (m *UserResultStream) XXX_Size() int {
	return xxx_messageInfo_UserResultStream.Size(m)
}
func (m *UserResultStream) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResultStream.DiscardUnknown(m)
}

var xxx_messageInfo_UserResultStream proto.InternalMessageInfo

func (m *UserResultStream) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *UserResultStream) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Users struct {
	User                 []*User  `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (m *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(m, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*UserResultStream)(nil), "pb.UserResultStream")
	proto.RegisterType((*Users)(nil), "pb.Users")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xd9, 0x34, 0xe9, 0x9f, 0x29, 0x14, 0x1d, 0x8a, 0x04, 0x51, 0xa8, 0x01, 0x21, 0xa7,
	0xb5, 0xc4, 0x83, 0x57, 0xed, 0xc9, 0x73, 0x8b, 0x1e, 0xbc, 0x65, 0xcd, 0x80, 0x81, 0xc6, 0x84,
	0x9d, 0x8d, 0x9f, 0xce, 0x0f, 0x27, 0x99, 0xac, 0x31, 0xe0, 0xa1, 0xb7, 0xf7, 0x5e, 0x7e, 0xef,
	0x65, 0x97, 0x05, 0x68, 0x99, 0xac, 0x6e, 0x6c, 0xed, 0x6a, 0x0c, 0x1a, 0x93, 0x3c, 0x42, 0xf8,
	0xc2, 0x64, 0x71, 0x05, 0x41, 0x59, 0xc4, 0x6a, 0xa3, 0xd2, 0xc5, 0x3e, 0x28, 0x0b, 0x44, 0x08,
	0x3f, 0xf3, 0x8a, 0xe2, 0x40, 0x12, 0xd1, 0xb8, 0x86, 0x88, 0xaa, 0xbc, 0x3c, 0xc6, 0x13, 0x09,
	0x7b, 0x93, 0x3c, 0xc3, 0x59, 0xb7, 0xb0, 0x27, 0x6e, 0x8f, 0xee, 0xe0, 0x2c, 0xe5, 0x15, 0x5e,
	0xc0, 0x94, 0x5d, 0xee, 0x5a, 0xf6, 0x8b, 0xde, 0xe1, 0x15, 0x84, 0xdd, 0xff, 0x65, 0x75, 0x99,
	0xcd, 0x75, 0x63, 0xb4, 0x74, 0x25, 0x4d, 0x6e, 0x21, 0xea, 0xdc, 0x1f, 0xa6, 0x36, 0x93, 0xff,
	0x58, 0xf6, 0xad, 0x60, 0xd9, 0xd9, 0x03, 0xd9, 0xaf, 0xf2, 0x9d, 0xf0, 0x1a, 0x66, 0x4f, 0x45,
	0x21, 0xb7, 0x18, 0xd0, 0xcb, 0x41, 0x61, 0x06, 0x2b, 0xff, 0xf9, 0x95, 0xac, 0xa9, 0x99, 0x46,
	0xd4, 0x7a, 0x98, 0x1e, 0x9d, 0x7e, 0xab, 0xf0, 0x06, 0xe6, 0xbe, 0xc3, 0x23, 0x7a, 0xf1, 0xab,
	0x38, 0x55, 0xf8, 0x00, 0xe7, 0x1e, 0xe9, 0x5b, 0xbb, 0xda, 0x7d, 0x9c, 0x5a, 0x4e, 0xd5, 0x56,
	0xed, 0x66, 0x6f, 0x91, 0xd6, 0x77, 0x8d, 0x31, 0x53, 0x79, 0x85, 0xfb, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4c, 0xc6, 0x0a, 0x01, 0x93, 0x01, 0x00, 0x00,
}

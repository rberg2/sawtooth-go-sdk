// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/rberg2/sawtooth-go-sdk/protobuf/setting_pb2/setting.proto

/*
Package setting_pb2 is a generated protocol buffer package.

It is generated from these files:
	github.com/rberg2/sawtooth-go-sdk/protobuf/setting_pb2/setting.proto

It has these top-level messages:
	Setting
*/
package setting_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Setting Container for on-chain configuration key/value pairs
type Setting struct {
	// List of setting entries - more than one implies a state key collision
	Entries []*Setting_Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *Setting) Reset()                    { *m = Setting{} }
func (m *Setting) String() string            { return proto.CompactTextString(m) }
func (*Setting) ProtoMessage()               {}
func (*Setting) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Setting) GetEntries() []*Setting_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

// Contains a setting entry (or entries, in the case of collisions).
type Setting_Entry struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Setting_Entry) Reset()                    { *m = Setting_Entry{} }
func (m *Setting_Entry) String() string            { return proto.CompactTextString(m) }
func (*Setting_Entry) ProtoMessage()               {}
func (*Setting_Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Setting_Entry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Setting_Entry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Setting)(nil), "Setting")
	proto.RegisterType((*Setting_Entry)(nil), "Setting.Entry")
}

func init() { proto.RegisterFile("github.com/rberg2/sawtooth-go-sdk/protobuf/setting_pb2/setting.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x2f, 0x4e, 0x2c, 0x2f,
	0xc9, 0xcf, 0x2f, 0xc9, 0x88, 0x2f, 0x4e, 0xc9, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a,
	0x4d, 0xd3, 0x2f, 0x4e, 0x2d, 0x29, 0xc9, 0xcc, 0x4b, 0x8f, 0x2f, 0x48, 0x32, 0x82, 0xb1, 0xf5,
	0xc0, 0x92, 0x4a, 0x29, 0x5c, 0xec, 0xc1, 0x10, 0x01, 0x21, 0x0d, 0x2e, 0xf6, 0xd4, 0xbc, 0x92,
	0xa2, 0xcc, 0xd4, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x3e, 0x3d, 0xa8, 0x94, 0x9e,
	0x2b, 0x50, 0xbc, 0x32, 0x08, 0x26, 0x2d, 0xa5, 0xcf, 0xc5, 0x0a, 0x16, 0x11, 0x12, 0xe0, 0x62,
	0xce, 0x4e, 0xad, 0x04, 0x2a, 0x67, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58, 0xcb,
	0x12, 0x73, 0x4a, 0x53, 0x25, 0x98, 0xc0, 0x62, 0x10, 0x8e, 0x93, 0x1a, 0x97, 0x28, 0xcc, 0x61,
	0x7a, 0x40, 0x87, 0xe9, 0xc1, 0x1c, 0x16, 0xc0, 0x18, 0xc5, 0x8d, 0xe4, 0xb6, 0x24, 0x36, 0xb0,
	0x84, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x06, 0x0d, 0xef, 0x1f, 0xc7, 0x00, 0x00, 0x00,
}

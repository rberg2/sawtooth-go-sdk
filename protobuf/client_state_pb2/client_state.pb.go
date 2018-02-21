// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/rberg2/sawtooth-go-sdk/protobuf/client_state_pb2/client_state.proto

/*
Package client_state_pb2 is a generated protocol buffer package.

It is generated from these files:
	github.com/rberg2/sawtooth-go-sdk/protobuf/client_state_pb2/client_state.proto

It has these top-level messages:
	ClientStateListRequest
	ClientStateListResponse
	ClientStateGetRequest
	ClientStateGetResponse
*/
package client_state_pb2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import client_list_control "github.com/rberg2/sawtooth-go-sdk/protobuf/client_list_control_pb2"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientStateListResponse_Status int32

const (
	ClientStateListResponse_STATUS_UNSET    ClientStateListResponse_Status = 0
	ClientStateListResponse_OK              ClientStateListResponse_Status = 1
	ClientStateListResponse_INTERNAL_ERROR  ClientStateListResponse_Status = 2
	ClientStateListResponse_NOT_READY       ClientStateListResponse_Status = 3
	ClientStateListResponse_NO_ROOT         ClientStateListResponse_Status = 4
	ClientStateListResponse_NO_RESOURCE     ClientStateListResponse_Status = 5
	ClientStateListResponse_INVALID_PAGING  ClientStateListResponse_Status = 6
	ClientStateListResponse_INVALID_SORT    ClientStateListResponse_Status = 7
	ClientStateListResponse_INVALID_ADDRESS ClientStateListResponse_Status = 8
	ClientStateListResponse_INVALID_ROOT    ClientStateListResponse_Status = 9
)

var ClientStateListResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	3: "NOT_READY",
	4: "NO_ROOT",
	5: "NO_RESOURCE",
	6: "INVALID_PAGING",
	7: "INVALID_SORT",
	8: "INVALID_ADDRESS",
	9: "INVALID_ROOT",
}
var ClientStateListResponse_Status_value = map[string]int32{
	"STATUS_UNSET":    0,
	"OK":              1,
	"INTERNAL_ERROR":  2,
	"NOT_READY":       3,
	"NO_ROOT":         4,
	"NO_RESOURCE":     5,
	"INVALID_PAGING":  6,
	"INVALID_SORT":    7,
	"INVALID_ADDRESS": 8,
	"INVALID_ROOT":    9,
}

func (x ClientStateListResponse_Status) String() string {
	return proto.EnumName(ClientStateListResponse_Status_name, int32(x))
}
func (ClientStateListResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

type ClientStateGetResponse_Status int32

const (
	ClientStateGetResponse_STATUS_UNSET    ClientStateGetResponse_Status = 0
	ClientStateGetResponse_OK              ClientStateGetResponse_Status = 1
	ClientStateGetResponse_INTERNAL_ERROR  ClientStateGetResponse_Status = 2
	ClientStateGetResponse_NOT_READY       ClientStateGetResponse_Status = 3
	ClientStateGetResponse_NO_ROOT         ClientStateGetResponse_Status = 4
	ClientStateGetResponse_NO_RESOURCE     ClientStateGetResponse_Status = 5
	ClientStateGetResponse_INVALID_ADDRESS ClientStateGetResponse_Status = 6
	ClientStateGetResponse_INVALID_ROOT    ClientStateGetResponse_Status = 7
)

var ClientStateGetResponse_Status_name = map[int32]string{
	0: "STATUS_UNSET",
	1: "OK",
	2: "INTERNAL_ERROR",
	3: "NOT_READY",
	4: "NO_ROOT",
	5: "NO_RESOURCE",
	6: "INVALID_ADDRESS",
	7: "INVALID_ROOT",
}
var ClientStateGetResponse_Status_value = map[string]int32{
	"STATUS_UNSET":    0,
	"OK":              1,
	"INTERNAL_ERROR":  2,
	"NOT_READY":       3,
	"NO_ROOT":         4,
	"NO_RESOURCE":     5,
	"INVALID_ADDRESS": 6,
	"INVALID_ROOT":    7,
}

func (x ClientStateGetResponse_Status) String() string {
	return proto.EnumName(ClientStateGetResponse_Status_name, int32(x))
}
func (ClientStateGetResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

// A request to list every entry in global state. Defaults to the most current
// tree, but can fetch older state by specifying a state root. Results can be
// further filtered by specifying a subtree with a partial address.
type ClientStateListRequest struct {
	StateRoot string                                    `protobuf:"bytes,1,opt,name=state_root,json=stateRoot" json:"state_root,omitempty"`
	Address   string                                    `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	Paging    *client_list_control.ClientPagingControls `protobuf:"bytes,4,opt,name=paging" json:"paging,omitempty"`
	Sorting   []*client_list_control.ClientSortControls `protobuf:"bytes,5,rep,name=sorting" json:"sorting,omitempty"`
}

func (m *ClientStateListRequest) Reset()                    { *m = ClientStateListRequest{} }
func (m *ClientStateListRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientStateListRequest) ProtoMessage()               {}
func (*ClientStateListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientStateListRequest) GetStateRoot() string {
	if m != nil {
		return m.StateRoot
	}
	return ""
}

func (m *ClientStateListRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClientStateListRequest) GetPaging() *client_list_control.ClientPagingControls {
	if m != nil {
		return m.Paging
	}
	return nil
}

func (m *ClientStateListRequest) GetSorting() []*client_list_control.ClientSortControls {
	if m != nil {
		return m.Sorting
	}
	return nil
}

type ClientStateListResponse struct {
	Status    ClientStateListResponse_Status            `protobuf:"varint,1,opt,name=status,enum=ClientStateListResponse_Status" json:"status,omitempty"`
	Entries   []*ClientStateListResponse_Entry          `protobuf:"bytes,2,rep,name=entries" json:"entries,omitempty"`
	StateRoot string                                    `protobuf:"bytes,3,opt,name=state_root,json=stateRoot" json:"state_root,omitempty"`
	Paging    *client_list_control.ClientPagingResponse `protobuf:"bytes,4,opt,name=paging" json:"paging,omitempty"`
}

func (m *ClientStateListResponse) Reset()                    { *m = ClientStateListResponse{} }
func (m *ClientStateListResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientStateListResponse) ProtoMessage()               {}
func (*ClientStateListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClientStateListResponse) GetStatus() ClientStateListResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientStateListResponse_STATUS_UNSET
}

func (m *ClientStateListResponse) GetEntries() []*ClientStateListResponse_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *ClientStateListResponse) GetStateRoot() string {
	if m != nil {
		return m.StateRoot
	}
	return ""
}

func (m *ClientStateListResponse) GetPaging() *client_list_control.ClientPagingResponse {
	if m != nil {
		return m.Paging
	}
	return nil
}

// An entry in the State
type ClientStateListResponse_Entry struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *ClientStateListResponse_Entry) Reset()         { *m = ClientStateListResponse_Entry{} }
func (m *ClientStateListResponse_Entry) String() string { return proto.CompactTextString(m) }
func (*ClientStateListResponse_Entry) ProtoMessage()    {}
func (*ClientStateListResponse_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

func (m *ClientStateListResponse_Entry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClientStateListResponse_Entry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// A request from a client for a particular entry in global state.
// Like State List, it defaults to the newest state, but a state root
// can be used to specify older data. Unlike State List the request must be
// provided with a full address that corresponds to a single entry.
type ClientStateGetRequest struct {
	StateRoot string `protobuf:"bytes,1,opt,name=state_root,json=stateRoot" json:"state_root,omitempty"`
	Address   string `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
}

func (m *ClientStateGetRequest) Reset()                    { *m = ClientStateGetRequest{} }
func (m *ClientStateGetRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientStateGetRequest) ProtoMessage()               {}
func (*ClientStateGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientStateGetRequest) GetStateRoot() string {
	if m != nil {
		return m.StateRoot
	}
	return ""
}

func (m *ClientStateGetRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// The response to a State Get Request from the client. Sends back just
// the data stored at the entry, not the address. Also sends back the
// head block id used to facilitate further requests.
//
// Statuses:
//   * OK - everything worked as expected
//   * INTERNAL_ERROR - general error, such as protobuf failing to deserialize
//   * NOT_READY - the validator does not yet have a genesis block
//   * NO_ROOT - the state_root specified was not found
//   * NO_RESOURCE - the address specified doesn't exist
//   * INVALID_ADDRESS - address isn't a valid, i.e. it's a subtree (truncated)
type ClientStateGetResponse struct {
	Status    ClientStateGetResponse_Status `protobuf:"varint,1,opt,name=status,enum=ClientStateGetResponse_Status" json:"status,omitempty"`
	Value     []byte                        `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	StateRoot string                        `protobuf:"bytes,3,opt,name=state_root,json=stateRoot" json:"state_root,omitempty"`
}

func (m *ClientStateGetResponse) Reset()                    { *m = ClientStateGetResponse{} }
func (m *ClientStateGetResponse) String() string            { return proto.CompactTextString(m) }
func (*ClientStateGetResponse) ProtoMessage()               {}
func (*ClientStateGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ClientStateGetResponse) GetStatus() ClientStateGetResponse_Status {
	if m != nil {
		return m.Status
	}
	return ClientStateGetResponse_STATUS_UNSET
}

func (m *ClientStateGetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ClientStateGetResponse) GetStateRoot() string {
	if m != nil {
		return m.StateRoot
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientStateListRequest)(nil), "ClientStateListRequest")
	proto.RegisterType((*ClientStateListResponse)(nil), "ClientStateListResponse")
	proto.RegisterType((*ClientStateListResponse_Entry)(nil), "ClientStateListResponse.Entry")
	proto.RegisterType((*ClientStateGetRequest)(nil), "ClientStateGetRequest")
	proto.RegisterType((*ClientStateGetResponse)(nil), "ClientStateGetResponse")
	proto.RegisterEnum("ClientStateListResponse_Status", ClientStateListResponse_Status_name, ClientStateListResponse_Status_value)
	proto.RegisterEnum("ClientStateGetResponse_Status", ClientStateGetResponse_Status_name, ClientStateGetResponse_Status_value)
}

func init() {
	proto.RegisterFile("github.com/rberg2/sawtooth-go-sdk/protobuf/client_state_pb2/client_state.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0xb3, 0x26, 0xf4, 0x74, 0x6c, 0x96, 0x47, 0x21, 0x9a, 0x04, 0x4c, 0xb9, 0x9a,
	0x84, 0xc8, 0xa4, 0x22, 0xfe, 0x5d, 0x86, 0x36, 0xaa, 0x2a, 0xaa, 0xa4, 0x72, 0x52, 0x24, 0xb8,
	0xb1, 0xd2, 0xd5, 0x8c, 0x8a, 0xaa, 0x29, 0xb1, 0x0b, 0xe2, 0x15, 0xb8, 0xe4, 0x45, 0xe0, 0x21,
	0x78, 0x30, 0x1c, 0x27, 0x46, 0x69, 0xc6, 0xc6, 0x05, 0x12, 0x77, 0xf1, 0xe7, 0x9f, 0x3f, 0x9f,
	0x7c, 0x3e, 0x3a, 0xf0, 0x82, 0xa7, 0x9f, 0x45, 0x96, 0x89, 0xf7, 0x94, 0x2f, 0x3e, 0x9c, 0x6d,
	0xf2, 0x4c, 0x64, 0xf3, 0xed, 0xbb, 0xb3, 0xf3, 0xd5, 0x92, 0xad, 0x05, 0xe5, 0x22, 0x15, 0x8c,
	0x6e, 0xe6, 0xfd, 0x1d, 0xc1, 0x53, 0xd8, 0xf1, 0xe8, 0xda, 0xa3, 0xab, 0x25, 0x17, 0xf4, 0x3c,
	0x5b, 0x8b, 0x3c, 0x5b, 0xd5, 0x1d, 0xea, 0x7a, 0x69, 0xe4, 0x7e, 0x37, 0xe0, 0xce, 0x40, 0xed,
	0xc6, 0x85, 0xfd, 0x44, 0x12, 0x84, 0x7d, 0xdc, 0x32, 0x2e, 0xf0, 0x3d, 0x80, 0xb2, 0x86, 0x5c,
	0x5e, 0xe4, 0x18, 0x27, 0xc6, 0x69, 0x87, 0x74, 0x94, 0x42, 0xa4, 0x80, 0x1d, 0xb0, 0xd3, 0xc5,
	0x22, 0x67, 0x9c, 0x3b, 0xa6, 0xda, 0xd3, 0x4b, 0xfc, 0x08, 0xac, 0x4d, 0x7a, 0xb1, 0x5c, 0x5f,
	0x38, 0x7b, 0x72, 0xa3, 0xdb, 0xef, 0x79, 0xe5, 0x0d, 0x53, 0x25, 0x0e, 0xca, 0xfb, 0x39, 0xa9,
	0x20, 0x89, 0xdb, 0x3c, 0xcb, 0x45, 0xc1, 0xb7, 0x4f, 0x4c, 0xc9, 0x1f, 0x55, 0x7c, 0x2c, 0xd5,
	0xdf, 0xb4, 0x66, 0xdc, 0x9f, 0x26, 0xdc, 0xbd, 0x54, 0x31, 0xdf, 0x64, 0x6b, 0xce, 0xf0, 0x33,
	0xb0, 0x8a, 0x02, 0xb7, 0x5c, 0x95, 0x7b, 0xd0, 0x7f, 0xe0, 0x5d, 0x41, 0x7a, 0xb1, 0xc2, 0x48,
	0x85, 0xe3, 0xe7, 0x60, 0x4b, 0x2c, 0x5f, 0x32, 0xee, 0xb4, 0x54, 0x0d, 0xf7, 0xaf, 0x3c, 0x19,
	0x48, 0xee, 0x0b, 0xd1, 0x78, 0x23, 0x25, 0xb3, 0x99, 0xd2, 0xf5, 0x59, 0x68, 0x53, 0x9d, 0xc5,
	0xf1, 0x13, 0x68, 0x2b, 0xff, 0x7a, 0xba, 0xc6, 0x6e, 0xba, 0x18, 0xf6, 0x16, 0xa9, 0x48, 0x65,
	0x9d, 0xc6, 0xe9, 0x3e, 0x51, 0xdf, 0xee, 0x0f, 0x03, 0xac, 0xf2, 0x8f, 0x30, 0x82, 0xfd, 0x38,
	0xf1, 0x93, 0x59, 0x4c, 0x67, 0x61, 0x1c, 0x24, 0xe8, 0x06, 0xb6, 0xa0, 0x15, 0xbd, 0x42, 0x86,
	0x3c, 0x78, 0x30, 0x0e, 0x93, 0x80, 0x84, 0xfe, 0x84, 0x06, 0x84, 0x44, 0x04, 0xb5, 0xf0, 0x2d,
	0xe8, 0x84, 0x51, 0x42, 0x49, 0xe0, 0x0f, 0xdf, 0x20, 0x13, 0x77, 0xc1, 0x0e, 0x23, 0x4a, 0xa2,
	0x28, 0x41, 0x7b, 0xf8, 0x10, 0xba, 0xc5, 0x22, 0x88, 0xa3, 0x19, 0x19, 0x04, 0xa8, 0x5d, 0x1a,
	0xbc, 0xf6, 0x27, 0xe3, 0x21, 0x9d, 0xfa, 0xa3, 0x71, 0x38, 0x42, 0x56, 0x71, 0x9d, 0xd6, 0xe2,
	0x88, 0x24, 0xc8, 0xc6, 0x47, 0x70, 0xa8, 0x15, 0x7f, 0x38, 0x94, 0xc7, 0x63, 0x74, 0xb3, 0x8e,
	0x29, 0xf7, 0x8e, 0x3b, 0x85, 0x5e, 0x2d, 0xe1, 0x11, 0xfb, 0xe7, 0xb6, 0x73, 0xbf, 0xb5, 0x76,
	0x5a, 0x59, 0x59, 0x56, 0x7d, 0xf1, 0xb4, 0xd1, 0x17, 0x3b, 0xaf, 0x5b, 0x03, 0x9b, 0x6d, 0x71,
	0x1b, 0xda, 0x9f, 0xd2, 0xd5, 0x96, 0x55, 0x61, 0x97, 0x8b, 0xbf, 0x3c, 0xb9, 0xfb, 0xf5, 0xbf,
	0x3c, 0xc6, 0x1f, 0x62, 0xb6, 0x2e, 0xc5, 0x6c, 0xbf, 0x7c, 0x08, 0x3d, 0x3d, 0x2a, 0x3c, 0x39,
	0x2a, 0x3c, 0x3d, 0x2a, 0xa6, 0xc6, 0x5b, 0xd4, 0x1c, 0x34, 0x73, 0x4b, 0xed, 0x3e, 0xfe, 0x15,
	0x00, 0x00, 0xff, 0xff, 0x80, 0xaf, 0x47, 0x28, 0x99, 0x04, 0x00, 0x00,
}

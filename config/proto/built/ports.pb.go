// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ports.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Port struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country              string   `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Alias                []string `protobuf:"bytes,5,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions              []string `protobuf:"bytes,6,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates          []string `protobuf:"bytes,7,rep,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province             string   `protobuf:"bytes,8,opt,name=province,proto3" json:"province,omitempty"`
	Timezone             string   `protobuf:"bytes,9,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs               []string `protobuf:"bytes,10,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code                 []string `protobuf:"bytes,11,rep,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_d80a8240fd02b040, []int{0}
}

func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (m *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(m, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Port) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Port) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Port) GetAlias() []string {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *Port) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *Port) GetCoordinates() []string {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func (m *Port) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *Port) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *Port) GetUnlocs() []string {
	if m != nil {
		return m.Unlocs
	}
	return nil
}

func (m *Port) GetCode() []string {
	if m != nil {
		return m.Code
	}
	return nil
}

type RequestWithId struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestWithId) Reset()         { *m = RequestWithId{} }
func (m *RequestWithId) String() string { return proto.CompactTextString(m) }
func (*RequestWithId) ProtoMessage()    {}
func (*RequestWithId) Descriptor() ([]byte, []int) {
	return fileDescriptor_d80a8240fd02b040, []int{1}
}

func (m *RequestWithId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestWithId.Unmarshal(m, b)
}
func (m *RequestWithId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestWithId.Marshal(b, m, deterministic)
}
func (m *RequestWithId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestWithId.Merge(m, src)
}
func (m *RequestWithId) XXX_Size() int {
	return xxx_messageInfo_RequestWithId.Size(m)
}
func (m *RequestWithId) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestWithId.DiscardUnknown(m)
}

var xxx_messageInfo_RequestWithId proto.InternalMessageInfo

func (m *RequestWithId) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CreatePortRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country              string   `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Alias                []string `protobuf:"bytes,5,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions              []string `protobuf:"bytes,6,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates          []string `protobuf:"bytes,7,rep,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province             string   `protobuf:"bytes,8,opt,name=province,proto3" json:"province,omitempty"`
	Timezone             string   `protobuf:"bytes,9,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs               []string `protobuf:"bytes,10,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code                 []string `protobuf:"bytes,11,rep,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePortRequest) Reset()         { *m = CreatePortRequest{} }
func (m *CreatePortRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePortRequest) ProtoMessage()    {}
func (*CreatePortRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d80a8240fd02b040, []int{2}
}

func (m *CreatePortRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePortRequest.Unmarshal(m, b)
}
func (m *CreatePortRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePortRequest.Marshal(b, m, deterministic)
}
func (m *CreatePortRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePortRequest.Merge(m, src)
}
func (m *CreatePortRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePortRequest.Size(m)
}
func (m *CreatePortRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePortRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePortRequest proto.InternalMessageInfo

func (m *CreatePortRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreatePortRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreatePortRequest) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *CreatePortRequest) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *CreatePortRequest) GetAlias() []string {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *CreatePortRequest) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *CreatePortRequest) GetCoordinates() []string {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func (m *CreatePortRequest) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *CreatePortRequest) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *CreatePortRequest) GetUnlocs() []string {
	if m != nil {
		return m.Unlocs
	}
	return nil
}

func (m *CreatePortRequest) GetCode() []string {
	if m != nil {
		return m.Code
	}
	return nil
}

type UpdatePortRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country              string   `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Alias                []string `protobuf:"bytes,5,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions              []string `protobuf:"bytes,6,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates          []string `protobuf:"bytes,7,rep,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province             string   `protobuf:"bytes,8,opt,name=province,proto3" json:"province,omitempty"`
	Timezone             string   `protobuf:"bytes,9,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs               []string `protobuf:"bytes,10,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code                 []string `protobuf:"bytes,11,rep,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePortRequest) Reset()         { *m = UpdatePortRequest{} }
func (m *UpdatePortRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePortRequest) ProtoMessage()    {}
func (*UpdatePortRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d80a8240fd02b040, []int{3}
}

func (m *UpdatePortRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePortRequest.Unmarshal(m, b)
}
func (m *UpdatePortRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePortRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePortRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePortRequest.Merge(m, src)
}
func (m *UpdatePortRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePortRequest.Size(m)
}
func (m *UpdatePortRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePortRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePortRequest proto.InternalMessageInfo

func (m *UpdatePortRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdatePortRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdatePortRequest) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *UpdatePortRequest) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *UpdatePortRequest) GetAlias() []string {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *UpdatePortRequest) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *UpdatePortRequest) GetCoordinates() []string {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func (m *UpdatePortRequest) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *UpdatePortRequest) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *UpdatePortRequest) GetUnlocs() []string {
	if m != nil {
		return m.Unlocs
	}
	return nil
}

func (m *UpdatePortRequest) GetCode() []string {
	if m != nil {
		return m.Code
	}
	return nil
}

type PortResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Data                 *Port    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortResponse) Reset()         { *m = PortResponse{} }
func (m *PortResponse) String() string { return proto.CompactTextString(m) }
func (*PortResponse) ProtoMessage()    {}
func (*PortResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d80a8240fd02b040, []int{4}
}

func (m *PortResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortResponse.Unmarshal(m, b)
}
func (m *PortResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortResponse.Marshal(b, m, deterministic)
}
func (m *PortResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortResponse.Merge(m, src)
}
func (m *PortResponse) XXX_Size() int {
	return xxx_messageInfo_PortResponse.Size(m)
}
func (m *PortResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PortResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PortResponse proto.InternalMessageInfo

func (m *PortResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *PortResponse) GetData() *Port {
	if m != nil {
		return m.Data
	}
	return nil
}

type PortsResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Data                 []*Port  `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortsResponse) Reset()         { *m = PortsResponse{} }
func (m *PortsResponse) String() string { return proto.CompactTextString(m) }
func (*PortsResponse) ProtoMessage()    {}
func (*PortsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d80a8240fd02b040, []int{5}
}

func (m *PortsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortsResponse.Unmarshal(m, b)
}
func (m *PortsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortsResponse.Marshal(b, m, deterministic)
}
func (m *PortsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortsResponse.Merge(m, src)
}
func (m *PortsResponse) XXX_Size() int {
	return xxx_messageInfo_PortsResponse.Size(m)
}
func (m *PortsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PortsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PortsResponse proto.InternalMessageInfo

func (m *PortsResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *PortsResponse) GetData() []*Port {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeleteResponse struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d80a8240fd02b040, []int{6}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*Port)(nil), "sync.Port")
	proto.RegisterType((*RequestWithId)(nil), "sync.RequestWithId")
	proto.RegisterType((*CreatePortRequest)(nil), "sync.CreatePortRequest")
	proto.RegisterType((*UpdatePortRequest)(nil), "sync.UpdatePortRequest")
	proto.RegisterType((*PortResponse)(nil), "sync.PortResponse")
	proto.RegisterType((*PortsResponse)(nil), "sync.PortsResponse")
	proto.RegisterType((*DeleteResponse)(nil), "sync.DeleteResponse")
}

func init() {
	proto.RegisterFile("ports.proto", fileDescriptor_d80a8240fd02b040)
}

var fileDescriptor_d80a8240fd02b040 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x54, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x6d, 0xb3, 0xd9, 0x8f, 0x4e, 0x68, 0x25, 0xdc, 0xaa, 0x58, 0x41, 0x82, 0x55, 0x4e, 0x3d,
	0xa5, 0x52, 0xe1, 0x82, 0x7a, 0xe3, 0xab, 0xe2, 0x86, 0xb6, 0x42, 0x48, 0xdc, 0x42, 0x32, 0x2c,
	0x96, 0xb2, 0x76, 0xb0, 0x27, 0x95, 0xc2, 0xef, 0xe0, 0x0f, 0xf0, 0x1b, 0xf8, 0x83, 0xc8, 0xe3,
	0x84, 0xdd, 0x55, 0x59, 0x50, 0xef, 0x3d, 0xc5, 0xef, 0x79, 0xde, 0x73, 0xfc, 0xe4, 0x19, 0x48,
	0x1a, 0x63, 0xc9, 0xe5, 0x8d, 0x35, 0x64, 0x44, 0xec, 0x3a, 0x5d, 0xa6, 0x8f, 0x97, 0xc6, 0x2c,
	0x6b, 0x3c, 0x67, 0xee, 0x73, 0xfb, 0xe5, 0x1c, 0x57, 0x0d, 0x75, 0xa1, 0x24, 0xfb, 0x11, 0x41,
	0xfc, 0xde, 0x58, 0x12, 0x27, 0x30, 0x26, 0x45, 0x35, 0xca, 0xfd, 0xf9, 0xfe, 0xd9, 0xc1, 0x22,
	0x00, 0x21, 0x20, 0xd6, 0xc5, 0x0a, 0x65, 0xc4, 0x24, 0xaf, 0x3d, 0x57, 0x2a, 0xea, 0xe4, 0x28,
	0x70, 0x7e, 0x2d, 0x24, 0x4c, 0x4b, 0xd3, 0x6a, 0xb2, 0x9d, 0x8c, 0x99, 0x1e, 0xa0, 0xf7, 0x2d,
	0x6a, 0x55, 0x38, 0x39, 0x9e, 0x8f, 0xbc, 0x2f, 0x03, 0x5f, 0x6f, 0x71, 0xa9, 0x8c, 0x76, 0x72,
	0xc2, 0xfc, 0x00, 0xc5, 0x1c, 0x92, 0xd2, 0x18, 0x5b, 0x29, 0x5d, 0x10, 0x3a, 0x39, 0xe5, 0xdd,
	0x4d, 0x4a, 0xa4, 0x30, 0x6b, 0xac, 0xb9, 0x51, 0xba, 0x44, 0x39, 0xe3, 0xc3, 0xfe, 0x60, 0xbf,
	0x47, 0x6a, 0x85, 0xdf, 0x8d, 0x46, 0x79, 0x10, 0xf6, 0x06, 0x2c, 0x4e, 0x61, 0xd2, 0xea, 0xda,
	0x94, 0x4e, 0x02, 0x9b, 0xf6, 0x88, 0xef, 0x63, 0x2a, 0x94, 0x09, 0xb3, 0xbc, 0xce, 0x9e, 0xc2,
	0xe1, 0x02, 0xbf, 0xb5, 0xe8, 0xe8, 0xa3, 0xa2, 0xaf, 0xef, 0x2a, 0x71, 0x04, 0x91, 0xaa, 0x38,
	0x9b, 0xf1, 0x22, 0x52, 0x55, 0xf6, 0x33, 0x82, 0x87, 0xaf, 0x2c, 0x16, 0x84, 0x3e, 0xbd, 0xbe,
	0xf6, 0x3e, 0xc4, 0xad, 0x10, 0x7d, 0x46, 0x1f, 0x9a, 0xea, 0x3e, 0xa3, 0x7f, 0x64, 0xf4, 0x16,
	0x1e, 0x84, 0x70, 0x5c, 0x63, 0xb4, 0x63, 0xad, 0xa3, 0x82, 0x5a, 0xd7, 0xbf, 0xb5, 0x1e, 0x89,
	0x27, 0x10, 0x57, 0x05, 0x15, 0x9c, 0x4f, 0x72, 0x01, 0xb9, 0xef, 0xec, 0x9c, 0x95, 0xcc, 0x67,
	0x57, 0x70, 0xe8, 0x91, 0xbb, 0x83, 0xd1, 0xe8, 0xaf, 0x46, 0x67, 0x70, 0xf4, 0x1a, 0x6b, 0x24,
	0xfc, 0x9f, 0xd3, 0xc5, 0xaf, 0x08, 0x92, 0xeb, 0x4e, 0x97, 0xd7, 0x68, 0x6f, 0x54, 0x89, 0xe2,
	0x39, 0x4c, 0xaf, 0x90, 0x78, 0x98, 0x1c, 0x07, 0xdb, 0xad, 0x16, 0x4a, 0xc5, 0xc6, 0x59, 0xbd,
	0x77, 0xb6, 0x27, 0x5e, 0xc0, 0xac, 0x57, 0x39, 0x71, 0x9a, 0x87, 0x51, 0x95, 0x0f, 0xa3, 0x2a,
	0x7f, 0xe3, 0x47, 0x55, 0x7a, 0xbc, 0x56, 0xba, 0x0d, 0xe9, 0x25, 0xc0, 0xba, 0x05, 0xc5, 0xa3,
	0x50, 0x74, 0xab, 0x29, 0x77, 0x9c, 0x7b, 0x09, 0xb0, 0x7e, 0x9b, 0x83, 0xf8, 0xd6, 0x6b, 0xdd,
	0xf9, 0xd3, 0x10, 0x42, 0xda, 0x7d, 0xdb, 0x93, 0x40, 0x6e, 0x67, 0x99, 0xed, 0xbd, 0x9c, 0x7e,
	0x1a, 0x87, 0xdb, 0x4d, 0xf8, 0xf3, 0xec, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x3b, 0x7f,
	0xff, 0xb2, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SyncServiceClient is the client API for SyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SyncServiceClient interface {
	GetPort(ctx context.Context, in *RequestWithId, opts ...grpc.CallOption) (*PortResponse, error)
	GetPorts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PortsResponse, error)
	CreatePort(ctx context.Context, in *CreatePortRequest, opts ...grpc.CallOption) (*PortResponse, error)
	UpdatePort(ctx context.Context, in *UpdatePortRequest, opts ...grpc.CallOption) (*PortResponse, error)
	DeletePort(ctx context.Context, in *RequestWithId, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type syncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncServiceClient(cc grpc.ClientConnInterface) SyncServiceClient {
	return &syncServiceClient{cc}
}

func (c *syncServiceClient) GetPort(ctx context.Context, in *RequestWithId, opts ...grpc.CallOption) (*PortResponse, error) {
	out := new(PortResponse)
	err := c.cc.Invoke(ctx, "/sync.SyncService/GetPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) GetPorts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PortsResponse, error) {
	out := new(PortsResponse)
	err := c.cc.Invoke(ctx, "/sync.SyncService/GetPorts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) CreatePort(ctx context.Context, in *CreatePortRequest, opts ...grpc.CallOption) (*PortResponse, error) {
	out := new(PortResponse)
	err := c.cc.Invoke(ctx, "/sync.SyncService/CreatePort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) UpdatePort(ctx context.Context, in *UpdatePortRequest, opts ...grpc.CallOption) (*PortResponse, error) {
	out := new(PortResponse)
	err := c.cc.Invoke(ctx, "/sync.SyncService/UpdatePort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) DeletePort(ctx context.Context, in *RequestWithId, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/sync.SyncService/DeletePort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncServiceServer is the server API for SyncService service.
type SyncServiceServer interface {
	GetPort(context.Context, *RequestWithId) (*PortResponse, error)
	GetPorts(context.Context, *empty.Empty) (*PortsResponse, error)
	CreatePort(context.Context, *CreatePortRequest) (*PortResponse, error)
	UpdatePort(context.Context, *UpdatePortRequest) (*PortResponse, error)
	DeletePort(context.Context, *RequestWithId) (*DeleteResponse, error)
}

// UnimplementedSyncServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSyncServiceServer struct {
}

func (*UnimplementedSyncServiceServer) GetPort(ctx context.Context, req *RequestWithId) (*PortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPort not implemented")
}
func (*UnimplementedSyncServiceServer) GetPorts(ctx context.Context, req *empty.Empty) (*PortsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPorts not implemented")
}
func (*UnimplementedSyncServiceServer) CreatePort(ctx context.Context, req *CreatePortRequest) (*PortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePort not implemented")
}
func (*UnimplementedSyncServiceServer) UpdatePort(ctx context.Context, req *UpdatePortRequest) (*PortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePort not implemented")
}
func (*UnimplementedSyncServiceServer) DeletePort(ctx context.Context, req *RequestWithId) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePort not implemented")
}

func RegisterSyncServiceServer(s *grpc.Server, srv SyncServiceServer) {
	s.RegisterService(&_SyncService_serviceDesc, srv)
}

func _SyncService_GetPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestWithId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sync.SyncService/GetPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetPort(ctx, req.(*RequestWithId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_GetPorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetPorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sync.SyncService/GetPorts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetPorts(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_CreatePort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).CreatePort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sync.SyncService/CreatePort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).CreatePort(ctx, req.(*CreatePortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_UpdatePort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).UpdatePort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sync.SyncService/UpdatePort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).UpdatePort(ctx, req.(*UpdatePortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_DeletePort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestWithId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).DeletePort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sync.SyncService/DeletePort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).DeletePort(ctx, req.(*RequestWithId))
	}
	return interceptor(ctx, in, info, handler)
}

var _SyncService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sync.SyncService",
	HandlerType: (*SyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPort",
			Handler:    _SyncService_GetPort_Handler,
		},
		{
			MethodName: "GetPorts",
			Handler:    _SyncService_GetPorts_Handler,
		},
		{
			MethodName: "CreatePort",
			Handler:    _SyncService_CreatePort_Handler,
		},
		{
			MethodName: "UpdatePort",
			Handler:    _SyncService_UpdatePort_Handler,
		},
		{
			MethodName: "DeletePort",
			Handler:    _SyncService_DeletePort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ports.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment.proto

package go_micro_srv_proto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
// const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 货轮承运的一批货物
type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_774f0b29a0c4f8e9, []int{0}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

// 单个集装箱
type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_774f0b29a0c4f8e9, []int{1}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// 托运结果
type Response struct {
	Created              bool           `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment   `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	Consignments         []*Consignment `protobuf:"bytes,3,rep,name=consignments,proto3" json:"consignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_774f0b29a0c4f8e9, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

// 查看货物信息的请求
// 客户端想要从服务端请求数据，必须有请求格式，哪怕为空
type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_774f0b29a0c4f8e9, []int{3}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Consignment)(nil), "proto.Consignment")
	proto.RegisterType((*Container)(nil), "proto.Container")
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterType((*GetRequest)(nil), "proto.GetRequest")
}

func init() { proto.RegisterFile("proto/consignment.proto", fileDescriptor_774f0b29a0c4f8e9) }

var fileDescriptor_774f0b29a0c4f8e9 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4b, 0x4e, 0xeb, 0x30,
	0x14, 0x7d, 0xe9, 0x3f, 0x37, 0xd5, 0xeb, 0xeb, 0x1d, 0xbc, 0x5a, 0x30, 0x20, 0xca, 0xa8, 0xa3,
	0x82, 0x0a, 0x02, 0x89, 0x69, 0x07, 0x55, 0xa7, 0xee, 0x02, 0x50, 0x89, 0xaf, 0x52, 0x4b, 0xd4,
	0x0e, 0xb6, 0x5b, 0x76, 0xc0, 0x02, 0x58, 0x04, 0xeb, 0x44, 0x75, 0x12, 0x6a, 0x44, 0x47, 0xd1,
	0xf9, 0xe5, 0x1c, 0xdb, 0x30, 0x29, 0x8d, 0x76, 0xfa, 0x3a, 0xd7, 0xca, 0xca, 0x42, 0xed, 0x48,
	0xb9, 0x99, 0x67, 0xb0, 0xeb, 0x3f, 0xd9, 0x67, 0x04, 0xc9, 0xe2, 0x24, 0xe2, 0x5f, 0x68, 0x49,
	0xc1, 0xa2, 0x34, 0x9a, 0xc6, 0xbc, 0x25, 0x05, 0xa6, 0x90, 0x08, 0xb2, 0xb9, 0x91, 0xa5, 0x93,
	0x5a, 0xb1, 0x96, 0x17, 0x42, 0x0a, 0xff, 0x43, 0xef, 0x8d, 0x64, 0xb1, 0x75, 0xac, 0x9d, 0x46,
	0xd3, 0x2e, 0xaf, 0x11, 0xde, 0x00, 0xe4, 0x5a, 0xb9, 0x8d, 0x54, 0x64, 0x2c, 0xeb, 0xa4, 0xed,
	0x69, 0x32, 0xff, 0x57, 0x95, 0xcf, 0x16, 0x8d, 0xc0, 0x03, 0x0f, 0x5e, 0x42, 0x7c, 0x20, 0x6b,
	0xe9, 0xe5, 0x49, 0x0a, 0xd6, 0xf5, 0x4d, 0x83, 0x8a, 0x58, 0x89, 0x6c, 0x07, 0xf1, 0x77, 0xea,
	0xd7, 0xca, 0x2b, 0x48, 0xf2, 0xbd, 0x75, 0x7a, 0x47, 0xe6, 0x98, 0xad, 0x56, 0x42, 0x43, 0xad,
	0xc4, 0x71, 0xa4, 0x36, 0xb2, 0x90, 0xca, 0x8f, 0x8c, 0x79, 0x8d, 0x70, 0x02, 0xfd, 0xbd, 0xad,
	0x42, 0x9d, 0x4a, 0x38, 0xc2, 0x95, 0xc8, 0x3e, 0x22, 0x18, 0x70, 0xb2, 0xa5, 0x56, 0x96, 0x90,
	0x41, 0x3f, 0x37, 0xb4, 0x71, 0x54, 0x75, 0x0e, 0x78, 0x03, 0xf1, 0x0e, 0x92, 0xe0, 0x6a, 0x7d,
	0x71, 0x32, 0xc7, 0xd3, 0x29, 0x1b, 0x85, 0x87, 0x36, 0xbc, 0x87, 0x61, 0x00, 0x2d, 0x6b, 0xfb,
	0xcb, 0x39, 0x17, 0xfb, 0xe1, 0xcb, 0x86, 0x00, 0x4b, 0x72, 0x9c, 0x5e, 0xf7, 0x64, 0xdd, 0xfc,
	0x3d, 0x82, 0xd1, 0x7a, 0x2b, 0xcb, 0x52, 0xaa, 0x62, 0x4d, 0xe6, 0x20, 0x73, 0xc2, 0x47, 0x18,
	0x2f, 0xfc, 0xb4, 0xf0, 0x4d, 0xcf, 0xfc, 0xf8, 0x62, 0x54, 0x73, 0xcd, 0x19, 0xb3, 0x3f, 0xf8,
	0x00, 0xa3, 0x25, 0xb9, 0xc0, 0x64, 0x71, 0x5c, 0xbb, 0x4e, 0xad, 0x67, 0x82, 0xcf, 0x3d, 0xcf,
	0xdc, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x15, 0x1d, 0x1d, 0x16, 0x6c, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShippingService service

type ShippingServiceClient interface {
	// 托运一批货物
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	// 查看托运货物的信息
	GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	c           client.Client
	serviceName string
}

func NewShippingServiceClient(serviceName string, c client.Client) ShippingServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "proto"
	}
	return &shippingServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.CreateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.GetConsignments", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceHandler interface {
	// 托运一批货物
	CreateConsignment(context.Context, *Consignment, *Response) error
	// 查看托运货物的信息
	GetConsignments(context.Context, *GetRequest, *Response) error
}

func RegisterShippingServiceHandler(s server.Server, hdlr ShippingServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ShippingService{hdlr}, opts...))
}

type ShippingService struct {
	ShippingServiceHandler
}

func (h *ShippingService) CreateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ShippingServiceHandler.CreateConsignment(ctx, in, out)
}

func (h *ShippingService) GetConsignments(ctx context.Context, in *GetRequest, out *Response) error {
	return h.ShippingServiceHandler.GetConsignments(ctx, in, out)
}
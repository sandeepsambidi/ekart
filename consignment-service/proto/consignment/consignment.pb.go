// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

package shippy_service_consignment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

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
	return fileDescriptor_e5e5ab05dfa973d5, []int{0}
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
	return fileDescriptor_e5e5ab05dfa973d5, []int{1}
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
	return fileDescriptor_e5e5ab05dfa973d5, []int{2}
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

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{3}
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
	proto.RegisterType((*Consignment)(nil), "shippy.service.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "shippy.service.consignment.Container")
	proto.RegisterType((*Response)(nil), "shippy.service.consignment.Response")
	proto.RegisterType((*GetRequest)(nil), "shippy.service.consignment.GetRequest")
}

func init() {
	proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor_e5e5ab05dfa973d5)
}

var fileDescriptor_e5e5ab05dfa973d5 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdb, 0x4e, 0xe3, 0x40,
	0x0c, 0xdd, 0xf4, 0x1e, 0xa7, 0xda, 0x6a, 0xe7, 0x61, 0x37, 0xea, 0x3e, 0x10, 0x85, 0x5b, 0x9f,
	0x82, 0x54, 0x3e, 0xa1, 0x42, 0x55, 0xc4, 0xdb, 0xf4, 0x03, 0x50, 0xc9, 0x58, 0xe9, 0x48, 0x74,
	0x26, 0xc4, 0xd3, 0x22, 0xbe, 0x0d, 0xbe, 0x84, 0xaf, 0x41, 0x99, 0x34, 0x74, 0x10, 0xa2, 0xf4,
	0xcd, 0xc7, 0xf6, 0xf1, 0x39, 0xb6, 0x0c, 0xa7, 0x45, 0xa9, 0x8d, 0xbe, 0xca, 0xb4, 0x22, 0x99,
	0xab, 0x35, 0x2a, 0xe3, 0xc6, 0x89, 0xad, 0xb2, 0x31, 0xad, 0x64, 0x51, 0x3c, 0x27, 0x84, 0xe5,
	0x56, 0x66, 0x98, 0x38, 0x1d, 0xf1, 0x8b, 0x07, 0xc1, 0x6c, 0x8f, 0xd9, 0x6f, 0x68, 0x49, 0x11,
	0x7a, 0x91, 0x37, 0xf1, 0x79, 0x4b, 0x0a, 0x16, 0x41, 0x20, 0x90, 0xb2, 0x52, 0x16, 0x46, 0x6a,
	0x15, 0xb6, 0x6c, 0xc1, 0x4d, 0xb1, 0xbf, 0xd0, 0x7b, 0x42, 0x99, 0xaf, 0x4c, 0xd8, 0x8e, 0xbc,
	0x49, 0x97, 0xef, 0x10, 0xbb, 0x01, 0xc8, 0xb4, 0x32, 0x4b, 0xa9, 0xb0, 0xa4, 0xb0, 0x13, 0xb5,
	0x27, 0xc1, 0xf4, 0x3c, 0xf9, 0xde, 0x4a, 0x32, 0x6b, 0xba, 0xb9, 0x43, 0x64, 0xff, 0xc1, 0xdf,
	0x22, 0x11, 0x3e, 0xdc, 0x49, 0x11, 0x76, 0xad, 0xfc, 0xa0, 0x4e, 0xa4, 0x22, 0x5e, 0x83, 0xff,
	0xc1, 0xfa, 0x62, 0xfd, 0x04, 0x82, 0x6c, 0x43, 0x46, 0xaf, 0xb1, 0xac, 0xb8, 0xb5, 0x75, 0x68,
	0x52, 0xa9, 0xa8, 0x9c, 0xeb, 0x52, 0xe6, 0x52, 0x59, 0xe7, 0x3e, 0xdf, 0x21, 0xf6, 0x0f, 0xfa,
	0x1b, 0xaa, 0x49, 0x9d, 0xba, 0x50, 0xc1, 0x54, 0xc4, 0xaf, 0x1e, 0x0c, 0x38, 0x52, 0xa1, 0x15,
	0x21, 0x0b, 0xa1, 0x9f, 0x95, 0xb8, 0x34, 0x58, 0x6b, 0x0e, 0x78, 0x03, 0x59, 0x0a, 0x81, 0xb3,
	0x97, 0x15, 0x0e, 0xa6, 0x97, 0x3f, 0xac, 0xde, 0xc4, 0xdc, 0xe5, 0xb2, 0x5b, 0x18, 0x3a, 0x90,
	0xc2, 0xb6, 0x3d, 0xe3, 0xd1, 0xb3, 0x3e, 0x91, 0xe3, 0x21, 0xc0, 0x1c, 0x0d, 0xc7, 0xc7, 0x0d,
	0x92, 0x99, 0xbe, 0x79, 0x30, 0x5a, 0x54, 0x63, 0xa4, 0xca, 0x17, 0xf5, 0x1c, 0x26, 0xe0, 0xcf,
	0xcc, 0x2e, 0xe1, 0xbe, 0xc4, 0xb1, 0x6a, 0xe3, 0xb3, 0x43, 0x8d, 0xcd, 0xdd, 0xe2, 0x5f, 0x6c,
	0x09, 0xa3, 0x39, 0x1a, 0x87, 0x49, 0xec, 0xe2, 0x10, 0x75, 0x6f, 0xfa, 0x58, 0x89, 0xfb, 0x9e,
	0xfd, 0xfc, 0xeb, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x4c, 0xb4, 0x37, 0x20, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShippingService service

type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
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
		serviceName = "shippy.service.consignment"
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
	CreateConsignment(context.Context, *Consignment, *Response) error
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

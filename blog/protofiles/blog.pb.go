// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog/protofiles/blog.proto

package blogpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type EntryRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Year                 string   `protobuf:"bytes,4,opt,name=year,proto3" json:"year,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntryRequest) Reset()         { *m = EntryRequest{} }
func (m *EntryRequest) String() string { return proto.CompactTextString(m) }
func (*EntryRequest) ProtoMessage()    {}
func (*EntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5480b3f1960d8f24, []int{0}
}

func (m *EntryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntryRequest.Unmarshal(m, b)
}
func (m *EntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntryRequest.Marshal(b, m, deterministic)
}
func (m *EntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntryRequest.Merge(m, src)
}
func (m *EntryRequest) XXX_Size() int {
	return xxx_messageInfo_EntryRequest.Size(m)
}
func (m *EntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EntryRequest proto.InternalMessageInfo

func (m *EntryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EntryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EntryRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *EntryRequest) GetYear() string {
	if m != nil {
		return m.Year
	}
	return ""
}

type CreateEntryRequest struct {
	EntryRequest         *EntryRequest `protobuf:"bytes,1,opt,name=entry_request,json=entryRequest,proto3" json:"entry_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateEntryRequest) Reset()         { *m = CreateEntryRequest{} }
func (m *CreateEntryRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEntryRequest) ProtoMessage()    {}
func (*CreateEntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5480b3f1960d8f24, []int{1}
}

func (m *CreateEntryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntryRequest.Unmarshal(m, b)
}
func (m *CreateEntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntryRequest.Marshal(b, m, deterministic)
}
func (m *CreateEntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntryRequest.Merge(m, src)
}
func (m *CreateEntryRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEntryRequest.Size(m)
}
func (m *CreateEntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntryRequest proto.InternalMessageInfo

func (m *CreateEntryRequest) GetEntryRequest() *EntryRequest {
	if m != nil {
		return m.EntryRequest
	}
	return nil
}

type ResponseCreateEntry struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseCreateEntry) Reset()         { *m = ResponseCreateEntry{} }
func (m *ResponseCreateEntry) String() string { return proto.CompactTextString(m) }
func (*ResponseCreateEntry) ProtoMessage()    {}
func (*ResponseCreateEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_5480b3f1960d8f24, []int{2}
}

func (m *ResponseCreateEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseCreateEntry.Unmarshal(m, b)
}
func (m *ResponseCreateEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseCreateEntry.Marshal(b, m, deterministic)
}
func (m *ResponseCreateEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseCreateEntry.Merge(m, src)
}
func (m *ResponseCreateEntry) XXX_Size() int {
	return xxx_messageInfo_ResponseCreateEntry.Size(m)
}
func (m *ResponseCreateEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseCreateEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseCreateEntry proto.InternalMessageInfo

func (m *ResponseCreateEntry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadEntryRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadEntryRequest) Reset()         { *m = ReadEntryRequest{} }
func (m *ReadEntryRequest) String() string { return proto.CompactTextString(m) }
func (*ReadEntryRequest) ProtoMessage()    {}
func (*ReadEntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5480b3f1960d8f24, []int{3}
}

func (m *ReadEntryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadEntryRequest.Unmarshal(m, b)
}
func (m *ReadEntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadEntryRequest.Marshal(b, m, deterministic)
}
func (m *ReadEntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadEntryRequest.Merge(m, src)
}
func (m *ReadEntryRequest) XXX_Size() int {
	return xxx_messageInfo_ReadEntryRequest.Size(m)
}
func (m *ReadEntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadEntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadEntryRequest proto.InternalMessageInfo

func (m *ReadEntryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadEntryResponse struct {
	EntryRequest         *EntryRequest `protobuf:"bytes,1,opt,name=entry_request,json=entryRequest,proto3" json:"entry_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReadEntryResponse) Reset()         { *m = ReadEntryResponse{} }
func (m *ReadEntryResponse) String() string { return proto.CompactTextString(m) }
func (*ReadEntryResponse) ProtoMessage()    {}
func (*ReadEntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5480b3f1960d8f24, []int{4}
}

func (m *ReadEntryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadEntryResponse.Unmarshal(m, b)
}
func (m *ReadEntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadEntryResponse.Marshal(b, m, deterministic)
}
func (m *ReadEntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadEntryResponse.Merge(m, src)
}
func (m *ReadEntryResponse) XXX_Size() int {
	return xxx_messageInfo_ReadEntryResponse.Size(m)
}
func (m *ReadEntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadEntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadEntryResponse proto.InternalMessageInfo

func (m *ReadEntryResponse) GetEntryRequest() *EntryRequest {
	if m != nil {
		return m.EntryRequest
	}
	return nil
}

type UpdateEntryRequest struct {
	EntryRequest         *EntryRequest `protobuf:"bytes,1,opt,name=entry_request,json=entryRequest,proto3" json:"entry_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdateEntryRequest) Reset()         { *m = UpdateEntryRequest{} }
func (m *UpdateEntryRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateEntryRequest) ProtoMessage()    {}
func (*UpdateEntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5480b3f1960d8f24, []int{5}
}

func (m *UpdateEntryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEntryRequest.Unmarshal(m, b)
}
func (m *UpdateEntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEntryRequest.Marshal(b, m, deterministic)
}
func (m *UpdateEntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEntryRequest.Merge(m, src)
}
func (m *UpdateEntryRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateEntryRequest.Size(m)
}
func (m *UpdateEntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEntryRequest proto.InternalMessageInfo

func (m *UpdateEntryRequest) GetEntryRequest() *EntryRequest {
	if m != nil {
		return m.EntryRequest
	}
	return nil
}

type UpdateEntryResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEntryResponse) Reset()         { *m = UpdateEntryResponse{} }
func (m *UpdateEntryResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateEntryResponse) ProtoMessage()    {}
func (*UpdateEntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5480b3f1960d8f24, []int{6}
}

func (m *UpdateEntryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEntryResponse.Unmarshal(m, b)
}
func (m *UpdateEntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEntryResponse.Marshal(b, m, deterministic)
}
func (m *UpdateEntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEntryResponse.Merge(m, src)
}
func (m *UpdateEntryResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateEntryResponse.Size(m)
}
func (m *UpdateEntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEntryResponse proto.InternalMessageInfo

func (m *UpdateEntryResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type DeleteEntryRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEntryRequest) Reset()         { *m = DeleteEntryRequest{} }
func (m *DeleteEntryRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteEntryRequest) ProtoMessage()    {}
func (*DeleteEntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5480b3f1960d8f24, []int{7}
}

func (m *DeleteEntryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEntryRequest.Unmarshal(m, b)
}
func (m *DeleteEntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEntryRequest.Marshal(b, m, deterministic)
}
func (m *DeleteEntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEntryRequest.Merge(m, src)
}
func (m *DeleteEntryRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteEntryRequest.Size(m)
}
func (m *DeleteEntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEntryRequest proto.InternalMessageInfo

func (m *DeleteEntryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteEntryResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEntryResponse) Reset()         { *m = DeleteEntryResponse{} }
func (m *DeleteEntryResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteEntryResponse) ProtoMessage()    {}
func (*DeleteEntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5480b3f1960d8f24, []int{8}
}

func (m *DeleteEntryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEntryResponse.Unmarshal(m, b)
}
func (m *DeleteEntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEntryResponse.Marshal(b, m, deterministic)
}
func (m *DeleteEntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEntryResponse.Merge(m, src)
}
func (m *DeleteEntryResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteEntryResponse.Size(m)
}
func (m *DeleteEntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEntryResponse proto.InternalMessageInfo

func (m *DeleteEntryResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*EntryRequest)(nil), "EntryRequest")
	proto.RegisterType((*CreateEntryRequest)(nil), "CreateEntryRequest")
	proto.RegisterType((*ResponseCreateEntry)(nil), "ResponseCreateEntry")
	proto.RegisterType((*ReadEntryRequest)(nil), "ReadEntryRequest")
	proto.RegisterType((*ReadEntryResponse)(nil), "ReadEntryResponse")
	proto.RegisterType((*UpdateEntryRequest)(nil), "UpdateEntryRequest")
	proto.RegisterType((*UpdateEntryResponse)(nil), "UpdateEntryResponse")
	proto.RegisterType((*DeleteEntryRequest)(nil), "DeleteEntryRequest")
	proto.RegisterType((*DeleteEntryResponse)(nil), "DeleteEntryResponse")
}

func init() {
	proto.RegisterFile("blog/protofiles/blog.proto", fileDescriptor_5480b3f1960d8f24)
}

var fileDescriptor_5480b3f1960d8f24 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0xa5, 0xb5, 0x14, 0x3b, 0x69, 0xc5, 0xee, 0x16, 0x09, 0x39, 0xc9, 0xa2, 0xe0, 0xa5, 0x5b,
	0x88, 0x97, 0x9c, 0xab, 0x62, 0xcf, 0x11, 0x2f, 0xbd, 0x48, 0x62, 0xc6, 0x12, 0x88, 0x49, 0xdc,
	0xdd, 0x08, 0xf9, 0xd5, 0xfe, 0x05, 0xd9, 0x24, 0xea, 0xe6, 0x03, 0x0f, 0xe2, 0x6d, 0xde, 0x63,
	0x78, 0x33, 0xef, 0x0d, 0x03, 0x4e, 0x98, 0x64, 0x87, 0x4d, 0x2e, 0x32, 0x95, 0xbd, 0xc4, 0x09,
	0xca, 0x8d, 0xc6, 0xbc, 0xc2, 0x6c, 0x0f, 0xf3, 0xbb, 0x54, 0x89, 0xd2, 0xc7, 0xb7, 0x02, 0xa5,
	0x22, 0x27, 0x30, 0x8e, 0x23, 0x7b, 0x74, 0x3e, 0xba, 0x9a, 0xf9, 0xe3, 0x38, 0x22, 0x04, 0x26,
	0x69, 0xf0, 0x8a, 0xf6, 0xb8, 0x62, 0xaa, 0x5a, 0x73, 0xaa, 0xcc, 0xd1, 0x3e, 0xaa, 0x39, 0x5d,
	0x6b, 0xae, 0xc4, 0x40, 0xd8, 0x93, 0x9a, 0xd3, 0x35, 0xdb, 0x01, 0xb9, 0x11, 0x18, 0x28, 0x6c,
	0x4d, 0x70, 0x61, 0x81, 0x1a, 0x3f, 0x89, 0x9a, 0xa8, 0x86, 0x59, 0xee, 0x82, 0x9b, 0x5d, 0xfe,
	0x1c, 0x0d, 0xc4, 0x2e, 0x81, 0xfa, 0x28, 0xf3, 0x2c, 0x95, 0x68, 0x28, 0x76, 0x97, 0x65, 0x0c,
	0x4e, 0x7d, 0x0c, 0xa2, 0xdf, 0x0c, 0xb1, 0x7b, 0x58, 0x1a, 0x3d, 0xb5, 0xe6, 0x9f, 0x76, 0xda,
	0x01, 0x79, 0xcc, 0xa3, 0xff, 0x70, 0xb7, 0x06, 0xda, 0x52, 0x6a, 0x96, 0x3a, 0x83, 0xa9, 0x54,
	0x81, 0x2a, 0x64, 0xb3, 0x7d, 0x83, 0xd8, 0x05, 0x90, 0x5b, 0x4c, 0xb0, 0x33, 0xb8, 0xeb, 0x73,
	0x0d, 0xb4, 0xd5, 0xf5, 0x23, 0x2a, 0x50, 0x16, 0x89, 0xfa, 0x12, 0xad, 0x91, 0xfb, 0x31, 0x02,
	0x6b, 0x9b, 0x64, 0x87, 0x07, 0x14, 0xef, 0xf1, 0x33, 0x12, 0x0f, 0x2c, 0x33, 0x69, 0xca, 0xfb,
	0x97, 0x74, 0x56, 0x7c, 0xe8, 0x28, 0x2e, 0xcc, 0xbe, 0x03, 0x26, 0x4b, 0xde, 0x3d, 0x88, 0x43,
	0x78, 0x3f, 0x7f, 0x0f, 0x2c, 0x23, 0x01, 0x42, 0x79, 0x3f, 0x59, 0x67, 0xc5, 0x87, 0x42, 0xf2,
	0xc0, 0x32, 0x6c, 0x12, 0xca, 0xfb, 0xd1, 0x38, 0x2b, 0x3e, 0x90, 0xc4, 0xf6, 0x78, 0x3f, 0xd5,
	0x7f, 0x90, 0x87, 0xe1, 0xb4, 0x7a, 0x85, 0xeb, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x83, 0xbe,
	0x2b, 0x22, 0x28, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServiceClient interface {
	CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*ResponseCreateEntry, error)
	ReadEntry(ctx context.Context, in *ReadEntryRequest, opts ...grpc.CallOption) (*ReadEntryResponse, error)
	UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*UpdateEntryResponse, error)
	DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*DeleteEntryResponse, error)
}

type blogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogServiceClient(cc grpc.ClientConnInterface) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateEntry(ctx context.Context, in *CreateEntryRequest, opts ...grpc.CallOption) (*ResponseCreateEntry, error) {
	out := new(ResponseCreateEntry)
	err := c.cc.Invoke(ctx, "/BlogService/CreateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ReadEntry(ctx context.Context, in *ReadEntryRequest, opts ...grpc.CallOption) (*ReadEntryResponse, error) {
	out := new(ReadEntryResponse)
	err := c.cc.Invoke(ctx, "/BlogService/ReadEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) UpdateEntry(ctx context.Context, in *UpdateEntryRequest, opts ...grpc.CallOption) (*UpdateEntryResponse, error) {
	out := new(UpdateEntryResponse)
	err := c.cc.Invoke(ctx, "/BlogService/UpdateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*DeleteEntryResponse, error) {
	out := new(DeleteEntryResponse)
	err := c.cc.Invoke(ctx, "/BlogService/DeleteEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
type BlogServiceServer interface {
	CreateEntry(context.Context, *CreateEntryRequest) (*ResponseCreateEntry, error)
	ReadEntry(context.Context, *ReadEntryRequest) (*ReadEntryResponse, error)
	UpdateEntry(context.Context, *UpdateEntryRequest) (*UpdateEntryResponse, error)
	DeleteEntry(context.Context, *DeleteEntryRequest) (*DeleteEntryResponse, error)
}

// UnimplementedBlogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (*UnimplementedBlogServiceServer) CreateEntry(ctx context.Context, req *CreateEntryRequest) (*ResponseCreateEntry, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEntry not implemented")
}
func (*UnimplementedBlogServiceServer) ReadEntry(ctx context.Context, req *ReadEntryRequest) (*ReadEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEntry not implemented")
}
func (*UnimplementedBlogServiceServer) UpdateEntry(ctx context.Context, req *UpdateEntryRequest) (*UpdateEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEntry not implemented")
}
func (*UnimplementedBlogServiceServer) DeleteEntry(ctx context.Context, req *DeleteEntryRequest) (*DeleteEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEntry not implemented")
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_CreateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlogService/CreateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateEntry(ctx, req.(*CreateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ReadEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ReadEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlogService/ReadEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ReadEntry(ctx, req.(*ReadEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlogService/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).UpdateEntry(ctx, req.(*UpdateEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_DeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).DeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BlogService/DeleteEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).DeleteEntry(ctx, req.(*DeleteEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEntry",
			Handler:    _BlogService_CreateEntry_Handler,
		},
		{
			MethodName: "ReadEntry",
			Handler:    _BlogService_ReadEntry_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _BlogService_UpdateEntry_Handler,
		},
		{
			MethodName: "DeleteEntry",
			Handler:    _BlogService_DeleteEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blog/protofiles/blog.proto",
}

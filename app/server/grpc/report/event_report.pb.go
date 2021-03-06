// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/event_report.proto

package report

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

type RespCode int32

const (
	RespCode_Success       RespCode = 0
	RespCode_BasRequest    RespCode = 400
	RespCode_InternalError RespCode = 500
	RespCode_KeepOff       RespCode = 600
	RespCode_Silence       RespCode = 601
	RespCode_Pause         RespCode = 602
	RespCode_Active        RespCode = 603
)

var RespCode_name = map[int32]string{
	0:   "Success",
	400: "BasRequest",
	500: "InternalError",
	600: "KeepOff",
	601: "Silence",
	602: "Pause",
	603: "Active",
}

var RespCode_value = map[string]int32{
	"Success":       0,
	"BasRequest":    400,
	"InternalError": 500,
	"KeepOff":       600,
	"Silence":       601,
	"Pause":         602,
	"Active":        603,
}

func (x RespCode) String() string {
	return proto.EnumName(RespCode_name, int32(x))
}

func (RespCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f6fc8250ac85f127, []int{0}
}

type ReportRequest struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportRequest) Reset()         { *m = ReportRequest{} }
func (m *ReportRequest) String() string { return proto.CompactTextString(m) }
func (*ReportRequest) ProtoMessage()    {}
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6fc8250ac85f127, []int{0}
}

func (m *ReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportRequest.Unmarshal(m, b)
}
func (m *ReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportRequest.Marshal(b, m, deterministic)
}
func (m *ReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportRequest.Merge(m, src)
}
func (m *ReportRequest) XXX_Size() int {
	return xxx_messageInfo_ReportRequest.Size(m)
}
func (m *ReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportRequest proto.InternalMessageInfo

func (m *ReportRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ReportRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ReportRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type ReportResponse struct {
	Code                 RespCode `protobuf:"varint,1,opt,name=code,proto3,enum=report.RespCode" json:"code,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportResponse) Reset()         { *m = ReportResponse{} }
func (m *ReportResponse) String() string { return proto.CompactTextString(m) }
func (*ReportResponse) ProtoMessage()    {}
func (*ReportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6fc8250ac85f127, []int{1}
}

func (m *ReportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportResponse.Unmarshal(m, b)
}
func (m *ReportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportResponse.Marshal(b, m, deterministic)
}
func (m *ReportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportResponse.Merge(m, src)
}
func (m *ReportResponse) XXX_Size() int {
	return xxx_messageInfo_ReportResponse.Size(m)
}
func (m *ReportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReportResponse proto.InternalMessageInfo

func (m *ReportResponse) GetCode() RespCode {
	if m != nil {
		return m.Code
	}
	return RespCode_Success
}

func (m *ReportResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type PollEvent struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PollEvent) Reset()         { *m = PollEvent{} }
func (m *PollEvent) String() string { return proto.CompactTextString(m) }
func (*PollEvent) ProtoMessage()    {}
func (*PollEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6fc8250ac85f127, []int{2}
}

func (m *PollEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollEvent.Unmarshal(m, b)
}
func (m *PollEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollEvent.Marshal(b, m, deterministic)
}
func (m *PollEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollEvent.Merge(m, src)
}
func (m *PollEvent) XXX_Size() int {
	return xxx_messageInfo_PollEvent.Size(m)
}
func (m *PollEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PollEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PollEvent proto.InternalMessageInfo

func (m *PollEvent) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *PollEvent) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("report.RespCode", RespCode_name, RespCode_value)
	proto.RegisterType((*ReportRequest)(nil), "report.ReportRequest")
	proto.RegisterType((*ReportResponse)(nil), "report.ReportResponse")
	proto.RegisterType((*PollEvent)(nil), "report.PollEvent")
}

func init() { proto.RegisterFile("proto/event_report.proto", fileDescriptor_f6fc8250ac85f127) }

var fileDescriptor_f6fc8250ac85f127 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x35, 0xcd, 0x9a, 0xda, 0xe9, 0x87, 0xeb, 0x80, 0x12, 0x8a, 0x87, 0x52, 0x3c, 0x14, 0x91,
	0x0a, 0xf5, 0xe0, 0xc1, 0x93, 0x2d, 0x3d, 0x88, 0x07, 0x4b, 0xea, 0x49, 0x0f, 0x12, 0xd3, 0x29,
	0x14, 0xd2, 0xdd, 0x74, 0x77, 0x5b, 0xf0, 0x5f, 0xf8, 0xf3, 0xfc, 0xf8, 0x1b, 0xfe, 0x00, 0xc9,
	0xa6, 0xb1, 0x35, 0x82, 0xb7, 0x79, 0xef, 0xc1, 0x9b, 0x37, 0x6f, 0xc0, 0x4f, 0x94, 0x34, 0xf2,
	0x9c, 0x56, 0x24, 0xcc, 0x93, 0xa2, 0x44, 0x2a, 0xd3, 0xb5, 0x14, 0x7a, 0x19, 0x6a, 0x3f, 0x42,
	0x3d, 0xb0, 0x53, 0x40, 0x8b, 0x25, 0x69, 0x83, 0x08, 0xcc, 0xbc, 0x24, 0xe4, 0x3b, 0x2d, 0xa7,
	0x53, 0x09, 0xec, 0x8c, 0x3e, 0x94, 0x23, 0x29, 0x0c, 0x09, 0xe3, 0x97, 0x2c, 0x9d, 0x43, 0x3c,
	0x86, 0x8a, 0x99, 0xcd, 0x49, 0x9b, 0x70, 0x9e, 0xf8, 0x6e, 0xcb, 0xe9, 0xb8, 0xc1, 0x86, 0x68,
	0xdf, 0x43, 0x23, 0x37, 0xd7, 0x89, 0x14, 0x9a, 0xf0, 0x04, 0x58, 0x24, 0x27, 0x99, 0x7b, 0xa3,
	0xc7, 0xbb, 0xeb, 0x4c, 0xa9, 0x3e, 0x90, 0x13, 0x0a, 0xac, 0xfa, 0xdb, 0xb5, 0x54, 0x74, 0x1d,
	0x40, 0x65, 0x24, 0xe3, 0x78, 0x98, 0x1e, 0xb5, 0x1d, 0xcd, 0xf9, 0x27, 0x5a, 0xd1, 0xe4, 0x74,
	0x01, 0x7b, 0xf9, 0x52, 0xac, 0x42, 0x79, 0xbc, 0x8c, 0x22, 0xd2, 0x9a, 0xef, 0xe0, 0x3e, 0x40,
	0x3f, 0xd4, 0xeb, 0x36, 0xf8, 0xab, 0x8b, 0x08, 0xf5, 0x1b, 0x61, 0x48, 0x89, 0x30, 0x1e, 0x2a,
	0x25, 0x15, 0xff, 0x72, 0xb1, 0x06, 0xe5, 0x5b, 0xa2, 0xe4, 0x6e, 0x3a, 0xe5, 0x6f, 0x2c, 0x45,
	0xe3, 0x59, 0x4c, 0x22, 0x22, 0xfe, 0xce, 0x10, 0x60, 0x77, 0x14, 0x2e, 0x35, 0xf1, 0x0f, 0x86,
	0x55, 0xf0, 0xae, 0x23, 0x33, 0x5b, 0x11, 0xff, 0x64, 0x3d, 0x03, 0x55, 0x9b, 0x79, 0x4c, 0x6a,
	0x45, 0x0a, 0x2f, 0xc1, 0xcb, 0xca, 0xc1, 0xc3, 0x4d, 0x0d, 0x5b, 0x9f, 0x68, 0x1e, 0x15, 0xe9,
	0x75, 0x87, 0x67, 0xc0, 0xd2, 0xfb, 0xf1, 0x20, 0xd7, 0x7f, 0xda, 0x68, 0xfe, 0xa5, 0xfa, 0x8d,
	0x87, 0x9a, 0xfd, 0xf8, 0x55, 0xa6, 0x3c, 0x7b, 0x16, 0x5d, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xe1, 0x47, 0xbf, 0xf2, 0x1b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventServerClient is the client API for EventServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventServerClient interface {
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
	Poll(ctx context.Context, in *PollEvent, opts ...grpc.CallOption) (*PollEvent, error)
}

type eventServerClient struct {
	cc *grpc.ClientConn
}

func NewEventServerClient(cc *grpc.ClientConn) EventServerClient {
	return &eventServerClient{cc}
}

func (c *eventServerClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := c.cc.Invoke(ctx, "/report.EventServer/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServerClient) Poll(ctx context.Context, in *PollEvent, opts ...grpc.CallOption) (*PollEvent, error) {
	out := new(PollEvent)
	err := c.cc.Invoke(ctx, "/report.EventServer/Poll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServerServer is the server API for EventServer service.
type EventServerServer interface {
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
	Poll(context.Context, *PollEvent) (*PollEvent, error)
}

// UnimplementedEventServerServer can be embedded to have forward compatible implementations.
type UnimplementedEventServerServer struct {
}

func (*UnimplementedEventServerServer) Report(ctx context.Context, req *ReportRequest) (*ReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}
func (*UnimplementedEventServerServer) Poll(ctx context.Context, req *PollEvent) (*PollEvent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Poll not implemented")
}

func RegisterEventServerServer(s *grpc.Server, srv EventServerServer) {
	s.RegisterService(&_EventServer_serviceDesc, srv)
}

func _EventServer_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServerServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/report.EventServer/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServerServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventServer_Poll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServerServer).Poll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/report.EventServer/Poll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServerServer).Poll(ctx, req.(*PollEvent))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "report.EventServer",
	HandlerType: (*EventServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _EventServer_Report_Handler,
		},
		{
			MethodName: "Poll",
			Handler:    _EventServer_Poll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/event_report.proto",
}

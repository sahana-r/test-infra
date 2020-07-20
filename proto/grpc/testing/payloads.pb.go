// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/testing/payloads.proto

package grpc_testing

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

type ByteBufferParams struct {
	ReqSize              int32    `protobuf:"varint,1,opt,name=req_size,json=reqSize,proto3" json:"req_size,omitempty"`
	RespSize             int32    `protobuf:"varint,2,opt,name=resp_size,json=respSize,proto3" json:"resp_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ByteBufferParams) Reset()         { *m = ByteBufferParams{} }
func (m *ByteBufferParams) String() string { return proto.CompactTextString(m) }
func (*ByteBufferParams) ProtoMessage()    {}
func (*ByteBufferParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e94ae443d7fd180, []int{0}
}

func (m *ByteBufferParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ByteBufferParams.Unmarshal(m, b)
}
func (m *ByteBufferParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ByteBufferParams.Marshal(b, m, deterministic)
}
func (m *ByteBufferParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ByteBufferParams.Merge(m, src)
}
func (m *ByteBufferParams) XXX_Size() int {
	return xxx_messageInfo_ByteBufferParams.Size(m)
}
func (m *ByteBufferParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ByteBufferParams.DiscardUnknown(m)
}

var xxx_messageInfo_ByteBufferParams proto.InternalMessageInfo

func (m *ByteBufferParams) GetReqSize() int32 {
	if m != nil {
		return m.ReqSize
	}
	return 0
}

func (m *ByteBufferParams) GetRespSize() int32 {
	if m != nil {
		return m.RespSize
	}
	return 0
}

type SimpleProtoParams struct {
	ReqSize              int32    `protobuf:"varint,1,opt,name=req_size,json=reqSize,proto3" json:"req_size,omitempty"`
	RespSize             int32    `protobuf:"varint,2,opt,name=resp_size,json=respSize,proto3" json:"resp_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleProtoParams) Reset()         { *m = SimpleProtoParams{} }
func (m *SimpleProtoParams) String() string { return proto.CompactTextString(m) }
func (*SimpleProtoParams) ProtoMessage()    {}
func (*SimpleProtoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e94ae443d7fd180, []int{1}
}

func (m *SimpleProtoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleProtoParams.Unmarshal(m, b)
}
func (m *SimpleProtoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleProtoParams.Marshal(b, m, deterministic)
}
func (m *SimpleProtoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleProtoParams.Merge(m, src)
}
func (m *SimpleProtoParams) XXX_Size() int {
	return xxx_messageInfo_SimpleProtoParams.Size(m)
}
func (m *SimpleProtoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleProtoParams.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleProtoParams proto.InternalMessageInfo

func (m *SimpleProtoParams) GetReqSize() int32 {
	if m != nil {
		return m.ReqSize
	}
	return 0
}

func (m *SimpleProtoParams) GetRespSize() int32 {
	if m != nil {
		return m.RespSize
	}
	return 0
}

type ComplexProtoParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComplexProtoParams) Reset()         { *m = ComplexProtoParams{} }
func (m *ComplexProtoParams) String() string { return proto.CompactTextString(m) }
func (*ComplexProtoParams) ProtoMessage()    {}
func (*ComplexProtoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e94ae443d7fd180, []int{2}
}

func (m *ComplexProtoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComplexProtoParams.Unmarshal(m, b)
}
func (m *ComplexProtoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComplexProtoParams.Marshal(b, m, deterministic)
}
func (m *ComplexProtoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComplexProtoParams.Merge(m, src)
}
func (m *ComplexProtoParams) XXX_Size() int {
	return xxx_messageInfo_ComplexProtoParams.Size(m)
}
func (m *ComplexProtoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ComplexProtoParams.DiscardUnknown(m)
}

var xxx_messageInfo_ComplexProtoParams proto.InternalMessageInfo

type PayloadConfig struct {
	// Types that are valid to be assigned to Payload:
	//	*PayloadConfig_BytebufParams
	//	*PayloadConfig_SimpleParams
	//	*PayloadConfig_ComplexParams
	Payload              isPayloadConfig_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PayloadConfig) Reset()         { *m = PayloadConfig{} }
func (m *PayloadConfig) String() string { return proto.CompactTextString(m) }
func (*PayloadConfig) ProtoMessage()    {}
func (*PayloadConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e94ae443d7fd180, []int{3}
}

func (m *PayloadConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayloadConfig.Unmarshal(m, b)
}
func (m *PayloadConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayloadConfig.Marshal(b, m, deterministic)
}
func (m *PayloadConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayloadConfig.Merge(m, src)
}
func (m *PayloadConfig) XXX_Size() int {
	return xxx_messageInfo_PayloadConfig.Size(m)
}
func (m *PayloadConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PayloadConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PayloadConfig proto.InternalMessageInfo

type isPayloadConfig_Payload interface {
	isPayloadConfig_Payload()
}

type PayloadConfig_BytebufParams struct {
	BytebufParams *ByteBufferParams `protobuf:"bytes,1,opt,name=bytebuf_params,json=bytebufParams,proto3,oneof"`
}

type PayloadConfig_SimpleParams struct {
	SimpleParams *SimpleProtoParams `protobuf:"bytes,2,opt,name=simple_params,json=simpleParams,proto3,oneof"`
}

type PayloadConfig_ComplexParams struct {
	ComplexParams *ComplexProtoParams `protobuf:"bytes,3,opt,name=complex_params,json=complexParams,proto3,oneof"`
}

func (*PayloadConfig_BytebufParams) isPayloadConfig_Payload() {}

func (*PayloadConfig_SimpleParams) isPayloadConfig_Payload() {}

func (*PayloadConfig_ComplexParams) isPayloadConfig_Payload() {}

func (m *PayloadConfig) GetPayload() isPayloadConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PayloadConfig) GetBytebufParams() *ByteBufferParams {
	if x, ok := m.GetPayload().(*PayloadConfig_BytebufParams); ok {
		return x.BytebufParams
	}
	return nil
}

func (m *PayloadConfig) GetSimpleParams() *SimpleProtoParams {
	if x, ok := m.GetPayload().(*PayloadConfig_SimpleParams); ok {
		return x.SimpleParams
	}
	return nil
}

func (m *PayloadConfig) GetComplexParams() *ComplexProtoParams {
	if x, ok := m.GetPayload().(*PayloadConfig_ComplexParams); ok {
		return x.ComplexParams
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PayloadConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PayloadConfig_BytebufParams)(nil),
		(*PayloadConfig_SimpleParams)(nil),
		(*PayloadConfig_ComplexParams)(nil),
	}
}

func init() {
	proto.RegisterType((*ByteBufferParams)(nil), "grpc.testing.ByteBufferParams")
	proto.RegisterType((*SimpleProtoParams)(nil), "grpc.testing.SimpleProtoParams")
	proto.RegisterType((*ComplexProtoParams)(nil), "grpc.testing.ComplexProtoParams")
	proto.RegisterType((*PayloadConfig)(nil), "grpc.testing.PayloadConfig")
}

func init() {
	proto.RegisterFile("grpc/testing/payloads.proto", fileDescriptor_7e94ae443d7fd180)
}

var fileDescriptor_7e94ae443d7fd180 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x4f, 0x4f, 0x02, 0x31,
	0x10, 0xc5, 0x59, 0x8c, 0x02, 0x23, 0x4b, 0xb4, 0xf1, 0xa0, 0x21, 0x51, 0xb2, 0x27, 0x4f, 0x25,
	0xd1, 0x6f, 0xb0, 0x24, 0x8a, 0x7a, 0xd9, 0x2c, 0x1f, 0x80, 0x74, 0x71, 0xba, 0x69, 0x02, 0xb4,
	0xb4, 0x25, 0x71, 0xf9, 0xe0, 0x9e, 0x4d, 0xff, 0x90, 0x80, 0x1c, 0xb9, 0xce, 0x9b, 0xf7, 0x9b,
	0xf7, 0x32, 0x30, 0xac, 0xb5, 0x5a, 0x8c, 0x2d, 0x1a, 0x2b, 0xd6, 0xf5, 0x58, 0xb1, 0x66, 0x29,
	0xd9, 0xb7, 0xa1, 0x4a, 0x4b, 0x2b, 0x49, 0xdf, 0x89, 0x34, 0x8a, 0xd9, 0x27, 0xdc, 0xe4, 0x8d,
	0xc5, 0x7c, 0xcb, 0x39, 0xea, 0x82, 0x69, 0xb6, 0x32, 0xe4, 0x01, 0xba, 0x1a, 0x37, 0x73, 0x23,
	0x76, 0x78, 0x9f, 0x8c, 0x92, 0xe7, 0xcb, 0xb2, 0xa3, 0x71, 0x33, 0x13, 0x3b, 0x24, 0x43, 0xe8,
	0x69, 0x34, 0x2a, 0x68, 0x6d, 0xaf, 0x75, 0xdd, 0xc0, 0x89, 0xd9, 0x17, 0xdc, 0xce, 0xc4, 0x4a,
	0x2d, 0xb1, 0x70, 0x87, 0xce, 0x84, 0xdd, 0x01, 0x99, 0x48, 0x07, 0xfb, 0x39, 0xa0, 0x65, 0xbf,
	0x09, 0xa4, 0x45, 0xe8, 0x33, 0x91, 0x6b, 0x2e, 0x6a, 0xf2, 0x0e, 0x83, 0xaa, 0xb1, 0x58, 0x6d,
	0xf9, 0x5c, 0xf9, 0x1d, 0x7f, 0xe5, 0xfa, 0xe5, 0x91, 0x1e, 0xf6, 0xa4, 0xff, 0x4b, 0x4e, 0x5b,
	0x65, 0x1a, 0x7d, 0x31, 0xe8, 0x1b, 0xa4, 0xc6, 0xa7, 0xdf, 0x73, 0xda, 0x9e, 0xf3, 0x74, 0xcc,
	0x39, 0x29, 0x38, 0x6d, 0x95, 0xfd, 0xe0, 0x8b, 0x9c, 0x0f, 0x18, 0x2c, 0x42, 0xf0, 0x3d, 0xe8,
	0xc2, 0x83, 0x46, 0xc7, 0xa0, 0xd3, 0x72, 0x2e, 0x52, 0x74, 0x86, 0x41, 0xde, 0x83, 0x4e, 0x7c,
	0x5e, 0x75, 0xe5, 0x9f, 0xf7, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x46, 0x17, 0x84, 0xdb,
	0x01, 0x00, 0x00,
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sensor.proto

/*
Package github_com_argoproj_argo_events_proto is a generated protocol buffer package.

It is generated from these files:
	sensor.proto

It has these top-level messages:
	SensorEvent
	SensorResponse
*/
package github_com_argoproj_argo_events_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SensorEvent struct {
	// name of the sensor
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// type is event/signal type
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *SensorEvent) Reset()                    { *m = SensorEvent{} }
func (m *SensorEvent) String() string            { return proto.CompactTextString(m) }
func (*SensorEvent) ProtoMessage()               {}
func (*SensorEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SensorEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SensorEvent) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type SensorResponse struct {
	// action sensor should take.
	Action string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
}

func (m *SensorResponse) Reset()                    { *m = SensorResponse{} }
func (m *SensorResponse) String() string            { return proto.CompactTextString(m) }
func (*SensorResponse) ProtoMessage()               {}
func (*SensorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SensorResponse) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func init() {
	proto.RegisterType((*SensorEvent)(nil), "github.com.argoproj.argo_events.proto.SensorEvent")
	proto.RegisterType((*SensorResponse)(nil), "github.com.argoproj.argo_events.proto.SensorResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SensorUpdate service

type SensorUpdateClient interface {
	// updates the state of sensor
	UpdateSensor(ctx context.Context, opts ...grpc.CallOption) (SensorUpdate_UpdateSensorClient, error)
}

type sensorUpdateClient struct {
	cc *grpc.ClientConn
}

func NewSensorUpdateClient(cc *grpc.ClientConn) SensorUpdateClient {
	return &sensorUpdateClient{cc}
}

func (c *sensorUpdateClient) UpdateSensor(ctx context.Context, opts ...grpc.CallOption) (SensorUpdate_UpdateSensorClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SensorUpdate_serviceDesc.Streams[0], c.cc, "/github.com.argoproj.argo_events.proto.SensorUpdate/UpdateSensor", opts...)
	if err != nil {
		return nil, err
	}
	x := &sensorUpdateUpdateSensorClient{stream}
	return x, nil
}

type SensorUpdate_UpdateSensorClient interface {
	Send(*SensorEvent) error
	CloseAndRecv() (*SensorResponse, error)
	grpc.ClientStream
}

type sensorUpdateUpdateSensorClient struct {
	grpc.ClientStream
}

func (x *sensorUpdateUpdateSensorClient) Send(m *SensorEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sensorUpdateUpdateSensorClient) CloseAndRecv() (*SensorResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SensorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SensorUpdate service

type SensorUpdateServer interface {
	// updates the state of sensor
	UpdateSensor(SensorUpdate_UpdateSensorServer) error
}

func RegisterSensorUpdateServer(s *grpc.Server, srv SensorUpdateServer) {
	s.RegisterService(&_SensorUpdate_serviceDesc, srv)
}

func _SensorUpdate_UpdateSensor_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SensorUpdateServer).UpdateSensor(&sensorUpdateUpdateSensorServer{stream})
}

type SensorUpdate_UpdateSensorServer interface {
	SendAndClose(*SensorResponse) error
	Recv() (*SensorEvent, error)
	grpc.ServerStream
}

type sensorUpdateUpdateSensorServer struct {
	grpc.ServerStream
}

func (x *sensorUpdateUpdateSensorServer) SendAndClose(m *SensorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sensorUpdateUpdateSensorServer) Recv() (*SensorEvent, error) {
	m := new(SensorEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SensorUpdate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.argoproj.argo_events.proto.SensorUpdate",
	HandlerType: (*SensorUpdateServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateSensor",
			Handler:       _SensorUpdate_UpdateSensor_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "sensor.proto",
}

func init() { proto.RegisterFile("sensor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0xcd, 0x2b,
	0xce, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4b, 0x2c, 0x4a, 0xcf, 0x2f, 0x28, 0xca, 0xcf, 0x02, 0x33, 0xe2,
	0x53, 0xcb, 0x52, 0xf3, 0x4a, 0x8a, 0x21, 0xca, 0x94, 0x4c, 0xb9, 0xb8, 0x83, 0xc1, 0xda, 0x5c,
	0x41, 0xa2, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x60, 0x36, 0x48, 0xac, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x09, 0x22, 0x06, 0x62, 0x2b, 0x69,
	0x70, 0xf1, 0x41, 0xb4, 0x05, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x89, 0x71, 0xb1,
	0x25, 0x26, 0x97, 0x64, 0xe6, 0xe7, 0x41, 0xf5, 0x42, 0x79, 0x46, 0xbd, 0x8c, 0x5c, 0x3c, 0x10,
	0xa5, 0xa1, 0x05, 0x29, 0x89, 0x25, 0xa9, 0x42, 0xb5, 0x5c, 0x3c, 0x10, 0x16, 0x44, 0x54, 0xc8,
	0x48, 0x8f, 0x28, 0x97, 0xea, 0x21, 0x39, 0x53, 0xca, 0x94, 0x24, 0x3d, 0x30, 0x37, 0x2a, 0x31,
	0x68, 0x30, 0x26, 0xb1, 0x81, 0x65, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xed, 0xe1, 0xc6,
	0x1b, 0x2e, 0x01, 0x00, 0x00,
}
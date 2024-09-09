// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: model.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Classifier_Classify_FullMethodName = "/model.Classifier/Classify"
)

// ClassifierClient is the client API for Classifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClassifierClient interface {
	Classify(ctx context.Context, in *ClassifyRequest, opts ...grpc.CallOption) (*ClassifyResponse, error)
}

type classifierClient struct {
	cc grpc.ClientConnInterface
}

func NewClassifierClient(cc grpc.ClientConnInterface) ClassifierClient {
	return &classifierClient{cc}
}

func (c *classifierClient) Classify(ctx context.Context, in *ClassifyRequest, opts ...grpc.CallOption) (*ClassifyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClassifyResponse)
	err := c.cc.Invoke(ctx, Classifier_Classify_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClassifierServer is the server API for Classifier service.
// All implementations must embed UnimplementedClassifierServer
// for forward compatibility.
type ClassifierServer interface {
	Classify(context.Context, *ClassifyRequest) (*ClassifyResponse, error)
	mustEmbedUnimplementedClassifierServer()
}

// UnimplementedClassifierServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClassifierServer struct{}

func (UnimplementedClassifierServer) Classify(context.Context, *ClassifyRequest) (*ClassifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Classify not implemented")
}
func (UnimplementedClassifierServer) mustEmbedUnimplementedClassifierServer() {}
func (UnimplementedClassifierServer) testEmbeddedByValue()                    {}

// UnsafeClassifierServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClassifierServer will
// result in compilation errors.
type UnsafeClassifierServer interface {
	mustEmbedUnimplementedClassifierServer()
}

func RegisterClassifierServer(s grpc.ServiceRegistrar, srv ClassifierServer) {
	// If the following call pancis, it indicates UnimplementedClassifierServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Classifier_ServiceDesc, srv)
}

func _Classifier_Classify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClassifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClassifierServer).Classify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Classifier_Classify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClassifierServer).Classify(ctx, req.(*ClassifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Classifier_ServiceDesc is the grpc.ServiceDesc for Classifier service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Classifier_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.Classifier",
	HandlerType: (*ClassifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Classify",
			Handler:    _Classifier_Classify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model.proto",
}

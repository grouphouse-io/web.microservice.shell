// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sg

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SGClient is the client API for SG service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SGClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type sGClient struct {
	cc grpc.ClientConnInterface
}

func NewSGClient(cc grpc.ClientConnInterface) SGClient {
	return &sGClient{cc}
}

var sGSayHelloStreamDesc = &grpc.StreamDesc{
	StreamName: "SayHello",
}

func (c *sGClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/sg.SG/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SGService is the service API for SG service.
// Fields should be assigned to their respective handler implementations only before
// RegisterSGService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type SGService struct {
	// Sends a greeting
	SayHello func(context.Context, *HelloRequest) (*HelloReply, error)
}

func (s *SGService) sayHello(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/sg.SG/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterSGService registers a service implementation with a gRPC server.
func RegisterSGService(s grpc.ServiceRegistrar, srv *SGService) {
	srvCopy := *srv
	if srvCopy.SayHello == nil {
		srvCopy.SayHello = func(context.Context, *HelloRequest) (*HelloReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "sg.SG",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "SayHello",
				Handler:    srvCopy.sayHello,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "sg.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewSGService creates a new SGService containing the
// implemented methods of the SG service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewSGService(s interface{}) *SGService {
	ns := &SGService{}
	if h, ok := s.(interface {
		SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	}); ok {
		ns.SayHello = h.SayHello
	}
	return ns
}

// UnstableSGService is the service API for SG service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableSGService interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

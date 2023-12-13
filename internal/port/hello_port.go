package port

import (
	"context"
	"github.com/lukmansjy/my-grpc-proto/protogen/go/hello"
	"google.golang.org/grpc"
)

type HelloClientPort interface {
	SayHello(ctx context.Context, in *hello.HelloRequest, opts ...grpc.CallOption) (*hello.HelloResponse, error)
	SayManyHellos(ctx context.Context, in *hello.HelloRequest, opts ...grpc.CallOption) (hello.HelloService_SayManyHellosClient, error)
}

package hello

import (
	"context"
	"github.com/lukmansjy/my-grpc-go-client/internal/port"
	"github.com/lukmansjy/my-grpc-proto/protogen/go/hello"
	"google.golang.org/grpc"
	"log"
)

type HelloAdapter struct {
	helloClient port.HelloClientPort
}

func NewHelloAdapter(conn *grpc.ClientConn) (*HelloAdapter, error) {
	client := hello.NewHelloServiceClient(conn)

	return &HelloAdapter{
		helloClient: client,
	}, nil
}

func (a *HelloAdapter) SayHello(ctx context.Context, name string) (*hello.HelloResponse, error) {
	helloRequest := &hello.HelloRequest{
		Name: name,
	}

	greet, err := a.helloClient.SayHello(ctx, helloRequest)
	if err != nil {
		log.Fatalf("Error on SayHello: %v\n", err)
	}

	return greet, nil
}

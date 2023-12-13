package hello

import (
	"context"
	"github.com/lukmansjy/my-grpc-go-client/internal/port"
	"github.com/lukmansjy/my-grpc-proto/protogen/go/hello"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
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

func (a *HelloAdapter) SayManyHellos(ctx context.Context, name string) {
	helloRequest := &hello.HelloRequest{
		Name: name,
	}

	greetStream, err := a.helloClient.SayManyHellos(ctx, helloRequest)
	if err != nil {
		log.Fatalf("Error on SayManyHellos: %v\n", err)
	}

	for {
		greet, err := greetStream.Recv()
		if err != nil {
			break
		}
		if err == io.EOF {
			break
		}

		log.Println(greet.Greet)
	}
}

func (a *HelloAdapter) SayHelloToEveryone(ctx context.Context, names []string) {
	greetStream, err := a.helloClient.SayHelloToEveryone(ctx)

	if err != nil {
		log.Fatalf("Error on SayHelloToEveryone: %v\n", err)
	}

	for _, name := range names {
		log.Printf("stream sending %s to server...\n", name)
		req := &hello.HelloRequest{Name: name}

		err := greetStream.Send(req)
		if err != nil {
			log.Fatalf("Error on send SayHelloToEveryone: %v\n", err)
		}
		time.Sleep(1 * time.Second)
	}

	log.Println("stream finish")

	res, err := greetStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error on close and recv SayHelloToEveryone: %v\n", err)
	}

	log.Println(res.Greet)
}

package main

import (
	"context"
	"github.com/lukmansjy/my-grpc-go-client/internal/adapter/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(":9090", opts...)
	if err != nil {
		log.Fatalf("Can't not connect to gRPC server %v\n", err)
	}

	defer conn.Close()

	hellAdapter, err := hello.NewHelloAdapter(conn)

	if err != nil {
		log.Fatalf("Can't not create HelloAdapter %v\n", err)
	}

	//runSayHello(hellAdapter, "Lukman Sanjaya")
	//runSayManyHello(hellAdapter, "Lukman Sanjaya")
	//runSayHelloToEveryone(hellAdapter, []string{"lukman sanjaya", "joko", "roni", "podo moro"})
	runSayHelloContinuous(hellAdapter, []string{"lukman sanjaya", "joko", "roni", "podo moro"})
}

func runSayHello(adapter *hello.HelloAdapter, name string) {
	greet, err := adapter.SayHello(context.Background(), name)

	if err != nil {
		log.Fatalf("Can't not call SayHello %v\n", err)
	}

	log.Println((greet.Greet))
}

func runSayManyHello(adapter *hello.HelloAdapter, name string) {
	adapter.SayManyHellos(context.Background(), name)
}

func runSayHelloToEveryone(adapter *hello.HelloAdapter, names []string) {
	adapter.SayHelloToEveryone(context.Background(), names)
}

func runSayHelloContinuous(adapter *hello.HelloAdapter, names []string) {
	adapter.SayHelloContinuous(context.Background(), names)
}

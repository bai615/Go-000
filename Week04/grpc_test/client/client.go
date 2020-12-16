package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"

	"github.com/bai615/Go-000/Week04/grpc_test/proto"
)

func main() {
	//stream
	conn, err := grpc.Dial("127.0.0.1:8088", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "world"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}

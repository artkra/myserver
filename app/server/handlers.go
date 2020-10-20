package server

import (
	"context"
	"fmt"
	pb "myserver/app/pb"
)

// SayHello says hello
func (g *Greeter) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello, %s!", r.Name),
	}, nil
}

// SayHello says hello
func (g *Greeter) SayHelloAgain(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello, again %s!", r.Name),
	}, nil
}

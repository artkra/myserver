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

// TODO: 6. Implement handlers for CRUD operations over LineItems and creatives.
// TODO: 7. In every CRUD operation update in-memory config
// update data in mongo;
// if success ->
// push diff data to special channel 'chan<- models.Diff' (also a field of server)

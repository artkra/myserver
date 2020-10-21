package server

import (
	"context"
	"log"

	pb "myserver/app/pb"
	"net"
	"net/http"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// TODO: 3. Enrich struct with mutex, models.Config instance, db connection, update channel
type Greeter struct {
	wg sync.WaitGroup
	pb.UnimplementedGreeterServer
}

// New creates new server greeter
func NewServer() *Greeter {
	return &Greeter{}
}

// TODO: 777. Ensure gracefull stop by signals.
// TODO: 888. Add logging.
// Start starts server
func (g *Greeter) Start() {
	g.wg.Add(1)
	go func() {
		log.Fatal(g.startGRPC())
		g.wg.Done()
	}()
	g.wg.Add(1)
	go func() {
		log.Fatal(g.startREST())
		g.wg.Done()
	}()
	g.wg.Wait()
}

func (g *Greeter) startGRPC() error {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, g)
	grpcServer.Serve(lis)
	return nil
}

func (g *Greeter) startREST() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	mux.HandlePath(
		"GET", "/swagger.json", runtime.HandlerFunc(SwaggerJSONHandler),
	)
	mux.HandlePath(
		"GET", "/swagger", runtime.HandlerFunc(SwaggerUIHandler),
	)
	// TODO: 999. Figure out how to serve dir in gateway
	mux.HandlePath(
		"GET", "/swagger/swagger-ui-bundle.js", runtime.HandlerFunc(SwaggerUIHandler),
	)
	mux.HandlePath(
		"GET", "/swagger/swagger-ui-standalone-preset.js", runtime.HandlerFunc(SwaggerUIHandler),
	)
	mux.HandlePath(
		"GET", "/swagger/swagger-ui.css", runtime.HandlerFunc(SwaggerUIHandler),
	)
	mux.HandlePath(
		"GET", "/swagger/favicon-16x16.png", runtime.HandlerFunc(SwaggerUIHandler),
	)
	mux.HandlePath(
		"GET", "/swagger/favicon-32x32.png", runtime.HandlerFunc(SwaggerUIHandler),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, ":8080", opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8090", mux)
}

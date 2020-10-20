package interceptors

import (
	"context"

	"google.golang.org/grpc"
)

func EnrichCtxInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	ctx = context.WithValue(ctx, "some", 2)
	return handler(ctx, req)
}

func WithServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(EnrichCtxInterceptor)
}

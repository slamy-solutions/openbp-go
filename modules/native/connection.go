package native

import (
	"time"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func makeGrpcDial(address string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(
		[]grpc.DialOption{
			grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
			grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
			grpc.WithBlock(),
			grpc.WithTimeout(time.Second * 5),
		},
		opts...,
	)

	return grpc.Dial(address, opts...)
}

func makeGrpcClient[T interface{}](clientFunction func(grpc.ClientConnInterface) T, address string, opts ...grpc.DialOption) (*grpc.ClientConn, T, error) {
	dial, err := makeGrpcDial(address, opts...)
	if err != nil {
		var result T
		return nil, result, err
	}

	client := clientFunction(dial)
	return dial, client, nil
}

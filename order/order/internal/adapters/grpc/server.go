package grpc

import (
	"fmt"
	"github.com/w3gop2p/microservices-proto/golang/order"
	"github.com/w3gop2p/microservices/order/order/config"
	"github.com/w3gop2p/microservices/order/order/internal/ports"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	// order "github.com/w3gop2p/microservices/order/order/internal/adapters/grpc/proto"
)

type Adapter struct {
	api                            ports.APIPort
	port                           int
	server                         *grpc.Server
	order.UnimplementedOrderServer // Forward compatibility support
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	)
	a.server = grpcServer
	order.RegisterOrderServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port ")
	}
}

func (a Adapter) Stop() {
	a.server.Stop()
}

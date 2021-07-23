package server

import (
	v1 "lbc/api/v1"
	"lbc/pkg/queue"
	"lbc/rpc"

	"github.com/go-logr/logr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RegisterRPCs registers all RPCs against the given grpc.Server
func RegisterRPCs(log logr.Logger, q *queue.FIFO, grpcServer *grpc.Server) {
	be := &rpc.BackendService{
		Log:   log,
		Queue: q,
	}
	v1.RegisterBackendServer(grpcServer, be)
	reflection.Register(grpcServer)
}

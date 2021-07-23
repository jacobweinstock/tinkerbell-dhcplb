package rpc

import (
	"context"
	v1 "lbc/api/v1"
	"lbc/pkg/queue"

	"github.com/go-logr/logr"
)

// BackendService
type BackendService struct {
	Log   logr.Logger
	Queue *queue.FIFO
	v1.UnimplementedBackendServer
}

// Set a DHCP backend for a MAC address
func (b *BackendService) Set(ctx context.Context, in *v1.BackendRequest) (*v1.BackendResponse, error) {
	b.Queue.Enqueue(queue.Data{Backend: in.Backend, MAC: in.MacAddress})
	b.Log.V(0).Info("debugging", "queue", b.Queue.Queue())
	return &v1.BackendResponse{Success: true}, nil
}

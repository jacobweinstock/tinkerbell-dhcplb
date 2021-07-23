package cmd

import (
	"context"
	"fmt"
	"lbc/pkg/queue"
	"lbc/server"
	"net"
	"sync"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

func Execute(ctx context.Context) error {
	port := "60061"

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	log := defaultLogger("info")
	q := &queue.FIFO{}

	// start the controller
	go startController(ctx, log, q)

	server.RegisterRPCs(log, q, grpcServer)

	go func() {
		<-ctx.Done()
		//log.V(0).Info("ctx cancelled, shutting down LBC")
		grpcServer.GracefulStop()
	}()

	log.V(0).Info("Starting LBC server")
	return grpcServer.Serve(listen)
}

// defaultLogger is zap logr implementation
func defaultLogger(level string) logr.Logger {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	switch level {
	case "debug":
		config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
	zapLogger, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("who watches the watchmen (%v)?", err))
	}

	return zapr.NewLogger(zapLogger)
}

func startController(ctx context.Context, log logr.Logger, q *queue.FIFO) {
	var controllerWg sync.WaitGroup
	controllerWg.Add(1)
	go controller(ctx, log, controllerWg, q)
	// graceful shutdown when a signal is caught
	<-ctx.Done()
	controllerWg.Wait()
}

func controller(ctx context.Context, log logr.Logger, stopControllerWg sync.WaitGroup, q *queue.FIFO) {
	for {
		select {
		case <-ctx.Done():
			log.V(0).Info("stopping controller")
			stopControllerWg.Done()
			return
		default:
			d := q.Dequeue()
			if d.Backend == "" {
				//log.V(0).Info("nothing to do")
				//time.Sleep(time.Second)
				continue
			}
			log.V(0).Info("update the dhcplb file", "data", d)
		}
	}
}

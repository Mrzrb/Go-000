package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"main/api"
	"main/internal"

	"github.com/jinzhu/gorm"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return serve(ctx)
	})

	g.Go(func() error {
		return signalHandle(ctx)
	})

	if err := g.Wait(); err != nil {
		log.Printf("[wait err] : %s", err)
	}
}

func signalHandle(ctx context.Context) error {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	for {
		select {
		case sig := <-sigs:
			return fmt.Errorf("rec sig : %+v", sig)
		case <-ctx.Done():
			return fmt.Errorf("ctx done")
		}
	}
}

func serve(ctx context.Context) error {
	log.Println("[grpc server start]")
	server := grpc.NewServer()

	go func(ctx context.Context) {
		<-ctx.Done()
		server.Stop()
	}(ctx)

	//srv :=
	//demoSrv := &internal.DemoService{}
	demoSrv := internal.InitDemoService(&gorm.DB{})
	api.RegisterDemoServiceServer(server, demoSrv)

	conn, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		return fmt.Errorf("error occur : %w", err)
	}
	defer conn.Close()

	return server.Serve(conn)
}

package main

import (
	"context"
	"log"
	"main/api"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := api.NewDemoServiceClient(conn)
	req := &api.GetDemoReq{
		Id: "123",
	}
	rsp, err := c.Demo(ctx, req)
	if err != nil {
		return
	}
	log.Println(rsp)
}

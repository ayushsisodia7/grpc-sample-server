package main

import (
	"context"
	"fmt"
	"log"
	"mymodule/invoicer"
	"net"

	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	
	serverRegistrar := grpc.NewServer()
	
	service := &myInvoicerServer{}
	
	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	
	err = serverRegistrar.Serve(lis)
	
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
	
}

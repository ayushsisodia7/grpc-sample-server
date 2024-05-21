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
	fmt.Println("Here1")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	fmt.Println("Here2")
	serverRegistrar := grpc.NewServer()
	fmt.Println("Here3")
	service := &myInvoicerServer{}
	fmt.Println("Here4")
	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	fmt.Println("Here5")
	err = serverRegistrar.Serve(lis)
	fmt.Println("Here6")
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
	fmt.Println("Here")
}

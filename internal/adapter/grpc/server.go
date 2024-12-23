package grpc

import (
	"fmt"
	"github.com/mahfuzon/grpc-server/internal/port"
	"log"
	"net"
)
import "github.com/mahfuzon/grpc-proto/protogen/go/bank"
import "google.golang.org/grpc"

/**
* Author Company: PT. ASLI ISOAE SOLUSINE
* Author Url: www.isoae.id
* Author Name : Mahfuzon Akhiar
* Author Email : mahfuzon0@gmail.com
* Date : 23/12/24
* Time : 12.36
* Project Name : grpc-server
* Licensed to :
**/

type GrpcAdapter struct {
	bankService port.BankService
	grpcPort    int
	server      *grpc.Server
	bank.BankServiceServer
}

func NewGrpcAdapter(bankService port.BankService, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{bankService: bankService, grpcPort: grpcPort}
}

func (a GrpcAdapter) Run() {
	var err error

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))

	if err != nil {
		log.Fatalf("failed to listen port: %v, %v", a.grpcPort, err)
	}

	log.Printf("grpc server listening on port %d", a.grpcPort)

	grpcServer := grpc.NewServer()
	a.server = grpcServer
	bank.RegisterBankServiceServer(grpcServer, a)

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v on port : %v", err, a.grpcPort)
	}
}

func (a GrpcAdapter) Stop() {
	a.server.Stop()
}

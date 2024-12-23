package main

import (
	"github.com/mahfuzon/grpc-server/internal/adapter/grpc"
	"github.com/mahfuzon/grpc-server/internal/aplication"
)

func main() {
	bankService := &aplication.BankService{}

	grpcAdapter := grpc.NewGrpcAdapter(bankService, 9090)

	grpcAdapter.Run()
}

package main

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/mahfuzon/grpc-server/db"
	"github.com/mahfuzon/grpc-server/internal/adapter/grpc"
	"github.com/mahfuzon/grpc-server/internal/aplication"
	"log"
)

func main() {

	sqlDb, err := sql.Open("pgx", "postgres://postgres:123456@localhost:5432/bank_db?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	db.Migrate(sqlDb)

	bankService := &aplication.BankService{}

	grpcAdapter := grpc.NewGrpcAdapter(bankService, 9090)

	grpcAdapter.Run()
}

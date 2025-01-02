package main

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	migration "github.com/mahfuzon/grpc-server/db"
	"github.com/mahfuzon/grpc-server/internal/adapter/db"
	"github.com/mahfuzon/grpc-server/internal/adapter/grpc"
	"github.com/mahfuzon/grpc-server/internal/aplication"
	"log"
)

func main() {
	sqlDb, err := sql.Open("pgx", "postgres://postgres:123456@localhost:5432/bank_db?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}

	migration.Migrate(sqlDb)

	dbAdapter, err := db.NewDatabaseAdapter(sqlDb)
	if err != nil {
		log.Fatalf("failed to create db adapter: %v", err)
	}

	bankService := aplication.NewBankService(dbAdapter)

	grpcAdapter := grpc.NewGrpcAdapter(bankService, 9090)

	grpcAdapter.Run()
}

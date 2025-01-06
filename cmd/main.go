package main

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	migration "github.com/mahfuzon/grpc-server/db"
	"github.com/mahfuzon/grpc-server/internal/adapter/db"
	"github.com/mahfuzon/grpc-server/internal/adapter/grpc"
	"github.com/mahfuzon/grpc-server/internal/aplication"
	"github.com/mahfuzon/grpc-server/internal/aplication/domain/bank"
	"log"
	"math/rand"
	"time"
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
	go generateExchangeRate(bankService, "USD", "IDR", time.Second*5)

	grpcAdapter := grpc.NewGrpcAdapter(bankService, 9090)

	grpcAdapter.Run()
}

func generateExchangeRate(bs *aplication.BankService, fromCurrency, toCurrency string, duration time.Duration) {
	ticker := time.NewTicker(duration)
	for range ticker.C {
		now := time.Now()
		validFrom := now.Truncate(time.Second).Add(3 * time.Second)
		validTo := validFrom.Add(duration).Add(-1 * time.Millisecond)

		dummyRate := bank.ExchangeRate{
			FromCurrency:       fromCurrency,
			ToCurrency:         toCurrency,
			Rate:               200 + float64(rand.Intn(100)),
			ValidFromTimeStamp: validFrom,
			ValidToTimeStamp:   validTo,
		}

		_, err := bs.CreateExchangeRate(dummyRate)
		if err != nil {
			log.Printf("failed to create exchange rate: %v", err)
		}
	}
}

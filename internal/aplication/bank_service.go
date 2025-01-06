package aplication

import (
	"github.com/google/uuid"
	"github.com/mahfuzon/grpc-server/internal/adapter/db"
	bankDomain "github.com/mahfuzon/grpc-server/internal/aplication/domain/bank"
	"github.com/mahfuzon/grpc-server/internal/port"
	"log"
	"time"
)

/**
* Author Company: PT. ASLI ISOAE SOLUSINE
* Author Url: www.isoae.id
* Author Name : Mahfuzon Akhiar
* Author Email : mahfuzon0@gmail.com
* Date : 23/12/24
* Time : 12.31
* Project Name : grpc-server
* Licensed to :
**/

type BankService struct {
	db port.BankDbPort
}

func NewBankService(db port.BankDbPort) *BankService {
	return &BankService{db: db}
}

func (s *BankService) FindCurrentBalance(acc string) float64 {
	bankAccount, err := s.db.GetBankAccountByAccountNumber(acc)
	if err != nil {
		log.Println("error finding current balance:", err)
	}
	return bankAccount.CurrentBalance
}

func (s *BankService) CreateExchangeRate(rate bankDomain.ExchangeRate) (uuid.UUID, error) {
	exchangeRateOrm := db.BankExchangeRateOrm{
		ExchangeRateUuid:   uuid.New(),
		FromCurrency:       rate.FromCurrency,
		ToCurrency:         rate.ToCurrency,
		Rate:               rate.Rate,
		ValidFromTimeStamp: rate.ValidFromTimeStamp,
		ValidToTimeStamp:   rate.ValidToTimeStamp,
	}

	return s.db.CreateExchangeRate(exchangeRateOrm)
}

func (s *BankService) FindExchangeRate(fromCurrency, toCurrency string, ts time.Time) float64 {
	exchangeRateOrm, err := s.db.GetExchangeRateAtTimeStamp(fromCurrency, toCurrency, ts)
	if err != nil {
		log.Printf("error finding exchange rate for %s to %s at %s", fromCurrency, toCurrency, ts)
	}

	return exchangeRateOrm.Rate
}

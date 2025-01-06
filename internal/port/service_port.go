package port

import (
	"github.com/google/uuid"
	bankDomain "github.com/mahfuzon/grpc-server/internal/aplication/domain/bank"
	"time"
)

/**
* Author Company: PT. ASLI ISOAE SOLUSINE
* Author Url: www.isoae.id
* Author Name : Mahfuzon Akhiar
* Author Email : mahfuzon0@gmail.com
* Date : 23/12/24
* Time : 12.33
* Project Name : grpc-server
* Licensed to :
**/

type BankService interface {
	FindCurrentBalance(string) float64
	CreateExchangeRate(rate bankDomain.ExchangeRate) (uuid.UUID, error)
	FindExchangeRate(fromCurrency, toCurrency string, ts time.Time) float64
}

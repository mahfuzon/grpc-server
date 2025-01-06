package port

import (
	"github.com/google/uuid"
	"github.com/mahfuzon/grpc-server/internal/adapter/db"
	"time"
)

/**
* Author Company: PT. ASLI ISOAE SOLUSINE
* Author Url: www.isoae.id
* Author Name : Mahfuzon Akhiar
* Author Email : mahfuzon0@gmail.com
* Date : 26/12/24
* Time : 21.48
* Project Name : grpc-server
* Licensed to :
**/

type BankDbPort interface {
	GetBankAccountByAccountNumber(accountNumber string) (db.BankAccountOrm, error)
	CreateExchangeRate(orm db.BankExchangeRateOrm) (uuid.UUID, error)
	GetExchangeRateAtTimeStamp(fromCurrency, toCurrency string, ts time.Time) (db.BankExchangeRateOrm, error)
}

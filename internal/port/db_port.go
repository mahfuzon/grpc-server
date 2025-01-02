package port

import (
	"github.com/mahfuzon/grpc-server/internal/adapter/db"
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
}

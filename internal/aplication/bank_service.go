package aplication

import (
	"github.com/mahfuzon/grpc-server/internal/port"
	"log"
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

package db

import (
	"github.com/google/uuid"
	"time"
)

/**
* Author Company: PT. ASLI ISOAE SOLUSINE
* Author Url: www.isoae.id
* Author Name : Mahfuzon Akhiar
* Author Email : mahfuzon0@gmail.com
* Date : 26/12/24
* Time : 18.58
* Project Name : grpc-server
* Licensed to :
**/

type BankAccountOrm struct {
	AccountUuid    uuid.UUID            `gorm:"primaryKey"`
	AccountNumber  string               `gorm:"column:account_a"`
	AccountName    string               `gorm:"column:account_name"`
	Currency       string               `gorm:"column:currency"`
	CurrentBalance float64              `gorm:"column:current_balance"`
	CreatedAt      time.Time            `gorm:"column:created_at"`
	UpdatedAt      time.Time            `gorm:"column:updated_at"`
	Transactions   []BankTransactionOrm `gorm:"foreignKey:AccountUuid"`
}

func (BankAccountOrm) TableName() string {
	return "bank_accounts"
}

type BankTransactionOrm struct {
	TransactionUuid      uuid.UUID `gorm:"primaryKey"`
	AccountUuid          uuid.UUID `gorm:"column:account_uuid"`
	TransactionTimestamp time.Time `gorm:"column:transaction_timestamp"`
	Amount               float64   `gorm:"column:amount"`
	TransactionType      string    `gorm:"column:transaction_type"`
	Notes                string    `gorm:"column:notes"`
	CreatedAt            time.Time `gorm:"column:created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at"`
}

func (BankTransactionOrm) TableName() string {
	return "bank_transactions"
}

type BankExchangeRateOrm struct {
	ExchangeRateUuid   uuid.UUID `gorm:"primaryKey"`
	FromCurrency       string
	ToCurrency         string
	Rate               float64
	ValidFromTimeStamp time.Time `gorm:"column:valid_from_timestamp"`
	ValidToTimeStamp   time.Time `gorm:"column:valid_to_timestamp"`
	CreatedAt          time.Time `gorm:"column:created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at"`
}

func (BankExchangeRateOrm) TableName() string {
	return "bank_exchange_rates"
}

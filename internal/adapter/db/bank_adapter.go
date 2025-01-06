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
* Time : 21.51
* Project Name : grpc-server
* Licensed to :
**/

func (a *dataBaseAdapter) GetBankAccountByAccountNumber(accountNumber string) (BankAccountOrm, error) {
	bankAccountOrm := BankAccountOrm{}
	err := a.db.Where("account_number = ?", accountNumber).First(&bankAccountOrm).Error
	return bankAccountOrm, err
}

func (a *dataBaseAdapter) CreateExchangeRate(orm BankExchangeRateOrm) (uuid.UUID, error) {
	err := a.db.Create(&orm).Error

	if err != nil {
		return uuid.Nil, err
	}

	return orm.ExchangeRateUuid, nil
}

func (a *dataBaseAdapter) GetExchangeRateAtTimeStamp(fromCurrency, toCurrency string, ts time.Time) (BankExchangeRateOrm, error) {
	BankExchangeRateOrm := BankExchangeRateOrm{}

	err := a.db.
		Where("from_currency = ?", fromCurrency).
		Where("to_currency = ?", toCurrency).
		Where("? BETWEEN valid_from_timestamp AND valid_to_timestamp", ts).
		First(&BankExchangeRateOrm).
		Error

	if err != nil {
		return BankExchangeRateOrm, err
	}

	return BankExchangeRateOrm, nil
}

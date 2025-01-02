package db

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

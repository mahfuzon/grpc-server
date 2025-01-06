package bank

import "time"

/**
* Author Company: PT. ASLI ISOAE SOLUSINE
* Author Url: www.isoae.id
* Author Name : Mahfuzon Akhiar
* Author Email : mahfuzon0@gmail.com
* Date : 05/01/25
* Time : 14.11
* Project Name : grpc-server
* Licensed to :
**/

type ExchangeRate struct {
	FromCurrency       string
	ToCurrency         string
	Rate               float64
	ValidFromTimeStamp time.Time
	ValidToTimeStamp   time.Time
}

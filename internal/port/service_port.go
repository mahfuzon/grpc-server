package port

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
}

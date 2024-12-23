package grpc

import (
	"context"
	"github.com/mahfuzon/grpc-proto/protogen/go/bank"
	"google.golang.org/genproto/googleapis/type/date"
	"time"
)

/**
* Author Company: PT. ASLI ISOAE SOLUSINE
* Author Url: www.isoae.id
* Author Name : Mahfuzon Akhiar
* Author Email : mahfuzon0@gmail.com
* Date : 23/12/24
* Time : 13.52
* Project Name : grpc-server
* Licensed to :
**/

func (a GrpcAdapter) GetCurrentBalance(ctx context.Context, request *bank.CurrentBalanceRequest) (*bank.CurrentBalanceResponse, error) {
	now := time.Now()
	balance := a.bankService.FindCurrentBalance(request.AccountNumber)

	return &bank.CurrentBalanceResponse{
		Amount: balance,
		CurrentDate: &date.Date{
			Year:  int32(now.Year()),
			Month: int32(now.Month()),
			Day:   int32(now.Day()),
		},
	}, nil
}

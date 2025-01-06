package grpc

import (
	"context"
	"github.com/mahfuzon/grpc-proto/protogen/go/bank"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc"
	"log"
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

func (a GrpcAdapter) FetchExchangeRates(req *bank.ExchangeRateRequest, stream grpc.ServerStreamingServer[bank.ExchangeRateResponse]) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("client canceled request")
			return nil
		default:
			now := time.Now().Truncate(time.Second)
			rate := a.bankService.FindExchangeRate(req.FromCurrency, req.ToCurrency, now)
			stream.Send(&bank.ExchangeRateResponse{
				FromCurrency: req.FromCurrency,
				ToCurrency:   req.ToCurrency,
				Rate:         rate,
				Timestamp:    now.Format("2006-01-02 15:04:05"),
			})

			log.Printf("exchange rate send to client form %v to %v with result %v", req.FromCurrency, req.ToCurrency, rate)

			time.Sleep(time.Second * 3)
		}
	}
}

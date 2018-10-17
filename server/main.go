package main

import (
	"context"
	"log"
	"net"

	"github.com/utkarsh17ife/grpc-example/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type Server struct{}

func (*Server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	return &calculatorpb.SumResponse{
		Result: req.GetFirstNum() + req.GetSecondNum(),
	}, nil
}

func (*Server) PrimeDecomp(req *calculatorpb.PrimeDecompRequest, stream calculatorpb.CalculatorService_PrimeDecompServer) error {

	var k int32 = 2
	N := req.GetNumForPrimedecomp()

	for N > 1 {
		if N%k == 0 {
			stream.Send(&calculatorpb.PrimeDecompResponse{Result: k})
			N = N / k
		} else {
			k = k + 1
		}
	}

	return nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Unable to start net listner :%v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to start Calculator service: %v", err)
	}

}

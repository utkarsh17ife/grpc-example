package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/utkarsh17ife/grpc-example/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to the server: %v", err)
	}

	cc := calculatorpb.NewCalculatorServiceClient(conn)

	fmt.Println("Following calculations are happening over GRPC")

	fmt.Printf("2+5=%v\n", calculateSum(cc, 2, 5))
	fmt.Printf("PrimeNumberDecomposition of 120=%v\n", calculatePrimeDecompos(cc, 120))
}

func calculateSum(cc calculatorpb.CalculatorServiceClient, a, b int32) int32 {

	res, err := cc.Sum(context.Background(), &calculatorpb.SumRequest{
		FirstNum:  a,
		SecondNum: b,
	})
	if err != nil {
		log.Fatalf("Server Failed to process sum requet : %v", err)
	}

	return res.GetResult()

}

func calculatePrimeDecompos(cc calculatorpb.CalculatorServiceClient, n int32) []int32 {

	stream, err := cc.PrimeDecomp(context.Background(), &calculatorpb.PrimeDecompRequest{
		NumForPrimedecomp: n,
	})
	if err != nil {
		log.Fatalf("Server Failed to process PrimeDecomp requet : %v", err)
	}
	decomps := make([]int32, 0)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receving the message from server: %v", err)
		}
		decomps = append(decomps, resp.GetResult())
	}
	return decomps
}

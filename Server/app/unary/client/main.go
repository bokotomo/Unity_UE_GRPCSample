package main

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"

	pb "grpc-sample/pb/calc"

	"google.golang.org/grpc"
)

func request(client pb.CalcClient, a, b int32) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()
	sumRequest := pb.SumRequest{
		A: a,
		B: b,
	}
	reply, err := client.Sum(ctx, &sumRequest)
	if err != nil {
		return errors.Wrap(err, "受取り失敗")
	}
	log.Printf("サーバからの受け取り\n %s", reply.GetMessage())
	return nil
}

func sum(a, b int32) error {
	address := "localhost:50051"
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return errors.Wrap(err, "コネクションエラー")
	}
	defer conn.Close()
	client := pb.NewCalcClient(conn)
	return request(client, a, b)
}

func main() {
	a := int32(300)
	b := int32(500)
	if err := sum(a, b); err != nil {
		log.Fatalf("%v", err)
	}
}

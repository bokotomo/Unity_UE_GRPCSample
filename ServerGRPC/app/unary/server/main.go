package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-sample/pb/calc"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const port = ":50051"

// ServerUnary is server
type ServerUnary struct {
	pb.UnimplementedCalcServer
}

// Sum 二つの値を受け取り、合計してクライアントへ返す
func (s *ServerUnary) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumReply, error) {
	a := in.GetA()
	b := in.GetB()
	fmt.Println(a, b)
	reply := fmt.Sprintf("%d + %d = %d", a, b, a+b)
	return &pb.SumReply{
		Message: reply,
	}, nil
}

func set() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return errors.Wrap(err, "ポート失敗")
	}
	s := grpc.NewServer()
	var server ServerUnary
	pb.RegisterCalcServer(s, &server)
	if err := s.Serve(lis); err != nil {
		return errors.Wrap(err, "サーバ起動失敗")
	}
	return nil
}

func main() {
	fmt.Println("起動")
	if err := set(); err != nil {
		log.Fatalf("%v", err)
	}
}

package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/pkg/errors"

	pb "grpc-sample/pb/chat"

	"google.golang.org/grpc"
)

func receive(stream pb.Chat_ChatClient) error {
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("エラー: %v", err)
			}
			log.Printf("サーバから：%s", in.Message)

			// お返し
			stream.Send(&pb.ChatRequest{
				Message: time.Now().Format("2006-01-02 15:04:05"),
			})
		}
	}()
	<-waitc
	return nil
}

func request(stream pb.Chat_ChatClient) error {
	return stream.Send(&pb.ChatRequest{
		Message: "こんにちは",
	})
}

func chat(client pb.ChatClient) error {
	stream, err := client.Chat(context.Background())
	if err != nil {
		return err
	}
	if err := request(stream); err != nil {
		return err
	}
	if err := receive(stream); err != nil {
		return err
	}
	stream.CloseSend()
	return nil
}

func exec() error {
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrap(err, "コネクションエラー")
	}
	defer conn.Close()
	client := pb.NewChatClient(conn)
	return chat(client)
}

func main() {
	if err := exec(); err != nil {
		log.Println(err)
	}
}

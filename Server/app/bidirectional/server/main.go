package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	pb "grpc-sample/pb/chat"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const port = ":50051"

// ServerBidirectional is server
type ServerBidirectional struct {
	pb.UnimplementedChatServer
}

func request(stream pb.Chat_ChatServer, message string) error {
	reply := fmt.Sprintf("%sを受け取ったよ！ありがとう＾＾", message)
	return stream.Send(&pb.ChatReply{
		Message: reply,
	})
}

// Chat クライアントから受け取った言葉に、言葉を返す
func (s *ServerBidirectional) Chat(stream pb.Chat_ChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		message := in.GetMessage()
		fmt.Println("受取：", message)

		if err := request(stream, message); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
}

func set() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return errors.Wrap(err, "ポート失敗")
	}
	s := grpc.NewServer()
	var server ServerBidirectional
	pb.RegisterChatServer(s, &server)
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

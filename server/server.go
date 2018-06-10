package main

import (
	pb "golang-terminal-messenger/proto"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

var user = make(map[string]bool)
var ch = make([]pb.Text, 0)

func (s *server) Enter(ctx context.Context, u *pb.Str) (*pb.Str, error) {
	if user[u.Noti] == true {
		return &pb.Str{Noti: "USER ALREADY EXIST"}, nil
	}
	user[u.Noti] = true
	log.Println(u.Noti, "Entered the room.")
	return &pb.Str{Noti: "Congo, you have entered the room"}, nil
}

func (s *server) Send(ctx context.Context, t *pb.Text) (*pb.Ack, error) {
	ch = append(ch, *t)
	log.Println("In Send: From:", t.Msg.From, "To:", t.Msg.To, "Mess:", t.Msg.Mess)
	return &pb.Ack{Done: true}, nil
}

func (s *server) Receive(ctx context.Context, usr *pb.Str) (*pb.Text, error) {
	for i, txt := range ch {
		if txt.Msg.To == usr.Noti {
			log.Println("From:", txt.Msg.From, "To:", txt.Msg.To, "Mess:", txt.Msg.Mess)
			ch[i] = ch[len(ch)-1]
			ch[len(ch)-1] = pb.Text{}
			ch = ch[:len(ch)-1]
			return &txt, nil
		}
	}
	return &pb.Text{Msg: &pb.TextMail{From: "", To: "", Mess: ""}}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Port is not listening: %v", err)
	}
	defer lis.Close()
	s := grpc.NewServer()
	pb.RegisterChatterServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

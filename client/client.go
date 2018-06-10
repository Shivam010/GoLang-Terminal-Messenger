package main

import (
	"fmt"
	pb "golang-terminal-messenger/proto"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// input takes in the input sending message from the client and is called using a sub-routine
func input(ctx context.Context, usr string, c pb.ChatterClient) {
	// Sender's loop, to send the message to another client
	for {
		var (
			ms string
			to string
		)
		fmt.Printf("Send Message To: ")
		fmt.Scanln(&to)
		fmt.Printf("Enter Message: ")
		fmt.Scanln(&ms)
		msg := &pb.Text{Msg: &pb.TextMail{From: usr, To: to, Mess: ms}}
		c.Send(ctx, msg)
	}
}

func main() {
	// setting up connection
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connection Not established: %v", err)
	}
	defer conn.Close()

	// making connection with the grpc protocol buffer
	c := pb.NewChatterClient(conn)

	ctx := context.Background()
	var usr string

	// registering username to the server
	for {
		fmt.Printf("Enter Username: ")
		fmt.Scanln(&usr)

		str := &pb.Str{Noti: usr}
		r, err := c.Enter(ctx, str)

		if err != nil {
			log.Fatalf("Could not connect: %v", err)
			continue
		}
		er := &pb.Str{Noti: "USER ALREADY EXIST"}
		if r.Noti == er.Noti {
			fmt.Println("Could not connect:", er.Noti)
			continue
		}
		fmt.Println("\t\t\t\t", r.Noti)
		break
	}

	// sub-routine for taking the input message from client
	go input(ctx, usr, c)

	// Receiver's loop, to receive the message sent to client
	for {
		txt, err := c.Receive(ctx, &pb.Str{Noti: usr})
		if err == nil && txt.Msg.From != "" {
			fmt.Println("\n\t\t\t\tFrom:", txt.Msg.From, "To:", txt.Msg.To, "Mess:", txt.Msg.Mess)
		}
	}
}

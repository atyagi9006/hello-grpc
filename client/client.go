package main

import (
	"context"
	"github.com/atyagi9006/hello-grpc/pkg/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	//get a connection by diailing grpc
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewPingClient(conn)
	res, err := client.SayHello(context.Background(), &proto.PingMessage{Greeting: "le greeting ...."})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server : %s \n",res.Greeting)

}

package api

import (
	"context"
	pb "github.com/atyagi9006/hello-grpc/pkg/proto"
	"log"
)

type HelloGRPCService struct{
}

func (svc *HelloGRPCService) SayHello(ctx context.Context, in *pb.PingMessage) (*pb.PingMessage,error){
	log.Printf("Received msg : %s in reuest \n",in.Greeting )
	res:= pb.PingMessage{
		Greeting: "Milgya greeting...bas rehn de....",
	}
	return  &res, nil
}
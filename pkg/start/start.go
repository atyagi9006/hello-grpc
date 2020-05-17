package start

import (
	"fmt"
	"github.com/atyagi9006/hello-grpc/pkg/api"
	pb "github.com/atyagi9006/hello-grpc/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Run() {
	//create a listener on tcp layer
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatal(err)
	}

	// create service hello service
	helloSvc := api.HelloGRPCService{}

	//create grpc service
	grpcServer := grpc.NewServer()
	pb.RegisterPingServer(grpcServer,&helloSvc)

	//start grpc server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

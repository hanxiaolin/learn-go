package rpc

import (
	pb "code.shihuo.cn/gin-demo/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

type server struct {
}

//type PersonServiceServer interface {
//	SayHello(context.Context, *Send) (*Say, error)
//}

func Serv() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5656))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterPersonServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil { //4.Serve()阻塞等待
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) SayHello(ctx context.Context, in *pb.Send) (*pb.Say, error) {
	log.Printf("Received from client: %s", in.Name)
	u := strings.ToUpper(in.Name) //注意.proto文件中字段是小写下划线，这里变成了大驼峰了
	return &pb.Say{Number: u + "..." + in.Name + "..."}, nil
}

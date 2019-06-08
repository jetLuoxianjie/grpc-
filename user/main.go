package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"xxm/user/impl"
	"xxm/user/proto"
)

func main(){

	lis, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 生成一个rpc服务器
	s := grpc.NewServer()
	// 使用pb包调用注册已实现的rpc接口类server
	//svcExport.RegisterExportTaskServer(s, &server{})

	// Register reflection service on gRPC server.
	proto.RegisterHelloServiceServer(s,&impl.UserServer{})
	s.Serve(lis)




}

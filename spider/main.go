package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"xxm/spider/impl"
	"xxm/spider/proto"
)

func main(){

	lis, err := net.Listen("tcp", "127.0.0.1:9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 生成一个rpc服务器
	s := grpc.NewServer()

	proto.RegisterSpiderServiceServer(s,&impl.SpiderServer{})
	s.Serve(lis)




}
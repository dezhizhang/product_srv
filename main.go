package main

import (
	"google.golang.org/grpc"
	"net"
	"sales-product-srv/handler"
	"sales-product-srv/proto"
)

func main() {
	server := grpc.NewServer()
	//proto.RegisterBannerServer(server, &handler.BannerServer{})
	proto.RegisterBrandServer(server, &handler.BrandServer{})
	listen, err := net.Listen("tcp", "127.0.0.1:8086")
	if err != nil {
		panic("failed to listen" + err.Error())

	}
	err = server.Serve(listen)
	if err != nil {
		panic("failed to start grpc" + err.Error())
	}
	//zap.S().Info("服务器运行在:8000端口上")
}

//// 初始化配置文件
//config.Init()
//// 初始化日志
//initialize.Logger()
//
//server := grpc.NewServer()
//proto.RegisterUserServer(server, &handler.UserServer{})
//listen, err := net.Listen("tcp", "127.0.0.1:8000")
//if err != nil {
//panic("failed to listen" + err.Error())
//
//}
//err = server.Serve(listen)
//if err != nil {
//panic("failed to start grpc" + err.Error())
//}
//zap.S().Info("服务器运行在:8000端口上")

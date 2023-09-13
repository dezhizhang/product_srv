package main

import (
	"google.golang.org/grpc"
	"net"
	"sales-product-srv/handler"
	"sales-product-srv/proto"
)

func main() {
	server := grpc.NewServer()
	proto.RegisterProductServer(server, &handler.ProductServer{})
	//proto.RegisterCategoryServer(server, &handler.CategoryServer{})
	////proto.RegisterBannerServer(server, &handler.BannerServer{})
	//proto.RegisterBrandServer(server, &handler.BrandServer{})
	listen, err := net.Listen("tcp", "127.0.0.1:8086")
	if err != nil {
		panic("failed to listen" + err.Error())

	}
	//err = server.Serve(listen)
	//if err != nil {
	//	panic("failed to start grpc" + err.Error())
	//}
	err = server.Serve(listen)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}

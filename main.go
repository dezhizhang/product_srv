package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"sales-product-srv/handler"
	"sales-product-srv/proto"
	"syscall"
)

func main() {
	server := grpc.NewServer()
	proto.RegisterCategoryServer(server, &handler.CategoryServer{})
	//proto.RegisterBannerServer(server, &handler.BannerServer{})
	proto.RegisterBrandServer(server, &handler.BrandServer{})
	listen, err := net.Listen("tcp", "127.0.0.1:8086")
	if err != nil {
		panic("failed to listen" + err.Error())

	}
	//err = server.Serve(listen)
	//if err != nil {
	//	panic("failed to start grpc" + err.Error())
	//}
	go func() {
		err = server.Serve(listen)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("服务注册成功")
	//if err = client.Agent().ServiceDeregister(serviceID); err != nil {
	//	zap.S().Info("注销失败")
	//}
	//zap.S().Info("注销成功")
}

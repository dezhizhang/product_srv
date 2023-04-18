package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"sales-product-srv/proto"
)

func main() {
	var err error
	conn, err := grpc.Dial("127.0.0.1:8086", grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	brandSrvClient := proto.NewBrandClient(conn)
	//for i := 0; i < 1000; i++ {
	//
	//	fmt.Println(rsp)
	//}
	rsp, err := brandSrvClient.CreateBrand(context.Background(), &proto.CreateBrandRequest{
		Id:   fmt.Sprintf("change%d", 1),
		Name: fmt.Sprintf("小米%d", "2"),
		Logo: fmt.Sprintf("https://www.baidu.com"),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(rsp)
	//rsp, err := brandSrvClient.GetBrandList(context.Background(), &proto.BrandRequest{
	//	PageIndex: 1,
	//	PageSize:  10,
	//})
	//
	//if err != nil {
	//	fmt.Printf("查询失败%s", err.Error())
	//	return
	//}
	//fmt.Println(rsp)

}

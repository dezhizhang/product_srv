package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"sales-product-srv/proto"
)

// 调用微服务
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
	//rsp, err := brandSrvClient.CreateBrand(context.Background(), &proto.CreateBrandRequest{
	//	Name: fmt.Sprintf("小米%d", "2"),
	//	Logo: fmt.Sprintf("https://www.baidu.com"),
	//})
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(rsp)
	rsp, err := brandSrvClient.GetBrandList(context.Background(), &proto.BrandRequest{
		PageIndex: 1,
		PageSize:  10,
	})

	if err != nil {
		fmt.Printf("查询失败%s", err.Error())
		return
	}
	fmt.Println(rsp)

}

// consul服务发现

//func Register(address string, port int, name string, tags []string, id string) error {
//	cfg := api.DefaultConfig()
//	//cfg.Address = address
//
//	client, err := api.NewClient(cfg)
//	if err != nil {
//		panic(err)
//	}
//	//生成对应的检查对象
//	check := &api.AgentServiceCheck{
//		HTTP:                           "http://192.168.1.102:8021/health",
//		Timeout:                        "5s",
//		Interval:                       "5s",
//		DeregisterCriticalServiceAfter: "10s",
//	}
//
//	//生成注册对象
//	registration := new(api.AgentServiceRegistration)
//	registration.Name = name
//	registration.ID = id
//	registration.Port = port
//	registration.Tags = tags
//	registration.Address = address
//	registration.Check = check
//
//	err = client.Agent().ServiceRegister(registration)
//	if err != nil {
//		panic(err)
//	}
//	return nil
//}
//
//func main() {
//	Register("192.168.152.92", 8500, "sales-product-srv", []string{"hello", "sales"}, "sales-user-product")
//}

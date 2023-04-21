package test

//func main() {
//	var err error
//	conn, err := grpc.Dial("127.0.0.1:8086", grpc.WithInsecure())
//	if err != nil {
//		log.Printf("连接失败%s", err.Error())
//		return
//	}
//
//	client := proto.NewCategoryClient(conn)
//	list, err := client.GetAllCategoryList(context.Background(), &empty.Empty{})
//	if err != nil {
//		log.Printf(err.Error())
//	}
//
//	fmt.Println(list)
//
//	//client := proto.NewBannerClient(conn)
//	//id, _ := utils.SnowflakeId()
//	//rsp, err := client.CreateBanner(context.Background(), &proto.CreateBannerRequest{
//	//	Index: int32(1),
//	//	Image: fmt.Sprintf("https://www.baidu.com"),
//	//	Id:    id,
//	//	Url:   fmt.Sprintf("https://www.baidu.com"),
//	//})
//	//
//	//if err != nil {
//	//	log.Printf("创建失败%s", err.Error())
//	//}
//	//fmt.Printf("创建成功%s", rsp)
//	//for i := 0; i < 10; i++ {
//	//	id, _ := utils.SnowflakeId()
//	//	rsp, err := client.CreateBanner(context.Background(), &proto.CreateBannerRequest{
//	//		Index: int32(i),
//	//		Image: fmt.Sprintf("https://www.baidu.com"),
//	//		Id:    id,
//	//		Url:   fmt.Sprintf("https://www.baidu.com"),
//	//	})
//	//
//	//	if err != nil {
//	//		log.Printf("创建失败%s", err.Error())
//	//	}
//	//	fmt.Printf("创建成功%s", rsp)
//	//
//	//}
//
//}

// 调用微服务
//func main() {
//	var err error
//	conn, err := grpc.Dial("127.0.0.1:8086", grpc.WithInsecure())
//	if err != nil {
//		fmt.Println("连接失败")
//		return
//	}
//	brandSrvClient := proto.NewBrandClient(conn)
//	//for i := 0; i < 1000; i++ {
//	//
//	//	fmt.Println(rsp)
//	//}
//	//rsp, err := brandSrvClient.CreateBrand(context.Background(), &proto.CreateBrandRequest{
//	//	Name: fmt.Sprintf("小米%d", "2"),
//	//	Logo: fmt.Sprintf("https://www.baidu.com"),
//	//})
//	//if err != nil {
//	//	fmt.Println(err.Error())
//	//}
//	//fmt.Println(rsp)
//	rsp, err := brandSrvClient.GetBrandList(context.Background(), &proto.BrandRequest{
//		PageIndex: 1,
//		PageSize:  10,
//	})
//
//	if err != nil {
//		fmt.Printf("查询失败%s", err.Error())
//		return
//	}
//	fmt.Println(rsp)
//
//}

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

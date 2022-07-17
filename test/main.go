package main

//func main() {
//	var err error
//	conn, err := grpc.Dial("127.0.0.1:8086", grpc.WithInsecure())
//	if err != nil {
//		fmt.Println("连接失败")
//		return
//	}
//	brandSrvClient := proto.NewBrandClient(conn)
//	rsp, err := brandSrvClient.GetBrandList(context.Background(), &proto.BrandRequest{
//		PageIndex: 1,
//		PageSize:  10,
//	})
//
//	if err != nil {
//		fmt.Printf("查询失败%s", err.Error())
//		return
//	}
//
//}

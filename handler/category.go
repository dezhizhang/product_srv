package handler

import (
	"context"
	"encoding/json"
	"sales-product-srv/driver"

	"google.golang.org/protobuf/types/known/emptypb"
	"sales-product-srv/model"
	"sales-product-srv/proto"
)

type CategoryServer struct {
	proto.UnimplementedCategoryServer
}

func (c *CategoryServer) GetAllCategoryList(ctx context.Context, empty *emptypb.Empty) (*proto.CategoryListResponse, error) {
	var category []model.Category
	driver.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&category)

	b, _ := json.Marshal(&category)
	return &proto.CategoryListResponse{JsonData: string(b)}, nil
}

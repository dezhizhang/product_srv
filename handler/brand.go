package handler

import (
	"context"
	"sales-product-srv/driver"
	"sales-product-srv/model"
	"sales-product-srv/proto"
	"sales-product-srv/utils"
)

type BrandServer struct {
}

func (b *BrandServer) GetBrandList(ctx context.Context, req *proto.BrandRequest) (*proto.BrandResponseList, error) {
	var count int64
	var brands []model.Brand
	var brandsList []*proto.BrandResponse
	result := driver.DB.Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}

	rsp := &proto.BrandResponseList{}

	//分页
	driver.DB.Scopes(utils.Paginate(int(req.PageIndex), int(req.PageSize))).Find(&brands)
	//查询所有总数
	driver.DB.Model(&model.Brand{}).Count(&count)

	for _, value := range brands {
		brandsList = append(brandsList, &proto.BrandResponse{
			Id:   value.Id,
			Name: value.Name,
			Logo: value.Logo,
		})
	}
	rsp.Total = int32(count)
	rsp.Data = brandsList
	return rsp, nil
}

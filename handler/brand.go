package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"sales-product-srv/global"
	"sales-product-srv/model"
	"sales-product-srv/proto"
	"sales-product-srv/utils"
	"time"
)

type BrandServer struct {
	proto.UnimplementedBrandServer
}

// 获取品牌列表

func (b *BrandServer) GetBrandList(ctx context.Context, req *proto.BrandRequest) (*proto.BrandResponseList, error) {
	var count int64
	var brands []model.Brands
	var brandsList []*proto.BrandResponse
	result := global.DB.Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}

	rsp := &proto.BrandResponseList{}

	//分页
	global.DB.Scopes(utils.Paginate(int(req.PageIndex), int(req.PageSize))).Find(&brands)
	//查询所有总数
	global.DB.Model(&model.Brands{}).Count(&count)

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

// 创建品牌

func (b *BrandServer) CreateBrand(ctx context.Context, req *proto.CreateBrandRequest) (*empty.Empty, error) {
	var brand model.Brands
	var err error
	if result := global.DB.Where("name=?", req.Name).Find(&brand); result.RowsAffected == 1 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌已存在")
	}
	id, err1 := utils.SnowflakeId()
	if err1 != nil {
		log.Printf("创建品牌%s", err.Error())
		return nil, err
	}
	brand.Id = id
	brand.Name = req.Name
	brand.Logo = req.Logo
	brand.DeletedAt = time.Now()
	err = global.DB.Save(&brand).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &empty.Empty{}, nil
}

// 删除品牌

func (b *BrandServer) DeleteBrand(ctx context.Context, req *proto.DeleteBrandRequest) (*empty.Empty, error) {
	var brand model.Brands
	result := global.DB.Where("id =? ", req.Id).Delete(&brand)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.Internal, "删除失败")
	}
	return &empty.Empty{}, nil
}

func (b *BrandServer) UpdateBrand(ctx context.Context, req *proto.UpdateBrandRequest) (*empty.Empty, error) {
	var brand model.Brands
	result := global.DB.Where("id = ?", req.Id).Find(&brand)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.Internal, "更新的品牌不存在")
	}
	brand.Name = req.Name
	brand.Logo = req.Logo
	err := global.DB.Save(&brand).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "更新品牌失败")
	}
	return &empty.Empty{}, nil
}

package handler

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sales-product-srv/global"
	"sales-product-srv/model"
	"sales-product-srv/proto"
	"sales-product-srv/utils"
	"time"
)

type BannerServer struct {
	proto.UnimplementedBannerServer
}

// CreateBanner 创建轮播图
func (b *BannerServer) CreateBanner(ctx context.Context, req *proto.CreateBannerRequest) (*empty.Empty, error) {

	banner := model.Banner{
		Name:        req.Name,
		Link:        req.Link,
		Url:         req.Url,
		Status:      req.Status,
		Position:    req.Position,
		Description: req.Description,
	}
	bannerId, _ := utils.SnowflakeId()

	banner.Id = bannerId
	banner.CreatedAt = time.Now()
	banner.UpdatedAt = time.Now()
	banner.DeletedAt = time.Now()

	err := global.DB.Save(&banner).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &empty.Empty{}, nil
}

// 获取轮播图列表

func (b *BannerServer) GetBannerList(ctx context.Context, req *proto.BannerRequest) (*proto.BannerResponseList, error) {
	var count int64
	var banners []model.Banner
	var bannerList []*proto.BannerResponse
	result := global.DB.Find(&banners)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "查询失败")
	}

	// 分页
	global.DB.Scopes(utils.Paginate(int(req.PageIndex), int(req.PageSize))).Find(&banners)
	//查询所有总数
	global.DB.Model(&model.Banner{}).Count(&count)
	for _, value := range banners {
		bannerList = append(bannerList, &proto.BannerResponse{
			Id:          value.Id,
			Name:        value.Name,
			Link:        value.Link,
			Url:         value.Url,
			Status:      value.Status,
			Position:    value.Position,
			Description: value.Description,
		})
	}

	fmt.Println("data", bannerList)
	fmt.Println("Total", count)

	return &proto.BannerResponseList{Data: bannerList, Total: int32(count)}, nil
}

// 删除轮播图

func (b *BannerServer) DeleteBanner(ctx context.Context, req *proto.DeleteBannerRequest) (*empty.Empty, error) {
	var banner model.Banner
	if result := global.DB.Where("id=?", req.Id).Find(&banner); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "轮播图不存在")
	}
	if result := global.DB.Where("id=?", req.Id).Delete(&banner); result.RowsAffected != 1 {
		return nil, status.Errorf(codes.Internal, "删除失败")
	}
	return &empty.Empty{}, nil
}

// UpdateBanner 更新轮播图
func (b *BannerServer) UpdateBanner(ctx context.Context, req *proto.UpdateBannerRequest) (*empty.Empty, error) {
	//var banner model.Banner
	result := global.DB.Where("id = ?", req.Id).Find(&model.Banner{})
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.Internal, "更新的轮播图不存在")
	}

	banner := &model.Banner{
		Name:        req.Name,
		Link:        req.Link,
		Url:         req.Url,
		Position:    req.Position,
		Status:      req.Status,
		Description: req.Description,
	}
	banner.Id = req.Id
	banner.UpdatedAt = time.Now()

	err := global.DB.Updates(&banner).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "更新轮播图失败")
	}
	return &empty.Empty{}, nil
}

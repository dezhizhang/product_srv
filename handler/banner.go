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

type BannerServer struct {
	proto.UnimplementedBannerServer
}

// 获取轮播图

func (b *BannerServer) CreateBanner(ctx context.Context, req *proto.CreateBannerRequest) (*empty.Empty, error) {
	var banner model.Banner
	var err error

	id, err1 := utils.SnowflakeId()
	if err1 != nil {
		log.Printf("创建轮播图%s", err.Error())
		return nil, err
	}
	banner.Id = id
	banner.Image = req.Image
	banner.Url = req.Url
	banner.Index = req.Index
	banner.CreatedAt = time.Now()
	banner.DeletedAt = time.Now()
	err = global.DB.Save(&banner).Error
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

	rsp := &proto.BannerResponseList{}

	// 分页
	global.DB.Scopes(utils.Paginate(int(req.PageIndex), int(req.PageSize))).Find(&banners)
	//查询所有总数
	global.DB.Model(&model.Brands{}).Count(&count)
	for _, value := range banners {
		bannerList = append(bannerList, &proto.BannerResponse{
			Id:    value.Id,
			Image: value.Image,
			Index: value.Index,
			Url:   value.Url,
		})
	}
	rsp.Total = int32(count)
	rsp.Data = bannerList

	return rsp, nil
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

// 更新轮播图

func (b *BannerServer) UpdateBanner(ctx context.Context, req *proto.UpdateBannerRequest) (*empty.Empty, error) {
	var banner model.Banner
	result := global.DB.Where("id = ?", req.Id).Find(&banner)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.Internal, "更新的轮播图不存在")
	}
	banner.Id = req.Id
	banner.Url = req.Url
	banner.Image = req.Image
	banner.Index = req.Index
	banner.UpdatedAt = time.Now()

	err := global.DB.Save(&banner).Error
	if err != nil {
		return nil, status.Errorf(codes.Internal, "更新轮播图失败")
	}
	return &empty.Empty{}, nil

}

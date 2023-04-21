package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sales-product-srv/driver"
	"sales-product-srv/model"
	"sales-product-srv/proto"
	"strconv"
)

type ProductServer struct {
	proto.UnimplementedProductServer
}

func ModelToResponse(product model.Product) proto.ProductInfoResponse {
	productId, _ := strconv.ParseInt(product.Id, 10, 32)
	categoryId, _ := strconv.ParseInt(product.Category.Id, 10, 32)
	brandsId, _ := strconv.ParseInt(product.Brands.Id, 10, 32)
	return proto.ProductInfoResponse{
		Id:           int32(productId),
		CategoryId:   product.CategoryId,
		Name:         product.Name,
		ProductSn:    product.ProductSn,
		ClickNum:     product.ClickNum,
		SoldNum:      product.SoldNum,
		FavNum:       product.FavNum,
		MarketPrice:  product.MarketPrice,
		ShopPrice:    product.ShopPrice,
		ProductBrief: product.Description,
		ShipFree:     product.ShipFree,
		IsNew:        product.IsNew,
		IsHot:        product.IsHot,
		OnSale:       product.OnSale,
		DescImages:   product.DescImages,
		Images:       product.Images,
		Category: &proto.CategoryBriefInfoResponse{
			Id:   int32(categoryId),
			Name: product.Category.Name,
		},
		Brand: &proto.BrandInfoResponse{
			Id:   int32(brandsId),
			Name: product.Brands.Name,
			Logo: product.Brands.Logo,
		},
	}
}

func (p *ProductServer) ProductList(ctx context.Context, req *proto.ProductFilterRequest) (*proto.ProductListResponse, error) {
	var product []model.Product
	localDB := driver.DB.Model(model.Product{})
	if req.KeyWords != "" {
		localDB = localDB.Where("name LIKE ? ", "%"+req.KeyWords+"%")
	}

	if req.IsHot {
		localDB = localDB.Where(model.Product{IsHot: true})
	}
	if req.IsNew {
		localDB = localDB.Where(model.Product{IsNew: true})
	}

	localDB.Find(&product)

	return nil, nil

}

// 批量获取商品

func (p *ProductServer) BatchGetProduct(ctx context.Context, req *proto.BatchProductIdInfo) (*proto.ProductListResponse, error) {
	var products []model.Product
	productListResponse := proto.ProductListResponse{}
	result := driver.DB.Where(&products, req.Id)

	for _, product := range products {
		productInfoResponse := ModelToResponse(product)
		productListResponse.Data = append(productListResponse.Data, &productInfoResponse)
	}
	productListResponse.Total = int32(result.RowsAffected)
	return &productListResponse, nil
}

// 获取商品详情

func (p *ProductServer) GetProductDetail(ctx context.Context, req *proto.ProductInfoRequest) (*proto.ProductInfoResponse, error) {
	var product model.Product
	result := driver.DB.First(&product, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}
	productInfoResponse := ModelToResponse(product)
	return &productInfoResponse, nil
}

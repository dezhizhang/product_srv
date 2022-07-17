package handler

import (
	"context"
	"sales-product-srv/proto"
)

type BannerServer struct {
}

func (b *BannerServer) GetBannerList(ctx context.Context, request *proto.BannerRequest) (*proto.BannerResponseList, error) {
	return nil, nil
}

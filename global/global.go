package global

import (
	"gorm.io/gorm"
	"sales-product-srv/model"
)

var (
	NacosConfig  *model.NacosConfig  = &model.NacosConfig{}
	ServerConfig *model.ServerConfig = &model.ServerConfig{}
	DB           *gorm.DB
)

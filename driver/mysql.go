package driver

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sales-product-srv/model"
)

var DB *gorm.DB

func InitDB() (err error) {
	dsn := "root:sdf@df%%$65#fdsbXT@tcp(127.0.0.1:3306)/product_srv?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		return err
	}
	DB = db
	return nil
}

func init() {
	InitDB()

	err := DB.AutoMigrate(model.Banner{}, model.Brand{}, model.Category{}, model.Product{})
	if err != nil {
		fmt.Printf("初始化表失败%s", err.Error())
	}
}

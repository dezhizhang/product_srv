package model

type Brand struct {
	BaseModel
	Name string `gorm:"type:varchar(32);not null" json:"name"`
	Logo string `gorm:"type:varchar(200);not null" json:"logo"`
}

type ProductCategoryBrand struct {
	BaseModel
	Category   Category `json:"category"`
	CategoryId int32    `gorm:"type:int;index:idx:category_brand,unique" json:"categoryId"`
	Brand      Brand    `json:"brand"`
	BrandId    int32    `gorm:"type:int;idx:category_brand,unique" json:"brandId"`
}

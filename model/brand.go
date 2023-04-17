package model

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(32);not null" json:"name"`
	Logo string `gorm:"type:varchar(200);default:'';not null" json:"logo"`
}

type ProductCategoryBrand struct {
	BaseModel
	Category   Category `json:"category"`
	CategoryId int32    `gorm:"type:int;index:idx_category_brand,unique" json:"categoryId"`
	Brands     Brands   `json:"brand"`
	BrandId    int32    `gorm:"type:int;index:idx_category_brand,unique" json:"brandId"`
}

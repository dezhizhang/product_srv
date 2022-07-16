package model

type Category struct {
	BaseModel
	ParentCategoryId string    `json:"ParentCategoryId"`
	ParentCategory   *Category `json:"parentCategory"`
	Name             string    `gorm:"type:varchar(42);not null" json:"name"`
	Level            int32     `gorm:"type:int;not null;default:1" json:"level"`
	IsTab            bool      `gorm:"default:false;not null"`
}

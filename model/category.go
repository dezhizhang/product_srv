package model

type Category struct {
	BaseModel
	Name             string    `gorm:"type:varchar(64);not null" json:"name"`
	ParentCategoryId int32     `json:"parentCategoryId"`
	ParentCategory   *Category `json:"productCategory"`
	Level         ss   int32     `gorm:"type:int;not null;default:1" json:"level"`
	IsTab            bool      `gorm:"type:bool;default:false;not null" json:"isTab"`
}

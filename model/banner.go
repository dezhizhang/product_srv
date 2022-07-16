package model

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null" json:"image"`
	Url   string `gorm:"type:varchar(200);not null" json:"url"`
	Index int32  `gorm:"type:int;default:1;not null" json:"index"`
}

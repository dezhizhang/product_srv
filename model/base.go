package model

import "time"

type BaseModel struct {
	Id        string    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updateAt" json:"updatedAt"`
	DeletedAt time.Time `gorm:"column:deletedAt" json:"deletedAt"`
	IsDeleted bool
}

package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type BaseModel struct {
	Id        string    `gorm:"primaryKey;type:varchar(200)"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updateAt" json:"updatedAt"`
	DeletedAt time.Time `gorm:"column:deletedAt" json:"deletedAt"`
	IsDeleted bool
}

type GormList []string

func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

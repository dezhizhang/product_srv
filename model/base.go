package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type BaseModel struct {
	Id        string    `gorm:"primaryKey;type:varchar(200)" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:update_at" json:"update_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	IsDeleted bool      `json:"is_deleted"`
}

type GormList []string

func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

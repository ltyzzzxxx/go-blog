package models

import "time"

type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index" json:"created-at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index" json:"updated-at,omitempty"`
}
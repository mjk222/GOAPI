// Package models 模型通用属性和方法
package models

import "time"

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// CommonTimetampsField 时间戳
type CommonTimetampsField struct {
	CreatedAt  time.Time `gorm:"column:create_at;index;" json:"create_at,omitempty"`
	UpdateedAt time.Time `gorm:"column:update_at;index;" json:"updat_at,omitempty"`
}

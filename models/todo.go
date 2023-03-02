package models

import (
	"time"

	"gorm.io/gorm"
)

type Priority string

const (
	VeryHigh Priority = "very-high"
	High     Priority = "high"
	Normal   Priority = "normal"
	Low      Priority = "low"
	VeryLow  Priority = "very-low"
)

type Todo struct {
	Id              uint     `gorm:"primaryKey" json:"id"`
	Title           string   `gorm:"type:varchar(255);not null" json:"title" binding:"required"`
	IsActive        bool     `gorm:"type:boolean" json:"is_active"`
	Priority        Priority `gorm:"type:enum('very-high', 'high', 'normal', 'low', 'very-low');default:very-high" json:"priority"`
	ActivityGroupId uint     `json:"activity_group_id"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

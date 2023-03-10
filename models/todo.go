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
	IsActive        bool     `gorm:"type:boolean;default:true" json:"is_active"`
	Priority        Priority `gorm:"type:enum('very-high', 'high', 'normal', 'low', 'very-low');default:very-high" json:"priority"`
	ActivityGroupId uint     `json:"activity_group_id" binding:"required"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type TodoUpdate struct {
	Title           string   `json:"title"`
	IsActive        *bool     `json:"is_active"`
	Priority        *Priority `json:"priority"`
	ActivityGroupId uint     `json:"activity_group_id"`
}
package models

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	Id        uint   `gorm:"primaryKey" json:"id"`
	Email     string `gorm:"type:varchar(255)" json:"email"`
	Title     string `gorm:"type:varchar(255);not null" json:"title" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

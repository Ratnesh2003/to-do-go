package models

import (
	"time"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TestField string
	Title       string
	Description string
	Deadline    time.Time
	IsCompleted bool `gorm:"default:false"`
}

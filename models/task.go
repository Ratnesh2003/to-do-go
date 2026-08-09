package models

import (
	"time"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TestNum int64
	TestField string
	Title       string
	Description string
	Deadline    time.Time
	IsCompleted bool `gorm:"default:false"`
}

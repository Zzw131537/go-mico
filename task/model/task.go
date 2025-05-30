package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"not null"`
	Status    int64  `gorm:"default:'0'"`
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}

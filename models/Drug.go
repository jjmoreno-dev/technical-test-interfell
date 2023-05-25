package models

import (
	"time"
)

type Drug struct {
	ID          uint64     `gorm:"primaryKey;autoIncrement:true"`
	Name        string     `gorm:"type:varchar(255);not null"`
	Approved    bool       `gorm:"not null;defult=false"`
	MinDose     string     `gorm:"not null"`
	MaxDose     string     `gorm:"not null"`
	AvailableAt time.Time  `gorm:"not null"`
	CreatedAt   time.Time  `gorm:"not null"`
	UpdatedAt   *time.Time `gorm:"default:null"`
}

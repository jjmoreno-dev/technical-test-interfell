package models

import (
	"time"
)

type Vaccination struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name   string `gorm:"type:varchar(255);not null"`
	DrugId uint64
	Drug   Drug      `gorm:"foreignKey:DrugId;references:ID"`
	dose   int       `gorm:"not null"`
	Date   time.Time `gorm:"not null"`
}

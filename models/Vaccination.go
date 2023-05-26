package models

import (
	"time"
)

type Vaccination struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name      string `gorm:"type:varchar(255);not null"`
	DrugId    uint64
	Drug      Drug       `gorm:"foreignKey:drug_id;references:id"`
	Dose      int        `gorm:"not null"`
	Date      time.Time  `gorm:"not null"`
	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt *time.Time `gorm:"default:null"`
}

/*
type User struct {
  orm.Model
  Name  string
  Phone   *Phone `gorm:"foreignKey:UserName;references:name"`
}

type Phone struct {
  orm.Model
  UserName string
  Name   string
}

*/

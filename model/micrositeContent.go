package models

import "time"

type MicrositeContent struct {
	Id        int    `gorm:"primary_key;column:Id"`
	Content   string `gorm:"column:Content"`
	Type 	  string `gorm:"column:Type"`
	Tenant	  string `gorm:"column:Tenant"`
	Status 	  int `gorm:"column:Status"`
	IsRequest int `gorm:"column:IsRequest"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
}

func (MicrositeContent) TableName() string {
	return "MicrositeContent"
}
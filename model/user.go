package models

type User struct {
	Userid  string `gorm:"primary_key;column:Userid"`
	Name   	string `gorm:"column:Name"`
}

func (User) TableName() string {
	return "users"
}

package models

type User struct {
	Id uint
	Email string `gorm:"unique"`
	Password []byte
}


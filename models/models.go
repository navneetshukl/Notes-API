package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

/*type Notes struct {
	gorm.Model
	Email       string
	Title       string `gorm:"unique"`
	Description string
}*/
type Notes struct {
	gorm.Model
	Email       string
	Heading     string `gorm:"unique"`
	Title       string
	Description string
}

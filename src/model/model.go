package model

import (
	"gorm.io/gorm"
)

// User model.
type User struct {
	gorm.Model
	Account  string
	Name     string
	Password string
}

// Diary model.
type Diary struct {
	gorm.Model
	UserId   int
	Date     int
	Title    string
	TargetId int
	Text     string
}

// Target model.
type Target struct {
	gorm.Model
	UserId       int
	Name         string
	Birthday     int
	Relationship string
	Profile      string
}

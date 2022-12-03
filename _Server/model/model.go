package model

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}
type Companies struct {
	gorm.Model
	Companie string `gorm:"not null;size:30"`
}
type Todos struct {
	gorm.Model
	Todoname string `gorm:"not null;size:30"`
}

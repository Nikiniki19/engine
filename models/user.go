package models

import "gorm.io/gorm"

type EngineUsers struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}

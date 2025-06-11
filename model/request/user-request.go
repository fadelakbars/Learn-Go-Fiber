package request

import (
	// "gorm.io/gorm"
)

type UserRequest struct {
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Address  string `json:"address" gorm:"not null"`			
	Phone    string `json:"phone" gorm:"not null"`
}
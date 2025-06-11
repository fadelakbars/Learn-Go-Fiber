package request

import (
	// "gorm.io/gorm"
)

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}
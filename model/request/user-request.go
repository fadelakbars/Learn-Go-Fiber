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

type UserUpdateRequest struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UserEmailRequest struct {
	Email    string `json:"email" validate:"required"`
}
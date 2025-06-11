package entity

import "time"

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Address  string `json:"address" gorm:"not null"`			
	Phone    string `json:"phone" gorm:"not null"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	DekletedAt   time.Time    `json:"deleted_at" gorm:"index"`
}
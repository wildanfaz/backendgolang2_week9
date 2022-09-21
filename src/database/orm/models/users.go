package models

import (
	"time"
)

type User struct {
	UserId       string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"user_id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Password     string    `json:"password,omitempty"`
	Gender       string    `json:"gender,omitempty"`
	Address      string    `json:"address,omitempty"`
	MobileNumber string    `json:"mobile_number,omitempty"`
	DisplayName  string    `json:"display_name,omitempty"`
	BirthDate    string    `json:"birth_date,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Users []User

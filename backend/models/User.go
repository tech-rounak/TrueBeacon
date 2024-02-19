package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name" validate:"required"`
	UserName  string    `gorm:"unique" json:"user_name"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Share struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Date           time.Time `json:"date"`
	Price          float64   `json:"price"`
	InstrumentName string    `json:"instrument_name"`
}

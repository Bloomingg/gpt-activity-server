package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null"`
	Total       int    `json:"total" gorm:"not null"`
	Current     int    `json:"current" gorm:"default:0;not null" `
	StartTime   uint   `json:"start_time" gorm:"not null" `
	EndTime     uint   `json:"end_time" gorm:"not null" `
	Location    string `json:"location" gorm:"not null" `
	Latng       string `json:"latng"`
	Description string `json:"description"`
	Banner      string `json:"banner"`
	Poster      string `json:"poster"`
}

type RespActivity struct {
	Title       string    `json:"title"`
	Total       int       `json:"total"`
	Current     int       `json:"current"`
	StartTime   uint      `json:"start_time"`
	EndTime     uint      `json:"end_time"`
	Location    string    `json:"location"`
	Latng       string    `json:"latng"`
	Description string    `json:"description"`
	Banner      string    `json:"banner"`
	Poster      string    `json:"poster"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

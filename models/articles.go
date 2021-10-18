package models

import (
	"time"
)

// Articles ...
type Articles struct {
	ID      int       `json:"id" gorm:"primary_key" tag:"auto_increment"`
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
	// Status string `json:"status" binding:"required"`
}

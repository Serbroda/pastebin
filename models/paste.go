package models

import (
	"gorm.io/gorm"
	"time"
)

type Paste struct {
	ID           uint           `gorm:"primarykey" json:"-"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	SID          string         `gorm:"not null;unique" json:"id"`
	Title        string         `gorm:"not null" json:"title"`
	Content      string         `json:"content"`
	AutoDeleteAt time.Time      `json:"autoDeleteAt"`
	Password     string         `json:"-"`
	SessionID    string         `json:"-"`
}

type CreatePasteDto struct {
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Password     string    `json:"password"`
	AutoDeleteAt time.Time `json:"autoDeleteAt"`
}

type UpdatePasteDto struct {
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Password     string    `json:"password"`
	AutoDeleteAt time.Time `json:"autoDeleteAt"`
}

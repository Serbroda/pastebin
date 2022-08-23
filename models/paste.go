package models

import (
	"gorm.io/gorm"
	"time"
)

type Paste struct {
	ID        uint           `gorm:"primarykey" json:"-"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	SID       string         `gorm:"not null;unique" json:"id"`
	Content   string         `json:"content"`
	Password  string         `json:"-"`
}

type CreatePasteDto struct {
	Content  string `json:"content"`
	Password string `json:"password"`
}

type UpdatePasteDto struct {
	Content  string `json:"content"`
	Password string `json:"password"`
}

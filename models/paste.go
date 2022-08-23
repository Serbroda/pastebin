package models

import "gorm.io/gorm"

type Paste struct {
	gorm.Model
	Content string
}

type CreatePasteDto struct {
	Content string
}

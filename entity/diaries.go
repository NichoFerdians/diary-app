package entity

import (
	"time"

	"gorm.io/gorm"
)

type Diary struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	UserID    int            `json:"user_id"`
	User      User           `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type CreateDiaryInput struct {
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body"`
	UserID int    `json:"user_id"`
}

type UpdateDiaryInput struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"user_id"`
}

type DairyResponse struct {
	ID        int            `json:"id,omitempty"`
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	UserID    int            `json:"user_id,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Diary) TableName() string {
	return "diaries"
}

package user

import (
	"time"

	"gorm.io/gorm"
)

type School struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name      string         `json:"name"`
	UserID    uint
}

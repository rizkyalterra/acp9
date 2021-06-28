package user

import (
	"time"
)

type UserResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	School    School    `gorm:"foreignKey:ID;references:School_id" json:"-"`
}

type ResponseUser struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type ResponseUserSingle struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
}

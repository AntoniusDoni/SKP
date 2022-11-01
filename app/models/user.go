package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserLoginReq struct {
	Usermail string `json:"Usermail"`
	Password string `json:"Password"`
}

type UserReq struct {
	IsLogin      bool
	RefreshToken string
}
type User struct {
	ID           uuid.UUID `gorm:"primarykey;type:uuid"`
	UserName     string    `gorm:"not null"`
	Name         string    `gorm:"not null"`
	Email        string    `gorm:"not null"`
	Password     string    `gorm:"not null"`
	IsLogin      *bool
	RefreshToken string
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}

package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Quisioner struct {
	ID          uuid.UUID `gorm:"primarykey;type:uuid"`
	IdQuis      uuid.UUID `gorm:"not null;type:uuid"`
	Name        string    `gorm:"not null"`
	Phone       string
	ResultQuist string
	Note        string
}

type ReqQuisioner struct {
	ID          uuid.UUID
	IdQuis      uuid.UUID `form:"idquis"`
	Name        string    `form:"name"`
	Phone       string    `form:"phone"`
	ResultQuist string    `form:"radio1"`
	Note        string    `form:"noted"`
}
type ResponseCart struct {
	Label string `json:"label"`
	Data  string `json:"data"`
	Color string `json:"color"`
}

func (quis *Quisioner) BeforeCreate(tx *gorm.DB) (err error) {
	quis.ID = uuid.New()
	return
}

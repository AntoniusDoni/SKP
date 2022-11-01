package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Quis struct {
	IdQuis   uuid.UUID `gorm:"primarykey;type:uuid"`
	Question string    `form:"question"`
}
type RequestQuis struct {
	// IdQuis   uuid.UUID
	Question string `form:"question"`
}

func (quis *Quis) BeforeCreate(tx *gorm.DB) (err error) {
	quis.IdQuis = uuid.New()
	return
}

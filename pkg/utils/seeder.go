package utils

import (
	"github.com/skp/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUser() *[]models.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	users := []models.User{
		{
			UserName: "admin",
			Name:     "Super Admin",
			Email:    "admin@gmail.com",
			Password: string(hashedPassword),
		},
	}
	return &users
}

func SeedQuis() *[]models.Quis {
	quis := []models.Quis{
		{
			Question: "Apakah Anda Puas dengan Pelayanan Kami?",
		},
	}
	return &quis
}

func Seed(db *gorm.DB) {
	tx := db.Begin()
	var users *[]models.User
	var quis *[]models.Quis
	db.Find(&users)
	if len(*users) <= 0 {
		users = SeedUser()
		tx.Create(users)
	}
	db.Find(&quis)
	if len(*quis) <= 0 {
		quis = SeedQuis()
		tx.Create(quis)
	}
	tx.Commit()
}

package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/skp/app/models"
)

func (repo Repository) GetUserEmailPass(email string, user *models.User) error {
	db, _ := repo.Gormdb.GetInstanceConnect()
	err := db.Where("(email=? OR user_name=?) ", email, email).First(&user).Error
	return err
}
func (repo Repository) GetUserEmail(email string, user *models.User) error {
	db, _ := repo.Gormdb.GetInstanceConnect()
	err := db.Where("(email=? OR user_name=?)", email, email).First(&user).Error
	return err
}
func (repo Repository) GetUserByRefreshToken(refreshToken string, user *models.User) error {
	db, _ := repo.Gormdb.GetInstanceConnect()
	err := db.Where("refresh_token=?", refreshToken).First(&user).Error
	// fmt.Println(err)
	return err
}
func (repo Repository) UpdateUserById(userId uuid.UUID, authtoken string, refreshToken string, user *models.User) error {
	db, _ := repo.Gormdb.GetInstanceConnect()
	mp := make(map[string]interface{})
	mp["RefreshToken"] = refreshToken
	mp["IsLogin"] = true
	err := db.Model(&user).Where("id=?", userId).UpdateColumns(mp).Error
	var expired time.Duration = 30
	repo.RedisClient.Set(authtoken, user, expired*time.Minute)
	return err

}

func (repo Repository) UpdateByRefreshToken(refreshToken string, usr *models.User) (rowaff int64, erros error) {
	db, _ := repo.Gormdb.GetInstanceConnect()
	mp := make(map[string]interface{})
	mp["RefreshToken"] = ""
	mp["IsLogin"] = false
	result := db.Model(&usr).Where("refresh_token=?", refreshToken).UpdateColumns(mp)
	// repo.RedisClient.Remove()
	rowaff = result.RowsAffected
	erros = result.Error
	return rowaff, erros
}

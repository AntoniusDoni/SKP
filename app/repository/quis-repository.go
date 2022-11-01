package repository

import (
	"github.com/skp/app/models"
)

func (repo Repository) GetRowQuis(quis *models.Quis) (error, *models.Quis) {
	db, _ := repo.Gormdb.GetInstanceConnect()
	err := db.Last(&quis).Error
	return err, quis
}
func (repo Repository) CreateSKP(param *models.ReqQuisioner) error {
	db, _ := repo.Gormdb.GetInstanceConnect()
	err := db.Table("quisioners").Create(&param).Error
	return err
}

func (repo Repository) CreateQuis(reqparam *models.Quis) error {
	db, _ := repo.Gormdb.GetInstanceConnect()
	err := db.Create(reqparam).Error
	return err
}

func (repo Repository) GetListQuisioner(quisioners *[]models.Quisioner) error {
	db, _ := repo.Gormdb.GetInstanceConnect()
	err := db.Find(quisioners).Error
	return err
}

func (repo Repository) GetResultQuisioner(datacart *[]models.ResponseCart) {
	db, _ := repo.Gormdb.GetInstanceConnect()
	db.Select("COUNT(result_quist) as data,result_quist as label,concat('#',SUBSTRING((lpad(hex(round(rand() * 10000000)),6,0)),-6)) as color").Group("result_quist").Table("quisioners").Scan(&datacart)

}

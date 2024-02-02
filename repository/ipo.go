package repository

import (
	"log"

	"github.com/gilangmahardhika/ipo-exchange/model"
	"gorm.io/gorm"
)

type IpoRepo struct {
	DB *gorm.DB
}

func NewIpoRepository(db *gorm.DB) IpoRepository {
	return &IpoRepo{
		DB: db,
	}
}

type IpoRepository interface {
	CreateOrUpdateIPO(ipo *model.Ipo) error
}

func (repo *IpoRepo) CreateOrUpdateIPO(ipo *model.Ipo) error {

	err := repo.DB.Where("id = ?", &ipo.ID).First(&ipo).Error
	if err == nil {
		return repo.updateIPO(ipo)
	}

	return repo.createIPO(ipo)
}

func (repo *IpoRepo) updateIPO(ipo *model.Ipo) error {
	err := repo.DB.Where("id = ?", &ipo.ID).Updates(&ipo).Error
	if err != nil {
		log.Println("Error on update id", &ipo.ID, err)
		return nil
	}
	log.Println("Success on update id", &ipo.ID)
	return nil
}

func (repo *IpoRepo) createIPO(ipo *model.Ipo) error {
	err := repo.DB.Create(&ipo).Error
	if err != nil {
		log.Println("Error on create id", &ipo, err)
		return nil
	}
	log.Println("Success create data id", &ipo.ID)
	return nil
}

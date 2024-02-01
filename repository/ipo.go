package repository

import (
	"fmt"
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
	fmt.Println(ipo)
	err := repo.DB.Where("id = ?", &ipo.ID).First(&ipo).Error
	if err == nil {
		err = repo.DB.Where("id = ?", &ipo.ID).Updates(&ipo).Error
		if err != nil {
			log.Println("Error on update id", &ipo.ID, err)
			return nil
		}
		log.Println("Success on update id", &ipo.ID)
		return nil
	}

	err = repo.DB.Create(&ipo).Error
	if err != nil {
		log.Println("Error on create id", &ipo, err)
		return nil
	}
	log.Println("Success create data id", &ipo.ID)
	return nil
}

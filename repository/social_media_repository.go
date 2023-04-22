package repository

import (
	"finalMygram/model"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	Create(socialMedia model.SocialMedia) (model.SocialMedia, error)
	Delete(ID int) error
	FindAll() ([]model.SocialMedia, error)
	FindById(ID int) (model.SocialMedia, error)
	Update(socialMedia model.SocialMedia) (model.SocialMedia, error)
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{db}
}

func (r *socialMediaRepository) Create(socialMedia model.SocialMedia) (model.SocialMedia, error) {
	err := r.db.Create(&socialMedia).Error
	return socialMedia, err
}

func (r *socialMediaRepository) Delete(ID int) error {
	return r.db.Delete(&model.SocialMedia{}, ID).Error
}

func (r *socialMediaRepository) FindAll() ([]model.SocialMedia, error) {
	var socialMedias []model.SocialMedia
	err := r.db.Preload("User").Find(&socialMedias).Error
	return socialMedias, err
}

func (r *socialMediaRepository) FindById(ID int) (model.SocialMedia, error) {
	var socialMedia model.SocialMedia
	err := r.db.Preload("User").Find(&socialMedia).Error
	return socialMedia, err
}

func (r *socialMediaRepository) Update(socialMedia model.SocialMedia) (model.SocialMedia, error) {
	err := r.db.Save(&socialMedia).Error
	return socialMedia, err
}

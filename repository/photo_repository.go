package repository

import (
	"finalMygram/model"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(photo model.Photo) (model.Photo, error)
	Delete(ID int) error
	FindAll() ([]model.Photo, error)
	FindById(ID int) (model.Photo, error)
	Update(photo model.Photo) (model.Photo, error)
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db}
}

func (r *photoRepository) Create(photo model.Photo) (model.Photo, error) {
	err := r.db.Create(&photo).Error
	return photo, err
}

func (r *photoRepository) Delete(ID int) error {
	err := r.db.Delete(&model.Photo{}, ID).Error
	return err
}

func (r *photoRepository) FindAll() ([]model.Photo, error) {
	var photos []model.Photo
	err := r.db.Preload("User").Find(&photos).Error
	return photos, err
}

func (r *photoRepository) FindById(ID int) (model.Photo, error) {
	var photo model.Photo
	err := r.db.Preload("User").Find(&photo, ID).Error
	return photo, err
}

func (r *photoRepository) Update(photo model.Photo) (model.Photo, error) {
	err := r.db.Save(&photo).Error
	return photo, err
}

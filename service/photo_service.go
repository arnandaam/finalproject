package service

import (
	"finalMygram/model"
	"finalMygram/repository"
)

type PhotoService interface {
	Create(photo model.Photo) (model.Photo, error)
	Delete(ID int) (model.Photo, error)
	FindAll() ([]model.Photo, error)
	FindById(ID int) (model.Photo, error)
	Update(ID int, newPhoto model.PhotoUpdate) (model.Photo, error)
}

type photoService struct {
	photoRepository repository.PhotoRepository
}

func NewPhotoService(r repository.PhotoRepository) *photoService {
	return &photoService{r}
}

func (s *photoService) Create(photo model.Photo) (model.Photo, error) {
	return s.photoRepository.Create(photo)
}

func (s *photoService) Delete(ID int) (model.Photo, error) {
	photo, err := s.photoRepository.FindById(ID)
	err = s.photoRepository.Delete(ID)
	return photo, err
}

func (s *photoService) FindAll() ([]model.Photo, error) {
	return s.photoRepository.FindAll()
}

func (s *photoService) FindById(ID int) (model.Photo, error) {
	return s.photoRepository.FindById(ID)
}

func (s *photoService) Update(ID int, newPhoto model.PhotoUpdate) (model.Photo, error) {
	photo, err := s.photoRepository.FindById(ID)

	photo.Title = newPhoto.Title
	photo.Caption = newPhoto.Caption
	photo.PhotoURL = newPhoto.PhotoURL

	newestPhoto, err := s.photoRepository.Update(photo)

	return newestPhoto, err
}

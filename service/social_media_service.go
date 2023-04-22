package service

import (
	"finalMygram/model"
	"finalMygram/repository"
)

type SocialMediaService interface {
	Create(socialMedia model.SocialMedia) (model.SocialMedia, error)
	Delete(ID int) (model.SocialMedia, error)
	FindAll() ([]model.SocialMedia, error)
	FindById(ID int) (model.SocialMedia, error)
	Update(ID int, newSocialMedia model.SocialMediaUpdate) (model.SocialMedia, error)
}

type socialMediaService struct {
	socialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaService(r repository.SocialMediaRepository) *socialMediaService {
	return &socialMediaService{r}
}

func (s *socialMediaService) Create(socialMedia model.SocialMedia) (model.SocialMedia, error) {
	return s.socialMediaRepository.Create(socialMedia)
}

func (s *socialMediaService) Delete(ID int) (model.SocialMedia, error) {
	socialMedia, err := s.socialMediaRepository.FindById(ID)
	err = s.socialMediaRepository.Delete(ID)
	return socialMedia, err
}

func (s *socialMediaService) FindAll() ([]model.SocialMedia, error) {
	return s.socialMediaRepository.FindAll()
}

func (s *socialMediaService) FindById(ID int) (model.SocialMedia, error) {
	return s.socialMediaRepository.FindById(ID)
}

func (s *socialMediaService) Update(ID int, newSocialMedia model.SocialMediaUpdate) (model.SocialMedia, error) {
	socialMedia, err := s.socialMediaRepository.FindById(ID)

	socialMedia.Name = newSocialMedia.Name
	socialMedia.SocialMediaURL = newSocialMedia.SocialMediaURL

	newestSocialMedia, err := s.socialMediaRepository.Update(socialMedia)

	return newestSocialMedia, err
}

package repository

import (
	"finalMygram/model"

	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment model.Comment) (model.Comment, error)
	Delete(ID int) error
	FindAll() ([]model.Comment, error)
	FindById(ID int) (model.Comment, error)
	Update(comment model.Comment) (model.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) Create(comment model.Comment) (model.Comment, error) {
	err := r.db.Create(&comment).Error
	return comment, err
}

func (r *commentRepository) Delete(ID int) error {
	err := r.db.Delete(&model.Comment{}, ID).Error
	return err
}

func (r *commentRepository) FindAll() ([]model.Comment, error) {
	var comments []model.Comment
	err := r.db.Preload("User").Preload("Photo").Find(&comments).Error
	return comments, err
}

func (r *commentRepository) FindById(ID int) (model.Comment, error) {
	var comment model.Comment
	err := r.db.Preload("User").Preload("Photo").Find(&comment, ID).Error
	return comment, err
}

func (r *commentRepository) Update(comment model.Comment) (model.Comment, error) {
	err := r.db.Save(&comment).Error
	return comment, err
}

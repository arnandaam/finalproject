package service

import (
	"finalMygram/model"
	"finalMygram/repository"
)

type CommentService interface {
	Create(comment model.Comment) (model.Comment, error)
	Delete(ID int) (model.Comment, error)
	FindAll() ([]model.Comment, error)
	FindById(ID int) (model.Comment, error)
	Update(ID int, newComment model.CommentUpdate) (model.Comment, error)
}

type commentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(r repository.CommentRepository) *commentService {
	return &commentService{r}
}

func (s *commentService) Create(comment model.Comment) (model.Comment, error) {
	return s.commentRepository.Create(comment)
}

func (s *commentService) Delete(ID int) (model.Comment, error) {
	comment, err := s.commentRepository.FindById(ID)
	err = s.commentRepository.Delete(ID)
	return comment, err
}

func (s *commentService) FindAll() ([]model.Comment, error) {
	return s.commentRepository.FindAll()
}

func (s *commentService) FindById(ID int) (model.Comment, error) {
	return s.commentRepository.FindById(ID)
}

func (s *commentService) Update(ID int, newComment model.CommentUpdate) (model.Comment, error) {
	comment, err := s.commentRepository.FindById(ID)

	comment.Message = newComment.Message

	newestComment, err := s.commentRepository.Update(comment)

	return newestComment, err
}

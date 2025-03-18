package service

import (
	"fmt"
	"github.com/eCo13rus/comments_service/internal/models"
	"github.com/eCo13rus/comments_service/internal/repository"
)

// CommentService представляет сервис для работы с комментариями
type CommentService struct {
	repo repository.CommentRepository
}

// NewCommentService создаёт новый экземпляр CommentService
func NewCommentService(repo repository.CommentRepository) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

// AddComment добавляет новый комментарий
func (s *CommentService) AddComment(req *models.CommentRequest) (int, error) {
	if req.Content == "" {
		return 0, fmt.Errorf("контент комментария не может быть пустым")
	}

	id, err := s.repo.AddComment(req)
	if err != nil {
		return 0, fmt.Errorf("ошибка добавления комментария: %v", err)
	}

	return id, nil
}

// GetCommentsByNewsID возвращает комментарии для указанной новости
func (s *CommentService) GetCommentsByNewsID(newsID int) ([]models.Comment, error) {
	comments, err := s.repo.GetCommentsByNewsID(newsID)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения комментариев: %v", err)
	}

	return comments, nil
}

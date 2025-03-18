package repository

import (
	"database/sql"
	"fmt"
	"github.com/eCo13rus/comments_service/internal/models"
	_ "github.com/lib/pq"
	_ "time"
)

// CommentRepository представляет интерфейс для работы с репозиторием комментариев
type CommentRepository interface {
	AddComment(comment *models.CommentRequest) (int, error)
	GetCommentsByNewsID(newsID int) ([]models.Comment, error)
	Close() error
}

// PostgresRepository реализует интерфейс CommentRepository для PostgreSQL
type PostgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository создаёт новый экземпляр PostgresRepository
func NewPostgresRepository(cfg *models.DatabaseConfig) (*PostgresRepository, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к БД: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка проверки соединения с БД: %v", err)
	}

	return &PostgresRepository{db: db}, nil
}

// AddComment добавляет новый комментарий в БД
func (r *PostgresRepository) AddComment(comment *models.CommentRequest) (int, error) {
	var id int
	query := `
        INSERT INTO comments (news_id, parent_id, content)
        VALUES ($1, $2, $3)
        RETURNING id
    `

	err := r.db.QueryRow(query, comment.NewsID, comment.ParentID, comment.Content).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("ошибка добавления комментария: %v", err)
	}

	return id, nil
}

// GetCommentsByNewsID возвращает комментарии для указанной новости
func (r *PostgresRepository) GetCommentsByNewsID(newsID int) ([]models.Comment, error) {
	query := `
        SELECT id, news_id, parent_id, content, created_at, updated_at
        FROM comments
        WHERE news_id = $1
        ORDER BY created_at ASC
    `

	rows, err := r.db.Query(query, newsID)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения комментариев: %v", err)
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.NewsID,
			&comment.ParentID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("ошибка сканирования результата: %v", err)
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка итерации по результатам: %v", err)
	}

	return comments, nil
}

func (r *PostgresRepository) Close() error {
	return r.db.Close()
}

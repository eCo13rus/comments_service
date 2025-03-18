package models

import "time"

// Comment представляет модель комментария в системе
type Comment struct {
	ID        int       `json:"id" db:"id"`
	NewsID    int       `json:"news_id" db:"news_id"`
	ParentID  *int      `json:"parent_id,omitempty" db:"parent_id"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// CommentRequest представляет запрос на создание нового комментария
type CommentRequest struct {
	NewsID   int    `json:"news_id"`
	ParentID *int   `json:"parent_id,omitempty"`
	Content  string `json:"content"`
}

// CommentResponse представляет ответ при запросе комментариев
type CommentResponse struct {
	Comments []Comment `json:"comments"`
}

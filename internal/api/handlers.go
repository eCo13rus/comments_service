package api

import (
	"encoding/json"
	"github.com/eCo13rus/comments_service/internal/models"
	"github.com/eCo13rus/comments_service/internal/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// Handler представляет обработчик HTTP-запросов
type Handler struct {
	commentService *service.CommentService
}

// NewHandler создаёт новый экземпляр Handler
func NewHandler(commentService *service.CommentService) *Handler {
	return &Handler{
		commentService: commentService,
	}
}

// AddComment обрабатывает запрос на добавление комментария
func (h *Handler) AddComment(w http.ResponseWriter, r *http.Request) {
	requestID := r.Context().Value(RequestIDKey).(string)

	var req models.CommentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("[%s] Ошибка декодирования запроса: %v", requestID, err)
		http.Error(w, "Некорректный формат запроса", http.StatusBadRequest)
		return
	}

	id, err := h.commentService.AddComment(&req)
	if err != nil {
		log.Printf("[%s] Ошибка добавления комментария: %v", requestID, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"id":      id,
		"status":  "success",
		"message": "Комментарий успешно добавлен",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("[%s] Ошибка кодирования ответа: %v", requestID, err)
	}
}

// GetComments обрабатывает запрос на получение комментариев к новости
func (h *Handler) GetComments(w http.ResponseWriter, r *http.Request) {
	requestID := r.Context().Value(RequestIDKey).(string)

	vars := mux.Vars(r)
	newsIDStr := vars["news_id"]
	newsID, err := strconv.Atoi(newsIDStr)
	if err != nil {
		log.Printf("[%s] Некорректный ID новости: %v", requestID, err)
		http.Error(w, "Некорректный ID новости", http.StatusBadRequest)
		return
	}

	comments, err := h.commentService.GetCommentsByNewsID(newsID)
	if err != nil {
		log.Printf("[%s] Ошибка получения комментариев: %v", requestID, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.CommentResponse{
		Comments: comments,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("[%s] Ошибка сериализации ответа: %v", requestID, err)
	}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{
		"status": "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/eCo13rus/comments_service/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/eCo13rus/comments_service/internal/api"
	"github.com/eCo13rus/comments_service/internal/models"
	"github.com/eCo13rus/comments_service/internal/repository"
)

func main() {
	config, err := loadConfig("configs/config.json")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	repo, err := repository.NewPostgresRepository(&config.Database)
	if err != nil {
		log.Fatalf("Ошибка инициализации репозитория: %v", err)
	}
	defer repo.Close()

	commentService := service.NewCommentService(repo)

	handler := api.NewHandler(commentService)

	addr := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	server := api.NewServer(handler, addr)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	log.Println("Сервис комментариев запущен")

	<-stop

	log.Println("Сервис комментариев останавливается...")
}

func loadConfig(path string) (*models.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла конфигурации: %v", err)
	}
	defer file.Close()

	var config models.Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("ошибка декодирования конфигурации: %v", err)
	}

	return &config, nil
}

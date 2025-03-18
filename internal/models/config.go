package models

// Config представляет конфигурацию приложения
type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
	Censor   CensorConfig   `json:"censor"`
}

// ServerConfig представляет настройки HTTP-сервера
type ServerConfig struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

// DatabaseConfig представляет настройки базы данных
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	SSLMode  string `json:"ssl_mode"`
}

// CensorConfig представляет настройки для модуля цензуры
type CensorConfig struct {
	ForbiddenWords []string `json:"forbidden_words"`
	CheckInterval  int      `json:"check_interval"`
	BatchSize      int      `json:"batch_size"`
}

package main

import (
	"fmt"
	"log"

	"github.com/Danyarbrg/securelogAPI/internal/config"
	"github.com/Danyarbrg/securelogAPI/internal/db"
	"github.com/Danyarbrg/securelogAPI/internal/models"
	"github.com/Danyarbrg/securelogAPI/internal/routes"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println("PORT:", cfg.Port)
	fmt.Println("DATABASE_URL:", cfg.DatabaseURL)

	database := db.Connect(cfg.DatabaseURL)
	defer database.Close()

	entry := models.LogEntry {
		EventType: "USER_LOGIN",
		UserID: "user-1234",
		Payload: `{"ip":"127.0.0.1"}`,
	}
	
	if err := db.InsertLog(database, entry); err != nil {
		log.Fatalf("Misstake while inserting log: %v", err)
	}

	router := routes.SetupRouter(database)

	log.Printf("🚀 Сервер запущен на порту %s\n", cfg.Port)
	err := router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

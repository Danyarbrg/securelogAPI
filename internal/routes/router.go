package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/Danyarbrg/securelogAPI/internal/handlers"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	logHandler := handlers.NewLogHandler(db)

	router.POST("/logs", logHandler.CreateLog)

	return router
}

package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Danyarbrg/securelogAPI/internal/db"
	"github.com/Danyarbrg/securelogAPI/internal/models"
)

type LogHandler struct {
	DB *sql.DB
}

func NewLogHandler(db *sql.DB) *LogHandler {
	return &LogHandler{DB: db}
}

func (h *LogHandler) CreateLog(c *gin.Context) {
	var input models.LogEntry

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong format JSON"})
		return 
	}

	if input.EventType == "" || input.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "event_type and user_id must be"})
		return
	}

	err := db.InsertLog(h.DB, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Misstake in inserting the log"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Log added"})
}

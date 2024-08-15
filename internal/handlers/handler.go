package handlers

import (
    "database/sql"
    "net/http"

    "github.com/gin-gonic/gin"
)

// Handler struct to hold the database connection
type Handler struct {
    DB *sql.DB
}

// NewHandler returns a new handler with the database connection
func NewHandler(db *sql.DB) *Handler {
    return &Handler{DB: db}
}

// PingHandler is an example handler that uses the database connection
func (h *Handler) PingHandler(c *gin.Context) {
    // Example database interaction
    var result string
    err := h.DB.QueryRow("SELECT 'pong'").Scan(&result)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": result})
}


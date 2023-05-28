package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Soundex(string) ([]string, error)
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{service: service}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/soundex", h.soundex)

	return router
}

func (h *Handler) soundex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

package handlers

import (
	"net/http"

	"github.com/artemKapitonov/soundex/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Service interface {
	Soundex(names models.Names) []string
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{service: service}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/soundex", h.soundex)

	return router
}

func (h *Handler) soundex(c *gin.Context) {
	var input models.Names

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input value")
	}

	var soundexes = h.service.Soundex(input)

	c.JSON(http.StatusOK, models.SoundexResponse{
		Soundexes: soundexes,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)

	c.AbortWithStatusJSON(statusCode, ErrorResponse{Message: message})
}

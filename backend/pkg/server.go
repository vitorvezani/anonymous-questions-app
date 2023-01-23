package pkg

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type Server struct {
	gin *gin.Engine
}

func NewServer(gin *gin.Engine, h *Handler) (*Server, error) {
	if gin == nil {
		return nil, errors.New("gin is require")
	}
	if h == nil {
		return nil, errors.New("h is require")
	}
	setupRoutes(gin, h)
	return &Server{gin}, nil
}

func (s Server) Start() error {
	return s.gin.Run() // listen and serve on 0.0.0.0:8080
}

func (s Server) Stop() {
	// do nothing
}

func setupRoutes(gin *gin.Engine, h *Handler) {
	v0 := gin.Group("/api/v0")

	v0.GET("/questions", h.listQuestions)
	v0.POST("/questions", h.addQuestion)
	v0.DELETE("/questions", h.deleteQuestions)
	v0.POST("/questions/:id/up-vote", h.upVoteQuestion)
}

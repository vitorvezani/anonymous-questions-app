package pkg

import (
	"errors"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

	gin.Use(cors.Default())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("questionValidator", questionValidatorFn)
	}

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

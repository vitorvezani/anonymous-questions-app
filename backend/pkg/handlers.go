package pkg

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) (*Handler, error) {
	if db == nil {
		return nil, errors.New("db is require")
	}
	return &Handler{db}, nil
}

func (h Handler) listQuestions(c *gin.Context) {
	var questions []Question
	err := h.db.Find(&questions).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, questions)
}

func (h Handler) addQuestion(c *gin.Context) {
	// TODO
}

func (h Handler) deleteQuestions(c *gin.Context) {
	// TODO
}

func (h Handler) upVoteQuestion(c *gin.Context) {
	// TODO
}

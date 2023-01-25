package pkg

import (
	"errors"
	"net/http"
	"strconv"

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
	var question Question
	err := c.ShouldBindJSON(&question)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.db.Create(&question).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, question)
}

func (h Handler) deleteQuestions(c *gin.Context) {
	err := h.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Question{}).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent)
}

func (h Handler) upVoteQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var upVotes int
	err = h.db.Raw("UPDATE questions SET up_votes = up_votes + 1 WHERE id = ? RETURNING up_votes", id).Scan(&upVotes).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"up_votes": upVotes})
}

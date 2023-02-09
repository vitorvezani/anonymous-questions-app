package pkg

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type Handler struct {
	tracer trace.Tracer
	db     *gorm.DB
}

func NewHandler(t trace.Tracer, db *gorm.DB) (*Handler, error) {
	if db == nil {
		return nil, errors.New("db is require")
	}
	if t == nil {
		return nil, errors.New("tracer is require")
	}
	return &Handler{t, db}, nil
}

func (h Handler) listQuestions(c *gin.Context) {
	var questions []Question
	err := h.db.WithContext(c.Request.Context()).Find(&questions).Error
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

	err = h.db.WithContext(c.Request.Context()).Create(&question).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = integrateWithExternalService(c, h.tracer)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, question)
}

func integrateWithExternalService(c *gin.Context, t trace.Tracer) error {
	_, span := t.Start(c.Request.Context(), "XYZ integration")
	defer span.End()

	time.Sleep(100 * time.Millisecond)

	return nil
}

func (h Handler) deleteQuestions(c *gin.Context) {
	err := h.db.WithContext(c.Request.Context()).Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Question{}).Error
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
	err = h.db.WithContext(c.Request.Context()).Raw("UPDATE questions SET up_votes = up_votes + 1 WHERE id = ? RETURNING up_votes", id).Scan(&upVotes).Error
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"up_votes": upVotes})
}

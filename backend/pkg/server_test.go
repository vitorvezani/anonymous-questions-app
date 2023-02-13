package pkg_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"fairwinds.com/anonymous-questions-app/pkg"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var dbFile = "sqlite_test.db"

func TestListQuestions(t *testing.T) {
	r := setupServer(t)
	defer cleanUp()

	req, _ := http.NewRequest("GET", "/api/v0/questions", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

func TestAddQuestionsError(t *testing.T) {
	r := setupServer(t)
	defer cleanUp()

	payload := strings.NewReader(`{"Text": "Is this a good question"}`)

	req, _ := http.NewRequest("POST", "/api/v0/questions", payload)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"Key: 'Question.Text' Error:Field validation for 'Text' failed on the 'questionValidator' tag"}`, w.Body.String())
}

func TestAddQuestions(t *testing.T) {
	// TODO
}

func TestDeleteQuestions(t *testing.T) {
	// TODO
}

func TestUpVoteQuestion(t *testing.T) {
	// TODO
}

func addQuestion(t *testing.T, r *gin.Engine) {
	payload := strings.NewReader(`{"Text": "Is this a good question?"}`)

	req, _ := http.NewRequest("POST", "/api/v0/questions", payload)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func setupServer(t *testing.T) *gin.Engine {
	r := gin.Default()

	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&pkg.Question{})

	h, err := pkg.NewHandler(db)
	assert.NoError(t, err)

	_, err = pkg.NewServer(r, h)
	assert.NoError(t, err)

	return r
}

func cleanUp() {
	err := os.Remove(dbFile)
	if err != nil {
		panic(err)
	}
}

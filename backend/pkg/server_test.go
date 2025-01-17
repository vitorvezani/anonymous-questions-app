package pkg_test

import (
	"encoding/json"
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
	r := setupServer(t)
	defer cleanUp()

	payload := strings.NewReader(`{"Text": "Is this a good question?"}`)

	req, _ := http.NewRequest("POST", "/api/v0/questions", payload)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var q pkg.Question
	err := json.Unmarshal(w.Body.Bytes(), &q)
	assert.NoError(t, err)
	assert.NotZero(t, q.ID)
	assert.Zero(t, q.UpVotes)
	assert.Equal(t, "Is this a good question?", q.Text)
}

func TestDeleteQuestions(t *testing.T) {
	r := setupServer(t)
	defer cleanUp()

	addQuestion(t, r)

	req, _ := http.NewRequest("DELETE", "/api/v0/questions", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestUpVoteQuestion(t *testing.T) {
	r := setupServer(t)
	defer cleanUp()

	addQuestion(t, r)

	req, _ := http.NewRequest("POST", "/api/v0/questions/1/up-vote", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"up_votes":1}`, w.Body.String())
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

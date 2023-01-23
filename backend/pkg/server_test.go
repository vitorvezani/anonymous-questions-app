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

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v0/questions", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

func TestAddQuestions(t *testing.T) {
	r := setupServer(t)
	defer cleanUp()

	w := httptest.NewRecorder()

	payload := strings.NewReader(`{"Text": "Is this a good question?"}`)

	req, _ := http.NewRequest("POST", "/api/v0/questions", payload)
	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	var q pkg.Question
	err := json.Unmarshal(w.Body.Bytes(), &q)
	assert.NoError(t, err)
	assert.NotZero(t, q.ID)
	assert.Zero(t, q.UpVotes)
	assert.Equal(t, "Is this a good question?", q.Text)
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

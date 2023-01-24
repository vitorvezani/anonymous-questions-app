package pkg_test

import (
	"os"
	"testing"

	"fairwinds.com/anonymous-questions-app/pkg"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var dbFile = "sqlite_test.db"

func TestListQuestions(t *testing.T) {
	// TODO
}

func TestAddQuestions(t *testing.T) {
	// TODO
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

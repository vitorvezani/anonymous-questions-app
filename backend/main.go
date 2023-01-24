package main

import (
	"fairwinds.com/anonymous-questions-app/pkg"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		logrus.Fatal("could not open db connection", err)
	}

	db.AutoMigrate(&pkg.Question{})

	h, err := pkg.NewHandler(db)
	if err != nil {
		logrus.Fatal("could not create handler", err)
	}

	r := gin.Default()

	s, err := pkg.NewServer(r, h)
	if err != nil {
		logrus.Fatal("could not create server", err)
	}

	s.Start()
}

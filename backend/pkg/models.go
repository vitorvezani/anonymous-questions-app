package pkg

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model

	Text    string `binding:"required"`
	UpVotes int
}

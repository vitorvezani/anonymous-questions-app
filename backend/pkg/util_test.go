package pkg

import (
	"testing"
)

// go test -run ^TestValidateQuestion$ fairwinds.com/anonymous-questions-app/pkg -v
func TestValidateQuestion(t *testing.T) {
	var tests = []struct {
		i    string
		want bool
	}{
		{"is this a question", false},
		{"?is this a question", false},
		{"is this a question¿", false},
		{"is this a question?", true},
	}

	for _, tt := range tests {
		t.Run(tt.i, func(t *testing.T) {
			ans := validateQuestion(tt.i)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

package pkg

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var questionValidatorFn validator.Func = func(fl validator.FieldLevel) bool {
	question, ok := fl.Field().Interface().(string)
	if ok {
		return validateQuestion(question)
	}
	return true
}

func validateQuestion(q string) bool {
	return strings.HasSuffix(q, "?")
}

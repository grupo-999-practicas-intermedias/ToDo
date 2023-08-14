package routes

import (
	"server/api-go-test/model"

	"github.com/go-playground/validator/v10"
)

var valdiate = validator.New()

func ValidateStructTodo(todo model.Todo) []*model.ErrorResponse {
	var errors []*model.ErrorResponse
	err := valdiate.Struct(todo)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element model.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

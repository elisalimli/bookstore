package utils

import "github.com/go-playground/validator"

var validate = validator.New()

type ErrorResponse struct {
	Field   string `json:"field"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
}

var customErrorMessages = map[string]string{
	"required": "This field cannot be empty",
}

func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.ActualTag()
			element.Tag = err.Tag()
			element.Message = customErrorMessages[element.Tag]

			errors = append(errors, &element)
		}
	}
	return errors
}

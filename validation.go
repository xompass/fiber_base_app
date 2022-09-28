package fiber_base_app

import (
	"github.com/go-playground/validator/v10"
)

var validatorInstance = validator.New()

type validationResult struct {
	Field string      `json:"field"`
	Tag   string      `json:"tag"`
	Value interface{} `json:"value"`
}

func ValidateStruct(_struct interface{}) error {
	err := validatorInstance.Struct(_struct)
	if err != nil {
		responseError := CustomHTTPError{
			Code: 400,
		}
		for _, e := range err.(validator.ValidationErrors) {
			responseError.Details = append(responseError.Details, validationResult{
				Field: e.Field(),
				Tag:   e.Tag(),
				Value: e.Value(),
			})
		}

		return responseError
	}

	return nil
}

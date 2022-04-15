package modelValidation

import (
	"github.com/go-playground/validator"
	"wb/app/wb/logger"
	"wb/app/wb/storage"
)

func Validate(model *storage.ModelJSON) bool {
	v := validator.New()

	err := v.Struct(model)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			logger.Log.Error(e)
		}
		return true
	}
	return false
}
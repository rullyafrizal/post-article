package validation

import "github.com/go-playground/validator/v10"

type Validator interface {
	Validate(data any) []error
}

type ValidatorImpl struct {
	validate *validator.Validate
}

func NewValidator() *ValidatorImpl {
	v := &ValidatorImpl{validate: validator.New()}
	return v
}

func (v *ValidatorImpl) Validate(data any) []error {
	err := v.validate.Struct(data)
	if err != nil {
		return MakeErrors(err.(validator.ValidationErrors))
	}
	return nil
}

package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func MakeErrors(errs validator.ValidationErrors) []error {
	var vErrs []error
	for _, err := range errs {
		b := &errorBuilder{}
		e := b.make(err)
		vErrs = append(vErrs, e)
	}
	return vErrs
}

type errorBuilder struct {
	field string
	verb  string
	tag   string
}

func (b *errorBuilder) make(err validator.FieldError) error {
	b.createValidationError(err.Field(), err.Tag())
	validationError := fmt.Errorf("%v %v %v", b.field, b.verb, b.tag)
	return validationError
}

func (b *errorBuilder) createValidationError(field, tag string) {
	switch tag {
	case "required":
		b.fieldIsRequired(field)
	case "min":
		b.fieldIsLessThanMinimumValue(field)
	case "max":
		b.fieldIsExceedsTheMaxValue(field)
	case "date", "businessName", "notOnlySpace", "oneof":
		b.fieldIsInvalid(field)
	case "initialBalance":
		b.fieldMustSatisfyValidInitialBalance(field)
	case "nefield":
		b.fieldCannotBeTheSame()
	default:
		b.fieldMustSatisfyTag(field, tag)
	}
}

func (b *errorBuilder) fieldIsRequired(field string) {
	b.field = field
	b.verb = "is"
	b.tag = "required"
}

func (b *errorBuilder) fieldIsLessThanMinimumValue(field string) {
	b.field = field
	b.verb = "is"
	b.tag = "less than minimum length / value"
}

func (b *errorBuilder) fieldIsExceedsTheMaxValue(field string) {
	b.field = field
	b.verb = "is"
	b.tag = "exceeds the maximum length / value"
}

func (b *errorBuilder) fieldMustSatisfyTag(field, tag string) {
	b.field = field
	b.verb = "must be"
	b.tag = tag
}

func (b *errorBuilder) fieldIsInvalid(field string) {
	b.field = field
	b.verb = "is"
	b.tag = "invalid"
}

func (b *errorBuilder) fieldMustSatisfyValidInitialBalance(field string) {
	b.field = field
	b.verb = "value must be"
	b.tag = "between 0 and 13 digits nominal"
}

func (b *errorBuilder) fieldCannotBeTheSame() {
	b.field = "source id and target id"
	b.verb = "cannot be"
	b.tag = "the same"
}

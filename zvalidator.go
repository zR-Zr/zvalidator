package zvalidator

import (
	"errors"
)

type Validator func(value any, rawData map[string]any, rule Rule) bool

var validators = map[string]Validator{
	"required": requiredValidator,
	"min":      minValidator,
	"max":      maxValidator,
	"range":    rangeValidator,
}

func Validate(data map[string]any, rules Rules) (bool, map[string]string) {
	validationErrors := make(map[string]string)

	for field, fieldRules := range rules {
		fieldValue, ok := getFieldValue(data, field)
		if !ok {
			fieldValue = nil
		}

		for _, rule := range fieldRules {

			if rule.Type != "required" && (fieldValue == nil || isEmptyValue(fieldValue)) {
				continue
			}

			isValid := true

			if rule.CustomValidator != nil {
				if !rule.CustomValidator(fieldValue, data) {
					validationErrors[field] = rule.Message
					isValid = false
				}
			} else {
				if registedValidator, ok := validators[rule.Type]; ok {
					if !registedValidator(fieldValue, data, rule) {
						validationErrors[field] = rule.Message
						isValid = false
					}
				} else {

					panic(errors.New("unknown validator type: " + rule.Type))
				}
			}

			if !isValid {
				break
			}
		}
	}

	return len(validationErrors) == 0, validationErrors
}

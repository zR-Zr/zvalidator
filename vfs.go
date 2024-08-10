package zvalidator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func isRequired() ValidatorFunc {
	return func(value any, extra ...any) bool {
		if value == nil {
			return false
		}

		strValue := strings.TrimSpace(fmt.Sprintf("%v", value))

		return strValue != ""
	}
}

func min(value any, extra ...any) bool {
	if len(extra) != 1 {
		panic("min validator need 1 extra arguments")
	}

	extraValue, ok := extra[0].(int)
	if !ok {
		panic("min validator extra arguments must be int")
	}

	switch typeValue := value.(type) {
	case string:
		return len(typeValue) >= extraValue
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(typeValue).Int() >= int64(extraValue)
	default:
		return false
	}

}

func max(value any, extra ...any) bool {
	if len(extra) != 1 {
		panic("max validator need 1 extra arguments")
	}

	extraValue, ok := extra[0].(int)
	if !ok {
		panic("max validator extra arguments must be int")
	}

	switch typeValue := value.(type) {
	case string:
		return len(typeValue) <= extraValue
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(typeValue).Int() <= int64(extraValue)
	default:
		return false
	}
}

func IsNumeric(value any, extra ...any) bool {
	if value == nil {
		return false
	}

	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	case string:
		_, err := strconv.ParseFloat(value.(string), 64)
		return err == nil

	default:
		return false
	}
}

func isEmail(value any, extra ...any) bool {
	if value == nil {
		return false
	}

	email, ok := value.(string)
	if !ok {
		return false
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)

	return re.MatchString(email)
}

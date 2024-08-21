package zvalidator

import (
	"regexp"

	"github.com/zR-Zr/zvalidator.git/optional"
)

type Range struct {
	Min float64
	Max float64
}

func Rangef(min, max float64) *Range {
	return &Range{Min: min, Max: max}
}

type CustomValidator func(data any, rawData map[string]any) bool

type Rule struct {
	Type            string // 验证规则类型
	Message         string // 自定义错误信息
	Min             optional.Optional[int]
	Max             optional.Optional[int]
	Range           *Range
	Pattern         regexp.Regexp
	Email           bool
	In              []any
	Criteria        any
	CustomValidator CustomValidator
}

type Rules map[string][]Rule

func RequiredRule(message string) Rule {
	return Rule{Type: "required", Message: message}
}

func MinRule(min int, message string) Rule {
	return Rule{Type: "min", Message: message, Min: optional.New(min)}
}

func MaxRule(max int, message string) Rule {
	return Rule{Type: "max", Message: message, Max: optional.New(max)}
}

func RangeRule(min, max int, message string) Rule {
	return Rule{Type: "range", Message: message, Range: Rangef(float64(min), float64(max))}
}

// func isEmptyValue(v reflect.Value) bool {
// 	return v.IsZero()
// }

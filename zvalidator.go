package zvalidator

import (
	"fmt"
	"reflect"
)

// 存储自定义验证函数
var validatorContainer = map[string]ValidatorFunc{
	"required": isRequired(),
	"min":      min,
	"max":      max,
	"numeric":  IsNumeric,
	"email":    isEmail,
}

func RegisterValidator(typeName string, validatorFunc ValidatorFunc) {
	validatorContainer[typeName] = validatorFunc
}

func Validate(data map[string]any, rules Rules) (bool, map[string]string) {
	errors := make(map[string]string)

	for field, fieldRules := range rules {
		// 使用 rules 的key ,到 data 中取值
		fieldValue, ok := data[field]
		if !ok {
			fieldValue = nil
		}

		for _, rule := range fieldRules { // Rule 数组
			// 如果该 字段不是必填, 且没有值则跳过, (处理,不是必填项,但是如果填入,就要验证格式的情况)
			// 例如 email 非必填,但是如果填了,就要复合email的格式
			if rule.Type != "required" && (fieldValue == nil || isEmptyValue(reflect.ValueOf(fieldValue))) {
				continue
			}

			isValid := true

			// 1. 优先使用 rule.Validators 中自定义的验证函数
			if len(rule.Validators) > 0 {
				for _, ruleInnerValidator := range rule.Validators {
					// 验证失败
					if !ruleInnerValidator(fieldValue, rule.Value) {
						errors[field] = rule.Message
						isValid = false // 设置验证失败标志
						break           // 如果一个验证眼熟失败,计算验证失败,不继续验证
					}
				}
			} else {
				// 如果 ruoe.Validators 为空, 则使用 validatorContainer 中 注册的验证函数
				if registedValidator, ok := validatorContainer[rule.Type]; ok {
					if !registedValidator(fieldValue, rule.Value) {
						errors[field] = rule.Message
						isValid = false
					}
				} else {
					panic(fmt.Sprintf("不支持的验证类型: %s", rule.Type))
				}
			}
			if !isValid {
				break // 当前规则验证失败,跳到下一个规则
			}
		}
	}

	return len(errors) == 0, errors
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String, reflect.Chan:
		return v.Len() == 0
	default:
		return v.IsZero()
	}
}

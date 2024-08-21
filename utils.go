package zvalidator

import (
	"reflect"
	"strings"
)

func isEmptyValue(v any) bool {
	if v == nil {
		return true
	}

	rValue := reflect.ValueOf(v)

	switch rValue.Kind() {
	case reflect.Interface, reflect.Ptr:
		return rValue.IsNil()
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String, reflect.Chan:
		return rValue.Len() == 0
	default:
		return rValue.IsZero()
	}
}

// 根据字段路径获取map中的值, 例如 fieldPath : "user.name",  {"user": {"name": "John"}} => "John"
func getFieldValue(data map[string]any, fieldPath string) (any, bool) {
	parts := strings.Split(fieldPath, ".")
	current := data

	for _, part := range parts {
		if value, exists := current[part]; exists {
			if next, ok := value.(map[string]any); ok {
				current = next
			} else {
				return value, true // 找到目标字段的值
			}
		} else {
			return nil, false // 字段路径不存在
		}
	}
	return current, true // 返回最后一个嵌套的结构体
}

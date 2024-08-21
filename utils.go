package zvalidator

import (
	"fmt"
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

func structToMap(obj any) (map[string]any, error) {
	data := make(map[string]any)
	v := reflect.ValueOf(obj)

	// 如果是指针类型, 则获取其指向的值
	if v.Kind() == reflect.Ptr {
		if v.IsNil() { // 处理空指针的情况
			return nil, fmt.Errorf("input cannot be a nil pointer")
		}
		v = v.Elem()
	}

	// 必须是结构体类型
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input must be a struct or a pointer to a struct")
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Tag.Get("json")
		if fieldName == "" {
			fieldName = field.Name // 如果没有json tag, 则使用字段名
		}

		fieldValue := v.Field(i)

		// 处理嵌套结构体
		if fieldValue.Kind() == reflect.Struct {
			nestedMap, err := structToMap(fieldValue.Interface())
			if err != nil {
				return nil, err
			}
			data[fieldName] = nestedMap
		} else if fieldValue.Kind() == reflect.Ptr {
			if !fieldValue.IsNil() {
				nestedElm := fieldValue.Elem()
				if nestedElm.Kind() == reflect.Struct {
					nestedMap, err := structToMap(nestedElm.Interface())
					if err != nil {
						return nil, err
					}
					data[fieldName] = nestedMap
				} else {
					data[fieldName] = fieldValue.Elem().Interface()
				}
			} else {
				data[fieldName] = nil
			}
		} else {
			data[fieldName] = v.Field(i).Interface()
		}

	}

	return data, nil
}

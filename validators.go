package zvalidator

import (
	"reflect"
)

func requiredValidator(value any, rawData map[string]any, rule Rule) bool {
	return !isEmptyValue(reflect.ValueOf(value))
}

func minValidator(value any, rawData map[string]any, rule Rule) bool {
	if !rule.Min.IsSet() {
		panic("min validator need Min ")
	}

	switch typeValue := value.(type) {
	case string:
		return len(typeValue) >= rule.Min.Value()
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(typeValue).Int() >= int64(rule.Min.Value())
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(typeValue).Uint() >= uint64(rule.Min.Value())
	case float32, float64:
		return reflect.ValueOf(typeValue).Float() >= float64(rule.Min.Value())
	default:
		return false
	}
}

func maxValidator(value any, rawData map[string]any, rule Rule) bool {
	if !rule.Max.IsSet() {
		panic("max validator need Max ")
	}

	switch typeValue := value.(type) {
	case string:
		return len(typeValue) <= rule.Max.Value()
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(typeValue).Int() <= int64(rule.Max.Value())
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(typeValue).Uint() <= uint64(rule.Max.Value())
	case float32, float64:
		return reflect.ValueOf(typeValue).Float() <= float64(rule.Max.Value())
	default:
		return false
	}
}

func rangeValidator(value any, rawData map[string]any, rule Rule) bool {
	if rule.Range == nil {
		panic("range validator need Range ")
	}

	min := rule.Range.Min
	max := rule.Range.Max

	switch typeValue := value.(type) {
	case string:
		return float64(len(typeValue)) >= float64(min) && float64(len(typeValue)) <= float64(max)
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(typeValue).Int() >= int64(min) && reflect.ValueOf(typeValue).Int() <= int64(max)
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(typeValue).Uint() >= uint64(min) && reflect.ValueOf(typeValue).Uint() <= uint64(max)
	case float32, float64:
		return reflect.ValueOf(typeValue).Float() >= float64(min) && reflect.ValueOf(typeValue).Float() <= float64(max)
	default:
		return false
	}
}

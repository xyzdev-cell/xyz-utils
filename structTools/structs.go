package xyz_struct

import (
	"fmt"
	"reflect"
)

// 把结构体反射为map, 以Tag为key
func StructToMap(in any, tagName string) (map[string]any, error) {
	out := make(map[string]any)

	reflectValue := reflect.ValueOf(in)
	if reflectValue.Kind() == reflect.Ptr { // 指针时需要取为 Elem()
		reflectValue = reflectValue.Elem()
	}

	if reflectValue.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", reflectValue)
	}

	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < reflectValue.NumField(); i++ {
		if tagValue := reflectValue.Type().Field(i).Tag.Get(tagName); tagValue != "" {
			out[tagValue] = reflectValue.Field(i).Interface()
		}
	}
	return out, nil
}

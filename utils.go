package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime/pprof"
	"time"
)

func PrettyJSON(jsonByte []byte) string {
	prettyJSON := &bytes.Buffer{}
	json.Indent(prettyJSON, jsonByte, "", "    ")
	return prettyJSON.String()
}

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

func CsharpStringHashV1(str string) int64 {
	var num1, num2 int32 = 5381, 5381

	length := len(str)
	for i, nexti := 0, 0; i < length; i += 2 {
		num1 = (num1 << 5) + num1 ^ int32(str[i])
		nexti = i + 1
		if nexti != length {
			num2 = (num2 << 5) + num2 ^ int32(str[nexti])
		}
	}

	return int64(num1+num2*1566083941) & 0xFFFFFFFF
}

func CsharpStringHashV2(str string) int64 {
	var num1 int32 = 352654597
	var num2 int32 = num1

	var length int
	for length = len(str); length > 2; length -= 4 {
		num1 = (num1 << 5) + num1 + (num1 >> 27) ^ int32(str[0])
		num2 = (num2 << 5) + num2 + (num2 >> 27) ^ int32(str[1])
		str = str[2:]
	}
	if length > 0 {
		num1 = (num1 << 5) + num1 + (num1 >> 27) ^ int32(str[0])
	}
	return int64(num1+num2*1566083941) & 0xFFFFFFFF
}

func StartCPUProfile(file string, duration time.Duration) (err error) {
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	log.Println("Start cpu profiling for", duration)
	err = pprof.StartCPUProfile(f)
	if err != nil {
		f.Close()
		return err
	}

	time.AfterFunc(duration, func() {
		pprof.StopCPUProfile()
		f.Close()
		log.Println("Stop CPU profiling after", duration)
	})
	return nil
}

func StartMemoryProfile(file string, duration time.Duration) (err error) {
	f, err := os.Create(file)
	if err != nil {
		return err
	}

	log.Println("Start memory profiling for", duration)
	time.AfterFunc(duration, func() {
		err = pprof.WriteHeapProfile(f)
		if err != nil {
			log.Println(err)
		}
		f.Close()
		log.Println("Stop memory profiling after", duration)
	})
	return nil
}

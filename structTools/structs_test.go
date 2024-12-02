package xyz_struct

import (
	"fmt"
	"testing"
)

type testTypeUser struct {
	Csv []string `toml:"Csv"`
}

func TestStructToMap(t *testing.T) {
	resMap, err := StructToMap(&testTypeUser{
		Csv: []string{
			"123",
			"456",
		},
	}, "toml")
	if err != nil {
		t.Errorf("StructToMap error:%s", err)
	} else {
		fmt.Println("res map:", resMap)
	}
}

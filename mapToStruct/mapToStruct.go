package main

import (
	"fmt"
	"reflect"
)

func map2Struct(m map[string]string, t interface{}) {
	if t == nil {
		return
	}
	rType := reflect.TypeOf(t).Elem()
	rValue := reflect.ValueOf(t).Elem()
	for k, v := range m {
		for i := 0; i < rValue.NumField(); i++ {
			tag := rType.Field(i).Tag.Get("json")
			if tag == k {
				rValue.Field(i).Set(reflect.ValueOf(v))
			}
		}
	}
}
func main() {
	t := struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	}{}
	m := map[string]string{
		"id": "11233cc",
	}
	map2Struct(m, &t)
	fmt.Printf("%v", t)
}

package utils

import (
	"fmt"
	"log"
	"net/url"
	"reflect"
)

// Marshal to map[string]string
func Marshal(v interface{}) (map[string]interface{}, error) {
	return valueEncoder(reflect.ValueOf(v))
}

func valueEncoder(v reflect.Value) (map[string]interface{}, error) {
	if !v.IsValid() {
		return nil, fmt.Errorf("value invalid")
	}
	return typeEncoder(v, v.Type())
}

func typeEncoder(v reflect.Value, t reflect.Type) (map[string]interface{}, error) {
	switch t.Kind() {
	case reflect.Struct:
		return newStructEncoder(v, t)
	case reflect.Ptr:
		return typeEncoder(v.Addr(), t.Elem())
	default:
		return nil, fmt.Errorf("unsupport type: %v", t.Kind())
	}
}

func newStructEncoder(v reflect.Value, t reflect.Type) (map[string]interface{}, error) {
	fields := typeFields(t)
	m := map[string]interface{}{}
	for _, field := range fields {
		m[field.name] = ""
	}

	return m, nil
}

// MarshalToValues marshal to url.Values
func MarshalToValues(s interface{}) (*url.Values, error) {
	res, err := Marshal(s)
	if err != nil {
		return nil, err
	}

	param := url.Values{}
	for k, v := range res {
		param[k] = []string{fmt.Sprintf("%v", v)}
	}

	return &param, nil
}

// ConvertValuesToMap convert url.Values to map
func ConvertValuesToMap(val url.Values) map[string]string {
	res := map[string]string{}
	for k, v := range val {
		if len(v) == 1 {
			res[k] = v[0]
		} else if len(v) == 0 {
			log.Printf("url.Values key %s value has zero value", k)
		} else {
			log.Printf("url.Values key %s value has mulit value: %v", k, v)
		}
	}

	return res
}

package utils

import (
	"fmt"
	"strconv"
)

// ConvertToInt64 convert to int64
func ConvertToInt64(val interface{}) (int64, error) {
	f, err := ConvertToFloat64(val)
	return int64(f), err
}

// ConvertToUint64 convert to uint64
func ConvertToUint64(val interface{}) (uint64, error) {
	f, err := ConvertToFloat64(val)
	return uint64(f), err
}

// ConvertToFloat64 convert to float64
func ConvertToFloat64(val interface{}) (float64, error) {
	s, ok := val.(string)
	if ok {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, err
		}
		return f, nil
	}

	switch val.(type) {
	case int8:
		return float64(val.(int8)), nil
	case int16:
		return float64(val.(int16)), nil
	case int32:
		return float64(val.(int32)), nil
	case int64:
		return float64(val.(int64)), nil
	case int:
		return float64(val.(int)), nil
	case uint8:
		return float64(val.(uint8)), nil
	case uint16:
		return float64(val.(uint16)), nil
	case uint32:
		return float64(val.(uint32)), nil
	case uint64:
		return float64(val.(uint64)), nil
	case uint:
		return float64(val.(uint)), nil
	case float32:
		return float64(val.(float32)), nil
	case float64:
		return float64(val.(float64)), nil
	default:
		return 0, fmt.Errorf("invalid type")
	}
}

// ConvertToInt64Must convert to int64, panic if err occurs
func ConvertToInt64Must(val interface{}) int64 {
	i64, err := ConvertToInt64(val)
	if err != nil {
		panic(err.Error())
	}

	return i64
}

// ConvertToUint64Must convert to uint64
func ConvertToUint64Must(val interface{}) uint64 {
	f, err := ConvertToFloat64(val)
	if err != nil {
		panic(err.Error())
	}
	return uint64(f)
}

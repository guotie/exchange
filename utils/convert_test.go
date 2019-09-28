package utils

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	var (
		i     int32 = 1
		inter interface{}
	)

	inter = i
	fi, ok := inter.(float64)
	if ok {
		fmt.Println("convert int32 to float64 success")
	} else {
		fmt.Println("cannot convert")
	}

	fmt.Println(fi)
}

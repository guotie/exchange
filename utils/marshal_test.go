package utils

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	var s = struct {
		A string `json:"a"`
	}{}

	ret, err := Marshal(s)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	fmt.Println(ret)
}

package main

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	t.SkipNow() // 掉过当前test case,只能写在case开头
	res := Print1to20()
	fmt.Println("hey")
	if res != 210 {
		t.Errorf("Wrong result of Print1to20")
	}
}

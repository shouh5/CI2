package main

import (
	"fmt"
	"testing"
)
func testInput(t *testing.T){
	enter :=input()
	fmt.Print(enter)
	if enter!="hello"{
		t.Errorf("input is %s",enter)
	}
}
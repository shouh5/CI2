package main

import "testing"
func testInput(t *testing.T){
	enter :=input()
	if enter!="hello"{
		t.Errorf("input is %s",enter)
	}
}
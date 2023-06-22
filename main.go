package main

import (
	"fmt"
)

func input() string{
	return "hello"
	//return "hello";
	//fmt.Print("hello")
}
func main(){
	enter :=input()
	fmt.Print(enter)
	say()
	security()

	username := "admin"
	var password = "f62e5bcda4fae4f82370da0c6f20697b8f8447ef"
	//var password = os.Getenv("password")
	fmt.Println("Doing something with: ", username, password)
}
package main

import (
	"fmt"
	"os"
)

func security(){
	//var password = "123456"

	var password = os.Getenv("password")

	fmt.Printf("my password is %s", password)
}

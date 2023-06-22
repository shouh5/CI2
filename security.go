package main

import (
	"fmt"
)

func security(){
	var password = "123456"

	//var password = os.Getenv("password")

	fmt.Printf("my password is %s", password)
}

package main

import "fmt"

func main() {
	var username string = "rhschuld"
	var password string = "5404945066"
	fmt.Println("Authorization: Basic", username+":"+password)
}

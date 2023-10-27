package main

import "fmt"

func main(){
	fmt.Println("TAKING INPUT FROM USER")

	var name string

	fmt.Println("Enter your Name: ")
	fmt.Scanf("%s", &name)

	fmt.Println("The name of user is ", name)
}
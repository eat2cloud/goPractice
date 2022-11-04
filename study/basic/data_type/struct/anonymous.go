package main

import "fmt"

func structAnonymous() {
	var user struct {
		Name string
		Age  int
	}
	user.Name = "小王子"
	user.Age = 18
	fmt.Printf("%#v\n", user)
}

func main() {
	structAnonymous()
}

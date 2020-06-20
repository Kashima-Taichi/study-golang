package main

import "fmt"

func init() {
	fmt.Println("init execution")
}

func sub() {
	fmt.Println("sub function")
}

func main() {
	sub()
	fmt.Println("Hello world!", "second declaration")
}

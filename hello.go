package main

import "fmt"

func Name(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Name("world"))
}


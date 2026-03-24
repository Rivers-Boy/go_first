package main

import "fmt"

func Hello(content string) string {
	return "Hello " + content
}

func main() {
	fmt.Println(Hello("River"))
}

package main

import "fmt"

// 加入常量重构代码
const prefix = "Hello, "

func Hello(content string) string {
	return prefix + content
}

func main() {
	fmt.Println(Hello("River"))
}

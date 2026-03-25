package main

import "fmt"

// 加入常量重构代码
const prefix = "Hello, "

func Hello(content string, language string) string {

	prefixNew := ""

	// 使用 switch 重构
	switch language {
	case "Spanish":
		prefixNew = "Hola, "
	case "French":
		prefixNew = "Bonjour, "
	default:
		prefixNew = prefix
	}

	return prefixNew + content
}

func main() {
	fmt.Println(Hello("River", ""))
}

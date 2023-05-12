package logger

import "fmt"

func Info(data interface{}) {
	fmt.Println("log info: ", data)
}

func Error(data interface{}) {
	fmt.Println("Log error: ", data)
}

func Warning(data interface{}) {
	fmt.Println("Warning error: ", data)
}

func fatal(data interface{}) {
	println("Fatal error: ", data)
}

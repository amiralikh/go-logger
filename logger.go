package logger

import "fmt"

func Info(message string, params interface{}) {
	fmt.Println("log info: ", message, "params:", params)
}

func Error(message string, params interface{}) {
	fmt.Println("Log error: ", message, "params:", params)
}

func Warning(message string, params interface{}) {
	fmt.Println("Warning error: ", message, "params:", params)
}

func fatal(message string, params interface{}) {
	println("Fatal error: ", message, "params:", params)
}

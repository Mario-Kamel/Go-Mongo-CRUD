package utils

import "fmt"

type Logger interface {
	Log(message string)
}

type FmtLogger struct{}

func (fl FmtLogger) Log(message string) {
	fmt.Println(message)
}

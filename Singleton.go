package main

import (
	"fmt"
	"sync"
)

type Logger struct {
	logs []string
	mu   sync.Mutex
}

var (
	singleLogger *Logger
	once         sync.Once
)

func GetLogger() *Logger {
	once.Do(func() {
		singleLogger = &Logger{}
	})
	return singleLogger
}

func (l *Logger) Log(message string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logs = append(l.logs, message)
}

func (l *Logger) DisplayLogs() {
	l.mu.Lock()
	defer l.mu.Unlock()
	fmt.Println("Log Entries:")
	for _, log := range l.logs {
		fmt.Println(log)
	}
}

func main() {
	logger := GetLogger()

	logger.Log("Application started.")
	logger.Log("User logged in.")
	logger.Log("Data saved.")

	logger.DisplayLogs()
}

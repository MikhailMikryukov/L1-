package main

import "fmt"

// Target интерфейс
type Logger interface {
	Log(message string, level string)
}

// Adaptee 1 - существующая система логирования
type FileLogger struct{}

func (f *FileLogger) WriteToFile(msg string, logLevel string) {
	fmt.Printf("[FILE %s] %s\n", logLevel, msg)
}

// Adaptee 2 - другая система логирования
type ConsoleWriter struct{}

func (c *ConsoleWriter) Print(text string) {
	fmt.Printf("[CONSOLE] %s\n", text)
}

// Adapter для FileLogger
type FileLoggerAdapter struct {
	fileLogger *FileLogger
}

func NewFileLoggerAdapter(logger *FileLogger) *FileLoggerAdapter {
	return &FileLoggerAdapter{fileLogger: logger}
}

func (a *FileLoggerAdapter) Log(message string, level string) {
	a.fileLogger.WriteToFile(message, level)
}

// Adapter для ConsoleWriter
type ConsoleAdapter struct {
	console *ConsoleWriter
}

func NewConsoleAdapter(console *ConsoleWriter) *ConsoleAdapter {
	return &ConsoleAdapter{console: console}
}

func (a *ConsoleAdapter) Log(message string, level string) {
	formattedMsg := fmt.Sprintf("[%s] %s", level, message)
	a.console.Print(formattedMsg)
}

// L1.21 Паттерн «Адаптер»
func main() {
	// Создаем адаптеры для разных систем логирования
	fileLogger := &FileLogger{}
	consoleWriter := &ConsoleWriter{}

	fileAdapter := NewFileLoggerAdapter(fileLogger)
	consoleAdapter := NewConsoleAdapter(consoleWriter)

	// Используем единый интерфейс
	loggers := []Logger{fileAdapter, consoleAdapter}

	for _, logger := range loggers {
		logger.Log("Application started", "INFO")
		logger.Log("Error occurred", "ERROR")
	}
}

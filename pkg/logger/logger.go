package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
	FATAL   = "FATAL"
)

var (
	logFile       *os.File
	consoleLogger *log.Logger
	fileLogger    *log.Logger
	customLogger  Logger
)

// Custom logger struct
type Logger interface {
	LogMessage(level, msg string, logToConsole bool)
}

// Set custom logger for testing
func SetLogger(logger Logger) {
	customLogger = logger
}

func findGoModDir(startDir string) (string, error) {
	// Traverse upwards to find the directory containing go.mod
	for {
		// Check if the current directory contains go.mod
		if _, err := os.Stat(filepath.Join(startDir, "go.mod")); err == nil {
			return startDir, nil
		}

		// Move up one directory level
		parentDir := filepath.Dir(startDir)

		// If we are already at the root directory, stop
		if parentDir == startDir {
			break
		}

		// Move to the parent directory
		startDir = parentDir
	}

	return "", fmt.Errorf("go.mod not found")
}

// Initialize loggers
func init() {
	// Create log file with a timestamp suffix
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
		return
	}

	// Get the absolute path to the "log" directory
	// Find the root of the project by looking for go.mod
	rootDir, err := findGoModDir(wd)
	if err != nil {
		fmt.Println("Error finding project root:", err)
		return
	}
	logDir := filepath.Join(rootDir, "/log/")
	timestamp := time.Now().Format("20060102_150405")
	logFileName := fmt.Sprintf("%s/app.%s.log", logDir, timestamp)

	logFile, err = os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	// Log to stdout
	consoleLogger = log.New(os.Stdout, "", log.LstdFlags)

	// Log to file
	fileLogger = log.New(logFile, "", log.LstdFlags)
}

// Log messages with severity levels
func LogMessage(level, msg string, logToConsole bool) {
	if customLogger != nil {
		// Use custom logger if it's set
		customLogger.LogMessage(level, msg, logToConsole)
		return
	}

	// Log to console if needed
	if logToConsole {
		printToConsole(level, msg)
	}

	// Log to file
	printToFile(level, msg)
}

// Print to console
func printToConsole(level, msg string) {
	switch level {
	case INFO:
		consoleLogger.SetPrefix("INFO: ")
		consoleLogger.Println(msg)
	case WARNING:
		consoleLogger.SetPrefix("WARNING: ")
		consoleLogger.Println(msg)
	case ERROR:
		consoleLogger.SetPrefix("ERROR: ")
		consoleLogger.Println(msg)
	case FATAL:
		consoleLogger.SetPrefix("FATAL: ")
		consoleLogger.Println(msg)
	default:
		consoleLogger.SetPrefix("INFO: ")
		consoleLogger.Println(msg)
	}
}

// Print to log file
func printToFile(level, msg string) {
	switch level {
	case INFO:
		fileLogger.SetPrefix("INFO: ")
		fileLogger.Println(msg)
	case WARNING:
		fileLogger.SetPrefix("WARNING: ")
		fileLogger.Println(msg)
	case ERROR:
		fileLogger.SetPrefix("ERROR: ")
		fileLogger.Println(msg)
	case FATAL:
		fileLogger.SetPrefix("FATAL: ")
		fileLogger.Println(msg)
	default:
		fileLogger.SetPrefix("INFO: ")
		fileLogger.Println(msg)
	}
}

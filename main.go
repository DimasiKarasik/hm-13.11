package main
//Создал чтобы закомитить
import (
	"fmt"
	"strings"
	"time"
)

type Result struct {
	Level   string
	Message string
}

func main() {
	startTime := time.Now()

	logs := []string{
		"INFO|2024-01-01 10:00:00|Server started",
		"WARN|2024-01-01 10:01:00|Low memory",
		"ERROR|2024-01-01 10:02:00|Connection failed",
		"INFO|2024-01-01 10:03:00|User login",
		"BROKEN LINE WITHOUT PIPES",
		"ERROR|2024-01-01 10:04:00|Timeout",
	}

	results := make(chan Result, len(logs))

	var infoCount, warnCount, errorCount int

	for _, log := range logs {
		parts := strings.Split(log, "|")
		if len(parts) == 3 {
			results <- Result{
				Level:   parts[0],
				Message: parts[2],
			}
		}
	}
	close(results)

	for res := range results {
		switch res.Level {
		case "INFO":
			infoCount++
		case "WARN":
			warnCount++
		case "ERROR":
			errorCount++
		}
	}

	totalValidLogs := infoCount + warnCount + errorCount

	fmt.Printf("Количество событий по уровням:\n")
	fmt.Printf("INFO:  %d\n", infoCount)
	fmt.Printf("WARN:  %d\n", warnCount)
	fmt.Printf("ERROR: %d\n", errorCount)

	if totalValidLogs > 0 {
		errorPercentage := (float64(errorCount) / float64(totalValidLogs)) * 100
		fmt.Printf("\nОбщее количество ошибок: %.1f%%\n", errorPercentage)
	}

	executionTime := time.Since(startTime)
	fmt.Printf("Время выполнения программы: %v\n", executionTime)
}
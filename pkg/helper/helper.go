package helper

import (
	"time"
)

// GetCurrentTime returns required time format as string
func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05") // https://golang.org/src/time/format.go
}

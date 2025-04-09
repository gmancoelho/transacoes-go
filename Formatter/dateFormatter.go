package formatter

import (
	"fmt"
	"time"
)

const (
	dateFormat = "02/01/2006 15:04:05"
)

// ConvertISO8601ToString converts an ISO 8601 date string to a formatted string.
func ConvertISO8601ToString(isoDate string) (string, error) {
	parsedTime, err := time.Parse(time.RFC3339, isoDate)
	if err != nil {
		return "", fmt.Errorf("invalid ISO 8601 date: %v", err)
	}

	return parsedTime.Format(dateFormat), nil
}

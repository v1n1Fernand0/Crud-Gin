package utils

import (
	"time"

	"github.com/google/uuid"
)

func GetCurrentTime() time.Time {
	return time.Now()
}

func GenerateUUID() string {
    return uuid.New().String()
}
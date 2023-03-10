package common

import (
	"math/rand"
	"strings"
	"time"
)

func UUID(uuid string, delimiter string, nbrOfGroups int) string {
	newUUID := ""

	uuidList := strings.Split(uuid, delimiter)
	for key, value := range uuidList {
		if key < nbrOfGroups {
			newUUID = newUUID + value
		}
	}

	return newUUID
}

func FormattedTime(timeNow time.Time, timeFormat string) string {
	return timeNow.Format(timeFormat)
}

func GenerateFloat(nbrRange float64, minVal float64) int {
	newRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Return randomly generated float64
	return int((newRand.Float64() * nbrRange) + minVal)
}

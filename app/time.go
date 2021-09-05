package app

import (
	"time"
)

func getTime(timezones []string) map[string]string {
	result := make(map[string]string)

	if timezones == nil {
		result["current_time"] = time.Now().UTC().String()
	} else {
		for _, tz := range timezones {

			loc, err := time.LoadLocation(tz)

			if err != nil {
				result["invalid timezone"] = "invalid timezone provided"
			} else {
				result[tz] = time.Now().In(loc).String()
			}
		}
	}

	return result
}

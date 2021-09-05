package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func handleTime(writer http.ResponseWriter, req *http.Request) {
	timezones := req.URL.Query().Get("tz")
	var mapTimes map[string]string

	if timezones == "" {
		mapTimes = getTime(nil)
	} else {
		timezonesStrings := strings.Split(timezones, ",")
		mapTimes = getTime(timezonesStrings)
	}

	if len(mapTimes) == 1 && mapTimes["invalid timezone"] != "" {
		writer.WriteHeader(404)
		fmt.Fprint(writer, "Timezones not found!")
	} else {
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(mapTimes)
	}

}

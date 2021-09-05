Usage:

Run the code locally: go run main.go

http://localhost:9090/api/time --> for UTC time
http://localhost:9090/api/time?tz=America/New_York for the time in New York, or whatever time you want
http://localhost:9090/api/time?tz=America/New_York,Europe/Athens,Indian/Comoro for multiple timezones (sepparated by commas)

Complete list of string timezones: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
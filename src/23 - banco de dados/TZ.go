package main

// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".

// fonte: https://yourbasic.org/golang/time-change-convert-location-timezone/

import (
	"fmt"
	"time"
)

func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func main() {

	for _, name := range []string{
		"",
		"Local",
		"Asia/Shanghai",
		"America/Sao_Paulo",
	} {
		t, err := TimeIn(time.Now(), name)
		if err == nil {
			fmt.Println(t.Location(), t.Format("15:04"))
		} else {
			fmt.Println(name, "<time unknown>")
		}
	}
}

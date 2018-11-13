package main

import (
	"fmt"
	"time"
)

func main() {
	location := "America/New_York"
	loc, err := time.LoadLocation(location)
	if err != nil {
		fmt.Printf("[FAIL] load location %q from tz data: %s\n", location, err)
	} else {
		fmt.Printf("[ OK ] loaded location: %s\n", loc.String())
	}
}

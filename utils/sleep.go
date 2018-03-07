package utils

import "time"

// wrap the tedious time.Sleep function
func Sleep(dur int) {
	time.Sleep(time.Duration(dur) * time.Second)
}

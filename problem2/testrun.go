package testrun

import (
	"fmt"
	"time"
)

func testrun() {
	t := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-t.C:
			fmt.Println("hi")
		}
	}
}

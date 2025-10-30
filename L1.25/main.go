package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	<-time.After(duration)
}

// L1.25 Своя функция Sleep
func main() {
	start := time.Now()

	sleep(time.Second)

	fmt.Println(time.Since(start))
}

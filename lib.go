package main

import (
	"fmt"
	"time"
)

func test() {
	fmt.Println("test " + time.Now().Weekday().String())
}

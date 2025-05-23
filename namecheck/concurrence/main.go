package main

import (
	"fmt"
	"sync"
)

func grindCoffeeBeans() {
	fmt.Println("grindCoffeeBeans")
}

func frotMilk() {
	fmt.Println("frotMilk")
}

func runFunc(fn func(), wg *sync.WaitGroup) {
	defer wg.Done()
	fn()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go runFunc(grindCoffeeBeans, &wg)
	wg.Add(1)
	go runFunc(frotMilk, &wg)

	wg.Wait()
}

package main

import "fmt"

func grindCoffeeBeans() {
	fmt.Println("grindCoffeeBeans")
}

func frotMilk() {
	fmt.Println("frotMilk")
}

func main() {
	go grindCoffeeBeans()
	go frotMilk()
}

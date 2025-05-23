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

// version qui provoque une race condition, à tester avec go run --race main.go
func main() {
	countCh := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		count := <-countCh // Lire depuis le canal
		count++            // Incrémenter la valeur
		countCh <- count   // Écrire la valeur mise à jour dans le canal
	}()

	countCh <- 0              // Envoyer 0 initialement
	updatedCount := <-countCh // Lire la valeur mise à jour
	if updatedCount == 0 {
		fmt.Println(updatedCount)
	} else {
		fmt.Println(updatedCount)
	}

	wg.Wait()
}

/*
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go runFunc(grindCoffeeBeans, &wg)
	wg.Add(1)
	go runFunc(frotMilk, &wg)

	wg.Wait()
}
*/

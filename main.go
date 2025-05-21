package main

import (
	"fmt"
	"time"
)

func inc(p *int) {
	*p++
}

type Celsius int
type Farenheit int

func (c Celsius) String() string {
	return fmt.Sprintf("%d°C", c)
}

func (f Farenheit) String() string {
	return fmt.Sprintf("%d°F", f)
}

func main() {
	var temp Celsius = 24
	var tempF Farenheit = Farenheit(temp + 32)

	fmt.Println("Hello, World! " + temp.String() + " / " + tempF.String())

	if time.Now().Month() == time.May {
		fmt.Println("It's May")
	} else {
		fmt.Println("It's not May")
	}

	i := 42
	p := &i

	test()

	// pointers
	fmt.Println(p)  // adresse mémoire
	fmt.Println(*p) // valeur pointée
	*p++
	fmt.Println(*p) // valeur pointée
	inc(p)
	inc(&i)
	fmt.Println(*p)       // valeur pointée
	fmt.Printf("%T\n", p) // type de la valeur pointée
	fmt.Printf("%T\n", i) // type de la variable
}

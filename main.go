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

func testslice(s []int) {
	for i, value := range s {
		s[i] = value * 3
	}
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

	// tableaux
	slice := make([]int, 6)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	slice[4] = 5
	slice[5] = 10
	testslice(slice)
	fmt.Println(slice)

	slice2 := make([]int, 3)
	testslice(slice2)
	fmt.Println(slice2)

	red := slice[4:5]
	fmt.Println(red)
	red[0] = 99
	fmt.Println(slice)
}

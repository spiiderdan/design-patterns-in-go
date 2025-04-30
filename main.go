package main

import (
	"fmt"
	"github.com/spiiderdan/design-patterns-in-go/creational/singleton"
)


func main() {
	c1 := singleton.GetInstance()
	c2 := singleton.GetInstance()

	c1.Add()
	c2.Add()

	fmt.Println("Value from c1:", c1.Get()) // should print 2
	fmt.Println("Value from c2:", c2.Get()) // should also print 2
	fmt.Println("Same instance?", c1 == c2) // true
}

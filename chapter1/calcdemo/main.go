package main

import (
	"fmt"

	"github.com/NikolaM-Dev/go-web/chapter1/calc"
)

func main() {
	x, y := 10, 5
	fmt.Println(calc.Add(x, y))
	fmt.Println(calc.Substract(x, y))
}

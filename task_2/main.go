package main

import (
	"calc/Calc"
	"fmt"
	"flag"
)

func main() {
	flag.Parse()
	if flag.NArg() == 1 {
		if result, err := calc.Calc(flag.Args()[0]); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	} else {
		fmt.Println("Не введено выражение")
	}
}

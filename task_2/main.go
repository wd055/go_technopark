package main

import (
	"calc/Calc"
	"fmt"
)

func main() {
	fmt.Println(calc.Calc("2+(1*3)/4"))
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

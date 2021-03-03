package main

import (
	"calc/Calc"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Не введено выражение")
		os.Exit(1)
		return
	}

	result, err := calc.Calc(flag.Args()[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	fmt.Println(result)
}

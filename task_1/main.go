package main

import (
	"flag"
	"fmt"
	"os"
	"uniq/Uniq"
)

func main() {
	if flag.NArg() > 2 {
		fmt.Println(`Ошибка при вводе аргументов. Используйте -help перед вызовом функции.`)
		os.Exit(1)
		return
	}

	flags, err := getFlags()
	if err != nil || flags.Help {
		fmt.Println(`Ошибка при вводе аргументов. Используйте -help перед вызовом функции.`)
		os.Exit(1)
		return
	}

	rows, err := input()
	if err != nil {
		os.Exit(1)
		return
	}

	result := uniq.Uniq(rows, flags)

	if err := output(result); err != nil {
		os.Exit(1)
		fmt.Println(err)
	}
}

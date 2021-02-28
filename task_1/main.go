package main

import (
	"flag"
	"log"
	"uniq/Uniq"
)

func main() {
	if flag.NArg() > 2 {
		log.Fatal(`Ошибка при вводе аргументов. Используйте -help перед вызовом функции.`)
		return
	}

	flags, err := getFlags()
	if err || flags.Help {
		log.Fatal(`Ошибка при вводе аргументов. Используйте -help перед вызовом функции.`)
		return
	}

	rows := input()
	result := uniq.Uniq(rows, flags)

	if err := output(result); err != nil {
		log.Fatal(err)
	}
}

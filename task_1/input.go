package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"uniq/Uniq"
)

func init() {
	flag.Bool("c", false, "Подсчитать количество встречаний строки во входных данныx.")
	flag.Bool("d", false, "Вывести только те строки, которые повторились во входных данных.")
	flag.Bool("u", false, "Вывести только те строки, которые не повторились во входных данных.")
	flag.Int("f", 0, "Не учитывать первые n полей в строке.")
	flag.Int("s", 0, "Не учитывать первые n символов в строке.")
	flag.Bool("i", false, "Не учитывать регистр букв.")
	flag.Bool("help", false, "Вывести помощь по команде")
	flag.Parse()
}

func getFlags() (flags uniq.FlagsStruct, err bool) {
	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "c":
			flags.C = true
		case "d":
			flags.D = true
		case "u":
			flags.U = true
		case "f":
			if num, err := strconv.Atoi(f.Value.String()); err != nil {
				flags.F = num
			}
		case "s":
			if num, err := strconv.Atoi(f.Value.String()); err != nil {
				flags.S = num
			}
		case "i":
			flags.I = true
		case "help":
			flags.Help = true
		}
	})

	if (flags.C && flags.D) || (flags.D && flags.U) || (flags.C && flags.U) {
		err = true
	}

	return flags, err
}

func input() (result []string) {
	in := bufio.NewScanner(os.Stdin)

	if flag.NArg() >= 1 {
		if fileIn, err := os.Open(flag.Args()[0]); err != nil {
			panic(err)
		} else {
			in = bufio.NewScanner(fileIn)
			defer fileIn.Close()
		}
	}

	for in.Scan() {
		if in.Err() != nil {
			panic(in.Err())
		}
		result = append(result, in.Text())
	}
	return
}

func output(result []string) error {
	out := bufio.NewWriter(os.Stdout)
	if flag.NArg() == 2 {
		if fileOut, err := os.Create(flag.Args()[1]); err != nil {
			panic(err)
		} else {
			out = bufio.NewWriter(fileOut)
			defer fileOut.Close()
		}
	}

	for _, str := range result {
		if _, err := fmt.Fprintln(out, str); err != nil {
			return err
		}
	}
	out.Flush()
	return nil
}
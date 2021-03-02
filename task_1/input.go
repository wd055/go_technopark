package main

import (
	"bufio"
	"errors"
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

func getFlags() (flags uniq.FlagsStruct, err error) {
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
		err = errors.New("Ошибка ввода флагов")
	}

	return flags, err
}

func input() ([]string, error) {
	var result []string
	var in *bufio.Scanner

	if flag.NArg() >= 1 {
		if fileIn, err := os.Open(flag.Args()[0]); err != nil {
			return nil, errors.New("Ошибка открытия файла")
		} else {
			defer fileIn.Close()
			in = bufio.NewScanner(fileIn)
		}
	} else {
		in = bufio.NewScanner(os.Stdin)
	}

	for in.Scan() {
		result = append(result, in.Text())
	}
	if in.Err() != nil {
		return nil, errors.New("Ошибка чтения с потока")
	}

	return result, nil
}

func output(result []string) error {
	var out *bufio.Writer

	if flag.NArg() == 2 {
		if fileOut, err := os.Create(flag.Args()[1]); err != nil {
			return err
		} else {
			out = bufio.NewWriter(fileOut)
			defer fileOut.Close()
		}
	} else {
		out = bufio.NewWriter(os.Stdout)
	}

	for _, str := range result {
		if _, err := fmt.Fprintln(out, str); err != nil {
			return err
		}
	}
	out.Flush()
	return nil
}

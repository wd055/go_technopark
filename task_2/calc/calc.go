package calc

import (
	stack "calc/calc/stack"
	"errors"
	"math"
)

func genNumByStr(number []float64) float64 {
	size := len(number)
	num := 0.0
	powTen := float64(math.Pow10(size - 1))
	for _, digit := range number {
		num += digit * powTen
		powTen /= 10
	}
	return num
}

func calculate(numberR float64, numberL float64, operator float64) (float64, error) {
	switch operator {
	case float64('+'):
		return (numberL + numberR), nil
	case float64('-'):
		return (numberL - numberR), nil
	case float64('*'):
		return (numberL * numberR), nil
	case float64('/'):
		return (numberL / numberR), nil
	}
	return 0, errors.New("Что за оператор ты кинул сюда?")
}

func Calc(str string) (float64, error) {
	var numbers stack.Stack
	var operators stack.Stack
	countBracket := 0

	priority := map[float64]float64{
		'+': 1.0,
		'-': 1.0,
		'/': 2.0,
		'*': 2.0,
	}

	var number []float64
	for _, ch := range str {
		if ch >= '0' && ch <= '9' {
			number = append(number, float64(ch-48))
		} else {
			if number != nil {
				numbers.Push(genNumByStr(number))
			}
			number = nil

			if _, contain := priority[float64(ch)]; contain {
				if operation, err := operators.Top(); err == nil {
					if priority[float64(ch)] <= priority[float64(operation)] {
						var numR, numL float64
						if numR, err := numbers.Pop(); err != nil {
							return 0.0, err
						}
						if numL, err := numbers.Pop(); err != nil {
							return 0.0, err
						}
						if result, err := calculate(numR, numL, float64(operation)); err == nil {
							numbers.Push(result)
						} else {
							return 0, err
						}
						if _, err := operators.Pop(); err != nil {
							return 0.0, err
						}
					}
				} else {
					return 0.0, err
				}
				operators.Push(float64(ch))
			} else if ch == '(' {
				countBracket++
				operators.Push(float64(ch))
			} else if ch == ')' {
				if countBracket < 1 {
					panic("Скобочки посчитай, у тебя что-то не то!")
				}
				for operation, err := operators.Pop(); err == nil && operation != '('; operation, err = operators.Pop() {

					var numR, numL float64
					if numR, err := numbers.Pop(); err != nil {
						return 0.0, err
					}
					if numL, err := numbers.Pop(); err != nil {
						return 0.0, err
					}
					if result, err := calculate(numR, numL, float64(operation)); err == nil {
						numbers.Push(result)
					} else {
						return 0, err
					}
					if _, err := operators.Pop(); err != nil {
						return 0.0, err
					}
				}
				countBracket--
			} else {
				panic("Ты шо ввел, дурашка!?)")
			}
		}
	}

	if number != nil {
		numbers.Push(genNumByStr(number))
	}
	for operation, err := operators.Pop(); err == nil; operation = operators.Pop() {
		numbers.Push(calculate(numbers.Pop(), numbers.Pop(), operation))
	}
	return numbers[0], nil
}

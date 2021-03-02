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

func calculate(numbers *stack.Stack, operators *stack.Stack, operator float64) error {
	numberR, err := (*numbers).Pop()
	if err != nil {
		return err
	}
	numberL, err := (*numbers).Pop()
	if err != nil {
		return err
	}

	var result float64
	switch operator {
	case float64('+'):
		result = numberL + numberR
	case float64('-'):
		result = numberL - numberR
	case float64('*'):
		result = numberL * numberR
	case float64('/'):
		result = numberL / numberR
	default:
		return errors.New("Что за оператор ты кинул сюда?")
	}

	(*numbers).Push(result)

	return nil
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
				if operation, err := operators.Top(); err == nil && priority[float64(ch)] <= priority[float64(operation)] {
					if err := calculate(&numbers, &operators, operation); err != nil {
						return 0, err
					}
					if _, err := operators.Pop(); err != nil {
						return 0.0, err
					}
				}
				operators.Push(float64(ch))
			} else if ch == '(' {
				countBracket++
				operators.Push(float64(ch))
			} else if ch == ')' {
				if countBracket < 1 {
					return 0.0, errors.New("Скобочки посчитай, у тебя что-то не то!")
				}
				for operation, err := operators.Pop(); err == nil && operation != '('; operation, err = operators.Pop() {
					if err := calculate(&numbers, &operators, operation); err != nil {
						return 0, err
					}
				}
				countBracket--
			}
		}
	}

	if number != nil {
		numbers.Push(genNumByStr(number))
	}
	for operation, err := operators.Pop(); err == nil; operation, err = operators.Pop() {
		if err := calculate(&numbers, &operators, operation); err != nil {
			return 0, err
		}
	}
	return numbers[0], nil
}

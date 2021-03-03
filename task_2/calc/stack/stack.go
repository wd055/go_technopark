package stack

import (
	"errors"
)

type Stack []float64

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

func (stack *Stack) Push(pushData float64) {
	*stack = append(*stack, pushData)
}

func (stack *Stack) Pop() (float64, error) {
	if (*stack).IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	popData := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return popData, nil
}

func (stack Stack) Top() (float64, error) {
	if stack.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	return stack[len(stack)-1], nil
}

package stack

import "testing"

func TestStack_Push(t *testing.T) {
	stack := Stack{}

	stack.Push(123)

	if num := stack.Pop(); num != 123 {
		t.Fatal("Push error")
	}
}

func TestStack_Top(t *testing.T) {
	stack := Stack{}

	stack.Push(123)

	num1 := stack.Top()
	num2 := stack.Pop()

	if num1 != num2 {
		t.Fatal("Top ≠ Error")
	}
}

func TestStack_IsEmpty(t *testing.T) {
	stack := Stack{}

	if !stack.IsEmpty() {
		t.Fatal("IsEmpty error")
	}
}

func TestStack_NotIsEmpty(t *testing.T) {
	stack := Stack{}

	stack.Push(123)

	if stack.IsEmpty() {
		t.Fatal("IsEmpty error")
	}
}

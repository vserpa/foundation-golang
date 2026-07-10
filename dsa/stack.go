package main

import "errors"

type Stack struct {
	values []interface{}
}

func (stack Stack) Size() int {
	return len(stack.values)
}

func (stack Stack) IsEmpty() bool {
	return stack.Size() == 0
}

func (stack *Stack) Push(value interface{}) {
	stack.values = append(stack.values, value)
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	value := stack.values[stack.Size()-1]
	stack.values = stack.values[:stack.Size()-1]
	return value, nil
}

package main

import "fmt"

type Stack struct { data []int64 }

func NewStack() *Stack {
	return &Stack{data: make([]int64,0,2)}
}

func (stack *Stack) Push(n int64)  {
	stack.data = append(stack.data, n)
}

func (stack *Stack) Pop() (ret int64) {
	ret = stack.data[len(stack.data)-1]
	stack.data = stack.data[:len(stack.data)-1]
	return
}

func main() {
	stack := NewStack()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
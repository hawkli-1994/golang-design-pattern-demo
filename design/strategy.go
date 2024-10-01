package design

import (
	"fmt"
	"testing"
)

// 策略模式
type Strategy interface {
	do(int, int) int
}

type add struct{}

func (a add) do(x, y int) int {
	return x + y
}

type reduce struct{}

func (r reduce) do(x, y int) int {
	return x - y
}

type Operator struct {
	strategy Strategy
}

func (operator *Operator) setStrategy(strategy Strategy) {
	operator.strategy = strategy
}

func (operator *Operator) calculate(a, b int) int {
	return operator.strategy.do(a, b)
}

func TestStrategy(t *testing.T) {
	operator := Operator{}

	operator.setStrategy(&add{})
	result := operator.calculate(1, 2)
	fmt.Println("add:", result)

	operator.setStrategy(&reduce{})
	result = operator.calculate(2, 1)
	fmt.Println("reduce:", result)
}

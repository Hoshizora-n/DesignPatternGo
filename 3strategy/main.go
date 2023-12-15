package strategy

import "fmt"

type Calculator struct {
	strategy CalculateStrategy
}

func NewCalculator() *Calculator {
	return &Calculator{
		strategy: SumStrat,
	}
}

func (r *Calculator) SetStrategy(strategy CalculateStrategy) {
	r.strategy = strategy
}

func (r *Calculator) Calculate(a, b int) {
	res := r.strategy.calculate(a, b)
	fmt.Println(res)
}

func Run() {
	calculator := NewCalculator()
	calculator.Calculate(7, 2)
	calculator.SetStrategy(SubStrat)
	calculator.Calculate(7, 2)
	calculator.SetStrategy(MulStrat)
	calculator.Calculate(7, 2)
}

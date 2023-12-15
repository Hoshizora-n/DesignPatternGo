package strategy

var (
	SumStrat = Sum{}
	SubStrat = Subtract{}
	MulStrat = Multiply{}
)

type CalculateStrategy interface {
	calculate(a, b int) int
}

type Sum struct{}

func (s Sum) calculate(a, b int) int {
	return a + b
}

type Subtract struct{}

func (s Subtract) calculate(a, b int) int {
	return a - b
}

type Multiply struct{}

func (s Multiply) calculate(a, b int) int {
	return a * b
}

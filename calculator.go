package main

import "fmt"

type Calculator struct {
	Result float64
}

func (calc *Calculator) Add(a, b float64) (float64, error) {
	sum := a + b
	calc.Result = calc.Result + sum
	return sum, nil
}

func (calc *Calculator) Div(a, b float64) (float64, error) {
	if b == float64(0) {
		return 0, DivByZero(0)
	}
	div := a / b
	calc.Result = calc.Result + div
	return div, nil
}

func (calc Calculator) String() string {
	return fmt.Sprintf("The result is %v",  calc.Result)
}

type DivByZero float64

func (err DivByZero) Error() string {
	return fmt.Sprintf("Cannot divide by number: %v", float64(err))
}

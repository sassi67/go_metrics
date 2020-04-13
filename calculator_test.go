package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/rcrowley/go-metrics"
	"github.com/stretchr/testify/assert"
)

const loops = 100

func TestCalculator(t *testing.T) {
	assertion := assert.New(t)

	registry := metrics.NewRegistry()
	counterAddOk := metrics.NewCounter()
	registry.Register("counter_add_ok", counterAddOk)
	counterAddErr := metrics.NewCounter()
	registry.Register("counter_add_err", counterAddErr)

	counterDivOk := metrics.NewCounter()
	registry.Register("counter_div_ok", counterDivOk)
	counterDivErr := metrics.NewCounter()
	registry.Register("counter_div_err", counterDivErr)

	var wg sync.WaitGroup

	workerAdd := func(calc *Calculator, a float64, b float64, wg *sync.WaitGroup) {
		defer wg.Done()
		result, err := calc.Add(float64(a), float64(b))
		if err == nil {
			counterAddOk.Inc(1)
			fmt.Printf("Add(%v, %v) = %v\n", a, b, result)
		} else {
			counterAddErr.Inc(1)
		}
	}

	workerDiv := func(calc *Calculator, a float64, b float64, wg *sync.WaitGroup) {
		defer wg.Done()
		result, err := calc.Div(float64(a), float64(b))
		if err == nil {
			counterDivOk.Inc(1)
			fmt.Printf("Add(%v, %v) = %v\n", a, b, result)
		} else {
			counterDivErr.Inc(1)
		}
	}

	calcAdd := &Calculator{}
	for i := 0; i < loops; i++ {
		wg.Add(1)
		go workerAdd(calcAdd, float64(i), float64(i), &wg)
	}
	calcDiv := &Calculator{}
	for i := 0; i < loops; i++ {
		wg.Add(1)
		go workerDiv(calcDiv, float64(i), float64(i), &wg)
	}

	wg.Wait()

	assertion.Equal(counterAddOk.Count(), int64(loops))
	assertion.Equal(counterAddErr.Count(), int64(0))
	assertion.Equal(counterDivOk.Count(), int64(loops-1))
	assertion.Equal(counterDivErr.Count(), int64(1))
}

// Package models
package models

import "fmt"

func init() {
	fmt.Println("models initialized")
}

type CalculationHistory struct {
	PreviousExpression        string
	PreviousCalculationResult CalcResult
}

var CalculationHistories = []CalculationHistory{
	{
		PreviousExpression:        "1+1",
		PreviousCalculationResult: 2,
	},
	{
		PreviousExpression:        "2+2",
		PreviousCalculationResult: 2,
	},
}

type CalcResult float64

var CalcResults []CalcResult

// Package models
package models

import "fmt"

func init() {
	fmt.Println("models initialized")
}

type CalcResult float64

type CalculationHistory struct {
	PreviousExpression        string
	PreviousCalculationResult CalcResult
}

var CalculationHistories = []CalculationHistory{}

var CalcResults []CalcResult

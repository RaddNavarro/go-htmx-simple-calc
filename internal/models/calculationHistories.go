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

var CalculationHistories = []CalculationHistory{}

type CalcResult float64

var CalcResults []CalcResult

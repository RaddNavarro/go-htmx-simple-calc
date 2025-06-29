package main

type CalculationHistory struct {
	PreviousExpression        string
	PreviousCalculationResult int
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

type TestResults struct {
	Result int
}

var Results = []TestResults{}

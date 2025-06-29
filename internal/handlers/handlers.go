// Package handlers
package handlers

import (
	"fmt"
	"net/http"

	"github.com/RaddNavarro/simple-calculator/internal/components"
	"github.com/RaddNavarro/simple-calculator/internal/models"
	"github.com/expr-lang/expr"
)

func Home(w http.ResponseWriter, r *http.Request) {
	component := components.Home()
	component.Render(r.Context(), w)
}

func ViewCalculationsHistory(w http.ResponseWriter, r *http.Request) {
	component := components.CalculationHistoriesList(models.CalculationHistories)
	component.Render(r.Context(), w)
}

func Calculate(w http.ResponseWriter, r *http.Request) {
	expression := r.PostFormValue("test-expression")

	program, err := expr.Compile(expression, expr.AsFloat64())
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
	result := models.CalcResult(output.(float64))
	component := components.DisplayResult(result)
	component.Render(r.Context(), w)

	models.CalcResults = append(models.CalcResults, result)

	currentCalculation := models.CalculationHistory{PreviousExpression: expression, PreviousCalculationResult: result}

	models.CalculationHistories = append(models.CalculationHistories, currentCalculation)
}

func GetExpression(w http.ResponseWriter, r *http.Request) {
	expression := r.PostFormValue("test-expression")

	program, err := expr.Compile(expression, expr.AsFloat64())
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}

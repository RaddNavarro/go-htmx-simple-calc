package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RaddNavarro/simple-calculator/internal/handlers"
)

// testing git fugitive
func main() {
	fmt.Println("Hello WOrld")

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/viewHistory", handlers.ViewCalculationsHistory)
	http.HandleFunc("/calculate", handlers.Calculate)
	http.HandleFunc("/get-expression", handlers.GetExpression)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	log.Fatal(http.ListenAndServe(":8000", nil))
}

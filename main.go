package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/RaddNavarro/simple-calculator/internal/handlers"
)

var static embed.FS

func main() {
	fmt.Println("Website running")

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/viewHistory", handlers.ViewCalculationsHistory)
	http.HandleFunc("/calculate", handlers.Calculate)
	http.HandleFunc("/get-expression", handlers.GetExpression)
	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/static/", http.StripPrefix("/static", fs))
	http.Handle("/static/", http.FileServer(http.FS(static)))

	log.Fatal(http.ListenAndServe(":8000", nil))
}

package main

import "fmt"

// Gibt Ergebnisse von Products() auf die Konsole aus.
// Der Kommentar unten gibt die erwarteten Ergebnisse an.
// Automatische Prüfung mittels des Befehls "go test" (statt "go run").
func ExampleProducts() {
	fmt.Println(Products([]int{1, 3, 5, 7}))
	fmt.Println(Products([]int{1, 1, 2, 80}))
	fmt.Println(Products([]int{7, 3, 1, 2}))
	fmt.Println(Products([]int{3, 3, 0, 2}))

	// Output:
	// [1 3 15 105]
	// [1 1 2 160]
	// [7 21 21 42]
	// [3 9 0 0]
}

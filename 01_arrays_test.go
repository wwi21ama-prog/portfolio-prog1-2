package main

import (
	"fmt"
)

// Gibt Ergebnisse von Multn() auf die Konsole aus.
// Der Kommentar unten gibt die erwarteten Ergebnisse an.
// Automatische Prüfung mittels des Befehls "go test" (statt "go run").
func ExampleMultn() {
	v1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12} // Produkt: 1*6*11 == 66
	v2 := []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1} // Produkt: 12*7*2 = 168
	v3 := []int{3, 29, 4, 0, 42, 2, 38}                // Produkt: 3*2 = 6
	v4 := []int{}                                      // Leere Liste, Produkt sollte 1 sein

	fmt.Println(Multn(v1, 5))
	fmt.Println(Multn(v2, 5))
	fmt.Println(Multn(v3, 5))
	fmt.Println(Multn(v4, 5))

	// Output:
	// 66
	// 168
	// 6
	// 1
}

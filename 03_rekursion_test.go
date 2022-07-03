package main

import "fmt"

// Gibt Ergebnisse von Power2() auf die Konsole aus.
// Der Kommentar unten gibt die erwarteten Ergebnisse an.
// Automatische Prüfung mittels des Befehls "go test" (statt "go run").
func ExamplePower2() {

	fmt.Println(Power2(2))
	fmt.Println(Power2(5))
	fmt.Println(Power2(0))
	fmt.Println(Power2(-2))

	// Output:
	// 4
	// 32
	// 1
	// 0.25
}

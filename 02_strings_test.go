package main

import "fmt"

// Gibt Ergebnisse von Separate() auf die Konsole aus.
// Der Kommentar unten gibt die erwarteten Ergebnisse an.
// Automatische Prüfung mittels des Befehls "go test" (statt "go run").
func ExampleSeparate() {

	fmt.Println(Separate("bce", "abcdef"))
	fmt.Println(Separate("abc", "adbecf"))
	fmt.Println(Separate("abc", "dabcef"))

	// Output:
	// adf
	// def
	// def
}

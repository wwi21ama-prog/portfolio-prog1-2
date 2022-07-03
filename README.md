# Prüfungsaufgaben zu Programmierung 1

[![Run on Repl.it](https://repl.it/badge/github/wwi21ama-prog/portfolio_prog1_2)](https://repl.it/github/wwi21ama-prog/portfolio_prog1_2)

Dieses Repo enthält Prüfungsaufgaben für die Vorlesung "Programmierung 1".

## Aufbau der Aufgaben

Zu jeder Aufgabe (außer der letzten) gibt es zwei Go-Dateien:  
In einer Datei wird die Aufgabe formuliert, in der anderen stehen Tests für die Aufgabe.  
Die Test-Datei heißt genau wie die Aufgaben-Datei mit einem zusätzlich angehängten "_test".  
Beispiel: `01_arrays.go` enthält die erste Aufgabe zu Arrays, `01_arrays_test.go` enthält
den zugehörigen Testcode.

Implementieren Sie Ihre Lösung in der Aufgaben-Datei und nutzen Sie die Funktion(en) aus der Testdatei,
um Ihre Lösung zu testen.

* Sie können die Tests automatisch ausführen lassen, indem Sie auf der Konsole den Befehl `go test` verwenden.  
  Dies führt alle Funktionen aus Test-Dateien aus, die mit `Example_` beginnen, und vergleicht deren Ausgabe  
  mit dem, was in der jeweiligen Funktion unten unter der Zeile `// Output:` steht.
* Um nur die Tests für eine einzelne Aufgabe auszuführen, verwenden Sie z.B. `go test -run 01`.  
  Hinter dem `-run` kann ein Test-Name oder ein Teil davon stehen, es werden dann alle Tests ausgeführt, die dazu passen.
* Sie können die Tests natürlich auch manuell in der `main()`-Funktion in `main.go` ausführen.  
  Die `main()`-Funktion ist bisher leer, kann baer für eigene Tests und ggf. für Aufgabe 6 oder 7 verwendet werden.

### Bearbeiten des Repos auf replit.com

* Der Badge oben auf dieser Seite öffnet das Repo auf replit.com.  
  Die Replit-Konfiguration ist so eingestellt, dass beim Klick auf den Run-Button alle Tests ausgeführt werden.  
  Um einzelne Tests oder die `main.go` auszuführen, ändern Sie die run-Zeile in .replit oder benutzen Sie die Konsole.

## Prüfungsmodalitäten

### Abgabe

* **Termin:** 25.07.2022

### Prüfungsgespräche

* Prüfungsgespräch mit jeder/jedem erforderlich.
  * Sollten bis zur Abgabe stattgefunden haben.
  * Ziele: Prüfung der Eigenständigkeit und Feedback/Hilfestellung
* Aufgaben müssen bis dahin nicht endgültig erledigt sein, es kann auch über Teillösungen geredet werden.
* Prüfungsgespräche dienen nicht primär der Notenfindung, können aber herangezogen werden.

## Bewertungskriterien

Für die Aufgaben 1-4 werden in erster Linie die Tests ausschlaggebend sein.  
D.h. die Tests müssen laufen, die Signaturen der Funktionen sollen auch nicht verändert werden.
Neben den Tests aus diesem Repo wird es weitere Tests für jede der Funktionen geben,  
mit denen andere Werte, Sonderfälle etc. geprüft werden.

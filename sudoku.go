// Wir bauen einen GoLang Sudoku Server. Hauptprogramm.

package main

import (
	"fmt"
	"math"
)

func main() {

	su := make([][]int, 9)

	beginSudo(su)
	printSudo(su)

	solveSudo(su)

	printSudo(su)
	fmt.Println("energy:", energySudo(su))
	checkSudo(su)
}

// Beginn: Die Sudoku-Aufgabe wird hier kodiert.
// arr ist das 9x9 Sudoku-Zahlenquadrat.
func beginSudo(arr [][]int) {
	arr[0] = []int{6, 0, 0, 4, 0, 0, 0, 0, 5}
	arr[1] = []int{2, 9, 4, 6, 0, 0, 7, 0, 0}
	arr[2] = []int{0, 0, 0, 0, 0, 0, 4, 6, 1}

	arr[3] = []int{0, 2, 3, 5, 0, 0, 0, 7, 4}
	arr[4] = []int{0, 0, 0, 0, 9, 0, 0, 0, 0}
	arr[5] = []int{8, 7, 0, 0, 0, 6, 5, 3, 0}

	arr[6] = []int{7, 6, 9, 0, 0, 0, 0, 0, 0}
	arr[7] = []int{0, 0, 5, 0, 0, 9, 3, 2, 7}
	arr[8] = []int{1, 0, 0, 0, 0, 7, 0, 0, 6}
}

func energySudo(arr [][]int) (energy float64) {
	energy = 0
	count := 0
	for i := 0; i < 9; i++ {
		count = 0
		for j := 0; j < 9; j++ {
			count += arr[i][j]
		}
		energy += math.Abs(float64(45 - count))
	}
	for i := 0; i < 9; i++ {
		count = 0
		for j := 0; j < 9; j++ {
			count += arr[j][i]
		}
		energy += math.Abs(float64(45 - count))
	}
	return
}

// Drucken des Sudoku-Zahlenquadrates
func printSudo(arr [][]int) {
	fmt.Println("---------------------")
	for i := 0; i < 9; i++ {
		fmt.Println(i, arr[i])
	}
	fmt.Println("---------------------")
}

// Versuche, das Sudoku-Rätsel zu lösen.
// Momentan passiert hier noch gar nichts.
func solveSudo(arr [][]int) {
	arr[0] = []int{6, 1, 7, 4, 3, 8, 2, 9, 5}
	arr[1] = []int{2, 9, 4, 6, 1, 5, 7, 8, 3}
	arr[2] = []int{3, 5, 8, 9, 7, 2, 4, 6, 1}

	arr[3] = []int{9, 2, 3, 5, 8, 1, 6, 7, 4}
	arr[4] = []int{5, 4, 6, 7, 9, 3, 8, 1, 2}
	arr[5] = []int{8, 7, 1, 2, 4, 6, 5, 3, 9}

	arr[6] = []int{7, 6, 9, 3, 2, 4, 1, 5, 8}
	arr[7] = []int{4, 8, 5, 1, 6, 9, 3, 2, 7}
	arr[8] = []int{1, 3, 2, 8, 5, 7, 9, 4, 6}
}

//Testet, ob der Sudoku-array OK ist.
//Die momentane Funktion ist noch nicht OK und muss wesentlich verbessert werden.
func checkSudo(arr [][]int) (ok bool) {
	var count = 0
	ok = true
	for i := 0; i < 9; i++ {
		count = 0
		for j := 0; j < 9; j++ {
			count += arr[i][j]
		}
		if count != 45 {
			fmt.Println("checkSudo: Not OK: line error", count)
			ok = false
			return
		}
	}

	for i := 0; i < 9; i++ {
		count = 0
		for j := 0; j < 9; j++ {
			count += arr[j][i]
		}
		if count != 45 {
			fmt.Println("checkSudo: Not OK: column error", count)
			return
		}
	}
	fmt.Println("checkSudo: OK")
	return
}

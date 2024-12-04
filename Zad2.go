package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	///"testing"
)

func detinc(A []int, B []int) bool {
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			return false
		}
	}
	return true
}

func detdec(A []int, B []int) bool {
	for i := 0; i < len(A); i++ {
		if A[i] != B[len(B)-1-i] {
			return false
		}
	}
	return true
}

func absval(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func determine(list []int) bool {
	sor := append([]int(nil), list...)
	sort.Ints(sor)
	guessinc := detinc(list, sor)
	guessdec := detdec(list, sor)

	if guessinc || guessdec {
		return true
	}
	return false
}
func determineif(list []int) int {
	sor := append([]int(nil), list...)
	sort.Ints(sor)
	guessinc := detinc(list, sor)
	guessdec := detdec(list, sor)
	if guessinc == true {
		return 1
	}
	if guessdec == true {
		return 2
	}
	return 0
}
func checkout(list []int) bool {
	if determine(list) {
		for i := 0; i < len(list)-1; i++ {
			ror := list[i] - list[i+1]
			if ror == 0 || absval(ror) > 3 {
				return false
			}
		}
		return true
	}
	return false
}
func listconv(list []string) []int {
	result := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		number, err := strconv.Atoi(list[i])
		if err != nil {
			fmt.Println("Błąd konwersji:", err)
			return nil
		}
		result[i] = number
	}
	return result
}

func repair(list []int) bool {
	safe := false
	for i := 0; i < len(list); i++ {
		copy := slices.Clone(list)
		copy = slices.Delete(copy, i, i+1)
		if checkout(copy) {
			safe = true
			break
		}
	}

	return safe
}
func main() {
	file, err := os.Open("DATA2.txt")
	if err != nil {
		fmt.Println("erorer", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	///scanner.Split(bufio.ScanWords)
	wynik := 0
	nova := 0
	for scanner.Scan() {
		lista := strings.Split(scanner.Text(), " ")
		line := listconv(lista)
		if checkout(line) == true {
			wynik++
		}
		if checkout(line) == false {
			red := repair(line)
			if red == true {
				nova++
			}

		}

	}

	fmt.Println("wynik", wynik)
	fmt.Println("nova", nova)
	fmt.Println(nova + wynik)
}

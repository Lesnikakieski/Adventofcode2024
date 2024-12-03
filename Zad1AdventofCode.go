package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func in(slice []int, value int) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			return true
		}
	}
	return false
}

func unique(slice []int) []int {
	uniques := []int{}
	for i := 0; i < len(slice); i++ {
		if !in(uniques, slice[i]) { // Sprawdzaj w `uniques`, nie w `slice`
			uniques = append(uniques, slice[i])
		}
	}
	return uniques
}

func main() {
	file, err := os.Open("Data.txt")
	if err != nil {
		fmt.Println("fatal Error!", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	lista := [][]int{
		{},
		{},
	}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error!", err)
			continue
		}
		if i%2 == 0 {
			lista[0] = append(lista[0], number)
		}
		if i%2 == 1 {
			lista[1] = append(lista[1], number)
		}
		i = i + 1
	}
	sort.Ints(lista[0])
	sort.Ints(lista[1])

	fmt.Println(lista[0])
	fmt.Println(lista[1])

	fmt.Println(len(lista[0]))
	fmt.Println(len(lista[1]))

	finlenght := 0
	for i := 0; i < len(lista[0]); i++ {
		temp := lista[0][i] - lista[1][i]
		if temp < 0 {
			temp = -temp

		}
		finlenght = finlenght + temp
	}
	//fmt.Println(finlenght)

	finpart2 := 0
	dodwa := unique(lista[0])
	fmt.Println(len(dodwa))
	fmt.Println(dodwa)
	for i := 0; i < len(dodwa); i++ {
		licznik := 0
		for j := 0; j < len(lista[1]); j++ {
			if dodwa[i] == lista[1][j] {
				licznik++
			}
		}
		finpart2 += dodwa[i] * licznik
	}
	fmt.Println(finpart2)
}

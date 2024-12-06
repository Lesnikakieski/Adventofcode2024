package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func mul(a int, b int) int {
	return a * b

}
func files(s string) *os.File {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err.Error())
	}
	return file
}
func main() {
	file, err := os.ReadFile("Data3.txt")
	if err != nil {
		print("error reading file", err)
		return
	}
	mule := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))`)
	allmule := mule.FindAllString(string(file), -1)

	all := 0
	for _, match := range allmule {
		nums := strings.Split(match[4:len(match)-1], ",")
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			print("error converting number", err)
			return
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			print("error converting number", err)
			return
		}
		all += mul(num1, num2)
	}
	fmt.Println(all)

}

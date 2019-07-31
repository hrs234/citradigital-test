package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	questA("omama")
}

func sameAlphabets(inputs string) string {
    newS := ""
    taken := make(map[rune]int)
    for _, value := range inputs {

		if vowellValid(string(value)) == 1 {

			if _, ok := taken[value]; !ok {
				taken[value] = 1
				newS += string(value)
			}
		} else {
			taken[value] = 1
			newS += string(value)
		}
		
	}
    
    // fmt.Println(newS)
	return newS
}

func vowellValid(inputs string) int {
	vowels := []string{"a", "i", "u", "e", "o"}
	var result int = 0
	for _, n := range vowels {
		if inputs == n {
			result = 1
		}
	}
	return result
}

func questA(inputs string) {


	// split data into per alphabets
	splited := strings.Split(sameAlphabets(inputs), "")

	var countVowell int = 0
	var countnonVowell int = 0

	for kata := range splited {
		if vowellValid(splited[kata]) == 1 {
			countVowell++
		} else {
			countnonVowell++
		}
		
	}

	converterA := strconv.Itoa(countVowell)
	converterB := strconv.Itoa(countnonVowell)
	
	fmt.Println("Jumlah Huruf Hidup", converterA)
	fmt.Println("Jumlah Huruf Mati", converterB)

}

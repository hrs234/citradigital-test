package main 

import (
    "fmt"
    "sort"
    "strings"
)

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

func selector(inputs string) {

    var mySliceA []string
    var mySliceB []string
    var mySliceC []string

    split := strings.Split(inputs, "")


    for kata := range split {
		if vowellValid(split[kata]) == 1 {
			mySliceA = append(mySliceA, split[kata])
		} else {
			mySliceB = append(mySliceB, split[kata])
		}
		
	}

    sort.Strings(mySliceA)
    sort.Strings(mySliceB)


    mySliceC = append(mySliceA, mySliceB...)

    
    fmt.Println(strings.Join(mySliceC[:], ""))

}

func main () {

    selector("osama")
}
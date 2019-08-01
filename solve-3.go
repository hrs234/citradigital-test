package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	_ "strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Roll struct {
	Hasil string `json:"hasil"`
	// ImageNumber string `json:"image_number"`
	// Name string `json:"id"`
	// Ingredients string `json:"name"`
}

var rolls []Roll

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

func selector(inputs string) string {

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

	// fmt.Println(strings.Join(mySliceC[:], ""))
	return strings.Join(mySliceC[:], "")

}

func getRolls(w http.ResponseWriter, r *http.Request) {
	// w.Header().set("Content-Type", "application/json")

	keys, ok := r.URL.Query()["word"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Param not inserted!!")
		return
	}

	key := keys[0]

	rolls = append(rolls, Roll{Hasil: selector(string(key))})
	log.Println(rolls)
    json.NewEncoder(w).Encode(rolls)
    rolls = nil
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/welcome", getRolls).Methods("GET")

	log.Printf("Started...")
	log.Fatal(http.ListenAndServe(":5000", router))
}

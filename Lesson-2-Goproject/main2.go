package main

import "fmt"

func main() {

	var word string = "Hello My Darling"
	fmt.Println(word)
	fmt.Println(word)
	fmt.Println(word)
	fmt.Println(word)
	fmt.Println(word)
	fmt.Println(word)

	var bill int = 42

	fmt.Println(100 + (100 * bill / 100))

	var bill2 float32 = 13.3
	fmt.Println(100 + 100*bill2/100)

	bill3 := 12

	fmt.Printf("Type: %T\n", bill3, bill2, bill)

	var words bool

	var word1 string = "Fabia"

	var word2 string = "Boa Hancoock"

	words = word1 == word2

	fmt.Println(words)

}

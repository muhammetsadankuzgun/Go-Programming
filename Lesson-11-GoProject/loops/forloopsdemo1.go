package loops

import "fmt"

func Demo1000() {

	var word string = "Hello My Darling"

	i := 1
	for i <= 7 {

		fmt.Println(word)
		i = i + 1
	}
	fmt.Println("Finished")
}

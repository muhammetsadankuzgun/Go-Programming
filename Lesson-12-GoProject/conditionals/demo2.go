package conditionals

import "fmt"

func Demo2() {

	var money float64 = 2000
	var want float64 = 2000

	if want > money {
		fmt.Println("invalid")

	} else if want == money {
		fmt.Println("Ready")
		fmt.Println("Your money can be finished")
	} else {
		fmt.Println("Ready")
	}
}

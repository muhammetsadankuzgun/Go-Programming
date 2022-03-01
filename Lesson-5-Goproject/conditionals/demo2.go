package conditionals

import "fmt"

func Demo2() {

	var money float64 = 2000
	var want float64 = 200

	if want > money {
		fmt.Println("invalid")

	} else {
		fmt.Println("Ready")

	}
}

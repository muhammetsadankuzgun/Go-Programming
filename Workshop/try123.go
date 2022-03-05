package loops

import "fmt"

func Demo4() {

	ages := map[string]int{
		"lane":    26,
		"preston": 28,
		"rory":    21,
	}
	for name, age := range ages {
		fmt.Println(name, age)
	}

	// prints:
	// lane 26
	// preston 28
	// rory 21
}

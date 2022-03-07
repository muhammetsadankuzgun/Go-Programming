package loops

import "fmt"

func Demo5() {

	n1 := 220
	n2 := 284

	total1 := 0

	total2 := 0

	for i := 1; i < n1; i++ {
		if n1%i == 0 {
			total1 = total1 + i
		}
	}

	for i := 1; i < n2; i++ {
		if n2%i == 0 {
			total2 = total2 + i
		}
	}
	if total1 == n2 && total2 == n1 {
		fmt.Printf("%v and %v Friendship Numbers", n1, n2)
	} else {
		fmt.Printf("%v and %v Not Friendship Numbers", n1, n2)
	}
}

package loops

import (
	"fmt"
)

func Demo3() {
	thenumberisinmymind := 52
	guessnumber := 0

	fmt.Println("Plz Guess a Number:")
	fmt.Scanln(&guessnumber)

	for thenumberisinmymind != guessnumber {

		if thenumberisinmymind > guessnumber {
			fmt.Println("Tell me the biggest number")
			fmt.Scanln(&guessnumber)
		}
		if thenumberisinmymind < guessnumber {
			fmt.Println("Tell me the smallest number")
			fmt.Scanln(&guessnumber)
		}
	}
	if thenumberisinmymind == guessnumber {
		fmt.Println("Yes It is a true number")
		fmt.Scanln(&guessnumber)
	}

}

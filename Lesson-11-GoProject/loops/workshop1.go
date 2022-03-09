package loops

import (
	"fmt"
)

func Demo3() {
	thenumberisinmymind := 52
	guessnumber := 0
	guessnumbers := 0

	fmt.Println("Plz Guess a Number:")
	fmt.Scanln(&guessnumber)

	for thenumberisinmymind != guessnumber {

		if thenumberisinmymind > guessnumber {
			fmt.Println("Tell me the biggest number")
			fmt.Scanln(&guessnumber)
			guessnumbers = guessnumbers + 1
		}
		if thenumberisinmymind < guessnumber {
			fmt.Println("Tell me the smallest number")
			fmt.Scanln(&guessnumber)
			guessnumbers = guessnumbers + 1

		}
	}
	if thenumberisinmymind == guessnumber {
		fmt.Println("Yes It is a true number")
		fmt.Scanln(&guessnumber)
		guessnumbers = guessnumbers + 1
	}

	successstatus := ""
	if guessnumbers > 0 && guessnumbers <= 3 {
		successstatus = "Perfect"
	} else if guessnumbers <= 6 {

		successstatus = "Good"

	} else {
		successstatus = "So So"
	}

	fmt.Printf("Bravo It's true %v guess: %v", guessnumbers, successstatus)
}

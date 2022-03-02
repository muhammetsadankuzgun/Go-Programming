package conditionals

import "fmt"

func Workshop1() {
	var n1, n2, n3 int = 10, 33, 48

	var thebiggest int = n1

	if n2 > thebiggest {

		thebiggest = n2
	}
	if n3 > thebiggest {
		thebiggest = n3
	}
	fmt.Printf("The Biggest Number is: %v", thebiggest)
}

package arrays

import "fmt"

func Demo3() {
	numbers := [7]int{10, 20, 95, 88, 9, 13, 1}

	fmt.Println(numbers)
	thebiggest := numbers[0]

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])

		if numbers[i] > thebiggest {
			thebiggest = numbers[i]

		}

	}
	fmt.Println("The Biggest Number:", thebiggest)
}

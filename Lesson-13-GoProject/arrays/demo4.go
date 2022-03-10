package arrays

import "fmt"

func Demo4() {

	var numbers [2][3]int
	numbers[0][0] = 5
	numbers[0][1] = 4
	numbers[0][2] = 6
	numbers[1][0] = 5
	numbers[1][1] = 9
	numbers[1][2] = 7

	//fmt.Println(numbers[0][1])
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(numbers[i][j])

			fmt.Print(" ")
		}
		fmt.Println()
	}
}

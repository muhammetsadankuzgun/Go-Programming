package conditionals

import "fmt"

func Demo1() {

	var money float64 = 2000
	var want float64 = 200

	if want > money {
		fmt.Println("invalid")

	}

	if want <= money {

		fmt.Println("Ready")
		money = money - want
	}

	fmt.Println("Finish. Final situation:" + fmt.Sprintf("%v", money))

}

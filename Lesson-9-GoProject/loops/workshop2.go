package loops

import "fmt"

func Demo4() {

	number := 0
	fmt.Println("Press a Number:")
	fmt.Scanln(&number)

	asalMı := true

	for i := 2; i < number; i++ {

		if number%i == 0 {
			asalMı = false

		}
	}
	if asalMı == true {
		fmt.Println("Asaldır")

	} else {
		fmt.Println("Asal değildir")
	}
}

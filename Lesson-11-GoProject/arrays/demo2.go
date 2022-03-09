package arrays

import "fmt"

func Demo2() {

	var cities [3]string
	cities[0] = "Istanbul"
	cities[1] = "Berlin"
	cities[2] = "Paris"
	fmt.Println(cities)

	for i := 0; i < 3; i++ {

		fmt.Println(cities[i])

	}
}

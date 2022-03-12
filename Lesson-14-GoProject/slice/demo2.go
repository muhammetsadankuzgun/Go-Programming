package slice

import "fmt"

func Demo2() {

	cities := []string{"Ankara,Istanbul,Adana"}
	fmt.Println(cities)
	citiescopy := make([]string, len(cities))
	fmt.Println(citiescopy)

	copy(citiescopy, cities)
	fmt.Println(citiescopy)

	cities = append(cities, "Karabuk")

	fmt.Println(cities)
	fmt.Println(citiescopy)
	fmt.Println(cities[1:2])
	fmt.Println(cities[:2])
	fmt.Println(cities[1:])

}

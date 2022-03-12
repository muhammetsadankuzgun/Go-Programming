package slice

import "fmt"

func Demo1() {

	names := make([]string, 3)

	//fmt.Println(names)

	names[0] = "Muhammet"
	names[1] = "Şadan"
	names[2] = "Kuzgun"
	names = append(names, "<3 Vanessa Kessa")
	fmt.Println(names)
	fmt.Println(len(names))

}

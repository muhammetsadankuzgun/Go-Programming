package functions

func AddVariadic(numbers ...int) int {

	add := 0
	for i := 0; i < len(numbers); i++ {
		add = add + numbers[i]
	}
	return add

}

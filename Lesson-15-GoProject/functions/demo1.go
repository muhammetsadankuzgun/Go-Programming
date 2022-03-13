package functions

func Fourcalculations(n1 int, n2 int) (int, int, int, float64) {
	add := n1 + n2
	subtraction := n1 - n2
	multiplication := n1 * n2
	divide := float64(n1 / n2)

	return add, subtraction, multiplication, divide
}

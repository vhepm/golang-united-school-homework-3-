package homework

func reverse(input []int64) (result []int64) {
	reversed := make([]int64, len(input))
	copy(reversed, input)

	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = input[j], input[i]
	}

	return reversed
}

package homework

func average(input [15]float32) (result float32) {
	var arraySum float32

	for _, f := range input {
		arraySum += f
	}
	return arraySum / float32(len(input))
}

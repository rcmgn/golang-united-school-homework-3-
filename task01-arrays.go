package homework

func average(input [15]float32) (result float32) {
	for _, i := range input {
		result += i
	}
	result = result / float32(len(input))
	return result
}

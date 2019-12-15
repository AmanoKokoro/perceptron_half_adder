package perceptron

//Data is struct weight and bias
type Data struct {
	w1   float64
	w2   float64
	bias float64
}

func perceptronxor(x1, x2 float64) float64 {
	var num Data
	var a float64 = 0
	num.w1, num.w2, num.bias = initialize()

	a = x1*num.w1 + x2*num.w2

	if a > num.bias {
		return 1
	}

	return 0
}

//Xor func is Return answerXor
func Xor(a, o, y float64) float64 {

	a = perceptronxor(a, o)

	if a == y {
		return a
	}

	return 0
}

//Initialize return weight and bias
func initialize() (float64, float64, float64) {

	return 0.3, 032, 0.5
}

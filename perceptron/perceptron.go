package perceptron

//Perceptron func return 1.0 or 0.0
func Perceptron(x1, x2, w1, w2, bias float64) int {
	var a float64 = 0

	a = x1*w1 + x2*w2

	if a > bias {
		return 1
	}
	return 0
}

//Dot is learning weight nd bias
func dot(vec0, vec1 []float64) float64 {
	tmp := 0.0
	for num := range vec0 {
		tmp += vec0[num] * vec1[num]
	}

	return tmp
}

func step(num float64) float64 {
	if num > 0 {
		return 1
	}
	return 0
}

//Feedforward is output func
func Feedforward(i, w []float64) float64 {
	return step(dot(i, w))
}

//Train is incremental learning func
func Train(w, i []float64, y, eta float64) []float64 {
	o := Feedforward(i, w)

	for wrange := range w {
		w[wrange] = w[wrange] + (y-o)*i[wrange]*eta
	}

	return w
}

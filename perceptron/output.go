package perceptron

import (
	"fmt"
	"math/rand"
	"time"
)

//W1 is Weight1 W2 is Weight2 Bias is bias
var (
	W1   float64
	W2   float64
	Bias float64
)

//Initialize func
func Initialize() {
	rand.Seed(time.Now().UnixNano())

	W1 = rand.Float64()*(0.0-1.0) + 1.0
	W2 = rand.Float64()*(0.0-1.0) + 1.0
	Bias := rand.Float64()*(0.0-1.0) + 1.0

	fmt.Printf("1 = %f\n", W1)
	fmt.Printf("2 = %f\n", W2)
	fmt.Printf("b = %f\n", Bias)
}

//Setans func
func Setans(y []int) {
	y[0] = Perceptron(0, 0, W1, W2, Bias)
	y[1] = Perceptron(0, 1, W1, W2, Bias)
	y[2] = Perceptron(1, 0, W1, W2, Bias)
	y[3] = Perceptron(1, 1, W1, W2, Bias)

	for i := 0; i < 4; i++ {
		fmt.Printf("y[%d] = %d\n", i, y[i])
	}
}

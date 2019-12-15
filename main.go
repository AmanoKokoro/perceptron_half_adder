package main

import (
	"./perceptron"
	"fmt"
	"math"
	"time"
)

func main() {
	x := [][]float64{{0, 0, 1}, {0, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	//input := [][]float64{{0, 0, 0, 1}, {0, 0, 1, 1}, {0, 1, 0, 1}, {0, 1, 1, 1}, {1, 0, 0, 1}, {1, 0, 1, 1}, {1, 1, 0, 1}, {1, 1, 1, 1}}
	yAND := []float64{0, 0, 0, 1} //鋸歯データ and
	yOR := []float64{0, 1, 1, 1}  //鋸歯データ or
	yNAND := []float64{1, 1, 1, 0}
	yXOR := []float64{0, 1, 1, 0}

	weightAND := []float64{0, 0, 0} //重み
	weightOR := []float64{0, 0, 0}  //重み
	weightNAND := []float64{0, 0, 0}
	//weightXOR := []float64{0, 0, 0}

	answerAND := []float64{0, 0, 0, 0}
	answerOR := []float64{0, 0, 0, 0}
	answerNAND := []float64{0, 0, 0, 0}
	answerXOR := []float64{0, 0, 0, 0}
	eta := 0.1 //学習率

	//AND, OR, NAND learning
	for cnt := 0; cnt < math.MaxUint16; cnt++ {
		for j := range yAND {
			weightAND = perceptron.Train(weightAND, x[j], yAND[j], eta)
			weightOR = perceptron.Train(weightOR, x[j], yOR[j], eta)
			weightNAND = perceptron.Train(weightNAND, x[j], yNAND[j], eta)
		}
	}

	fmt.Printf("Learning.")
	for cnt := 0; cnt < 20; cnt++ {
		fmt.Printf(".")
		time.Sleep(1 * time.Second) // 3秒休む
	}

	fmt.Println("\nC       , S       ")
	for i, j := range x {
		answerAND[i] = perceptron.Feedforward(j, weightAND)
		answerOR[i] = perceptron.Feedforward(j, weightOR)
		answerNAND[i] = perceptron.Feedforward(j, weightNAND)
		answerXOR[i] = perceptron.Xor(answerOR[i], answerNAND[i], yXOR[i])

		fmt.Printf("%f, %f\n", answerAND[i], answerXOR[i])
	}
}

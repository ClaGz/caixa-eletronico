package main

import (
	"caixaeletronico/models"
	"fmt"
)

func main() {
	fmt.Println(saque(0.00))
	fmt.Println(AndaArray([]int{1, 2, 3, 4}))
}

// TODO: aprender a usar o DEBUG!!!
func saque(quantia float64) []string {
	caixa := models.GetCaixa()

	return caixa.GetNotasFromValue(quantia)
}

func AndaArray(array []int) []int {
	hold := array[len(array)-1]

	for i := len(array) - 1; i > 0; i-- {
		//hold = array[i]
		fmt.Println(array[i])
		fmt.Println(array[i-1])

		array[i] = array[i-1]
	}

	array[0] = hold

	//fmt.Println()
	return array
}

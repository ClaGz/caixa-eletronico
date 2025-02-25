package models

import (
	"sort"
)

type caixa struct {
	notas map[float64]string
}

func GetCaixa() *caixa {
	return &caixa{
		notas: map[float64]string{
			10.00:  "10",
			20.00:  "20",
			50.00:  "50",
			100.00: "100",
		},
	}
}

func (c *caixa) GetNotasFromValue(value float64) []string {
	response := []string{}

	if c.notas[value] == "" {
		faltaAinda := 0.00
		notasCaixaArr := c.sortArray(c.toArray())

		for i := len(notasCaixaArr) - 1; i >= 0; i-- {

			faltaAinda = value - notasCaixaArr[i]

			if faltaAinda >= 0 {
				response = append(response, c.notas[notasCaixaArr[i]])
				value -= notasCaixaArr[i]

				if faltaAinda >= notasCaixaArr[i] {
					i++
				}
			}
		}

	} else {
		response = append(response, c.notas[value])

	}

	return response
}

func (c *caixa) toArray() []float64 {
	keys := make([]float64, len(c.notas))

	i := 0
	for k := range c.notas {
		keys[i] = k
		i++
	}

	return keys
}

func (c *caixa) sortArray(notasArray []float64) []float64 {
	response := notasArray

	sort.Slice(response, func(i, j int) bool {
		return response[i] < response[j]
	})

	return response
}

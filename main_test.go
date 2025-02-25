package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("Quando o saque de 10.00 então devemos retornar uma nota de 10 reais apenas", func(t *testing.T) {
		expected := []string{"10"}
		saque := saque(10.00)

		fmt.Println("SAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEE : ", saque)

		if expected[0] != saque[0] {
			t.Errorf("saque() = %v, want %v", saque, expected)
		}
	})

	t.Run("Quando o saque de 30.00 então devemos retornar uma nota de 10 e uma de 20 reais", func(t *testing.T) {
		expected := []string{"20", "10"}
		saque := saque(30.00)

		fmt.Println("SAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEE : ", saque)

		if expected[0] != saque[0] && expected[1] != saque[1] {
			t.Errorf("saque() = %v, want %v", saque, expected)
		}
	})

	t.Run("Quando o saque de 40.00 então devemos retornar uma nota de 20 e uma de 20 reais", func(t *testing.T) {
		expected := []string{"20", "20"}
		saque := saque(40.00)

		fmt.Println("SAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEE : ", saque)

		if expected[0] != saque[0] && expected[1] != saque[1] {
			t.Errorf("saque() = %v, want %v", saque, expected)
		}
	})

	t.Run("Quando o saque de 90.00 então devemos retornar uma nota de 50 e duas de 20 reais", func(t *testing.T) {
		expected := []string{"50", "20", "20"}
		saque := saque(90.00)

		fmt.Println("SAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEE : ", saque)

		for i, it := range saque {
			if it != expected[i] {
				t.Errorf("saque() = %v, want %v", it, expected[i])
			}
		}
	})

	t.Run("Quando o saque de 280.00 então devemos retornar 2 notas de 100, uma nota de 50, uma de 20 e uma de 10 reais", func(t *testing.T) {
		expected := []string{"100", "100", "50", "20", "10"}
		saque := saque(280.00)

		fmt.Println("SAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEESAQUEEE : ", saque)

		for i, it := range saque {
			if it != expected[i] {
				t.Errorf("saque() = %v, want %v", it, expected[i])
			}
		}
	})
}

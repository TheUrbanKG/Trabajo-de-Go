package main

import (
	"fmt"
	"strconv"
)

func calculate(expression string) (float64, error) {
	num1, operador, num2, err := parseExpression(expression)
	if err != nil {
		return 0, err
	}

	switch operador {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("bro, no puedes dividir por 0")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("ha elegido un operador invalido")
	}
}

func parseExpression(expression string) (float64, string, float64, error) {
	var num1, num2 float64
	var operador string

	fmt.Sscanf(expression, "%f %s %f", &num1, &operador, &num2)

	return num1, operador, num2, nil
}

func main() {
	fmt.Println("Calculadora de Kevo")
	fmt.Println("----------------")

	for {
		fmt.Print("Introduce una expresion matematica (ej 2+2): ")
		var expression string
		fmt.Scanln(&expression)

		result, err := calculate(expression)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Resultado: %f\n", result)
		}
	}
}
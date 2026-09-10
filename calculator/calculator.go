package calculator

import (
	"errors"
)

// Барлық функциялар БАС ӘРІППЕН басталады (Exported)
func Add(num1, num2 int) int {
	return num1 + num2
}

func Subtract(num1, num2 int) int {
	return num1 - num2
}

func Multiply(num1, num2 int) int {
	return num1 * num2
}

func Divide(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("нөлге бөлуге болмайды")
	}
	return num1 / num2, nil
}

// power емес, Power деп БАС ӘРІППЕН жазамыз!
func Power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("теріс дәрежеге шығаруға болмайды")
	} else if exponent == 0 {
		return 1, nil
	}

	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}

	return result, nil
}

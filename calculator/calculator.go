package calculator

import (
	"errors"
	"fmt"
)

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
		return 0, errors.New("Nolge boluge balmaidi")
	}
	return num1 / num2, nil
}
func power(base, exponent int) (int, error) {
	if exponent < 0 {
		return 0, errors.New("Теріс дәрежеге бөлуге болмайды")
	} else if exponent == 0 {
		return 1, nil
	}
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}

	return result, nil
}

func main() {
	for {
		var num1, num2 int
		var amal string
		fmt.Println("This is calculator")

		fmt.Println("give me number num1")
		_, err := fmt.Scan(&num1)
		if num1 == 99 {
			fmt.Println("Congradulations your won")
			break
		}
		if err != nil {
			fmt.Println("The wrong number")
			return
		}
		fmt.Print("Амалды таңдаңыз (+, -, *, /, ^): ")
		fmt.Scan(&amal)

		fmt.Print("Екінші санды енгізіңіз: ")
		_, err = fmt.Scan(&num2)
		if err != nil {
			fmt.Println("Қате: тек сан енгізіңіз!")
			return
		}
		switch amal {
		case "+":
			fmt.Println("Kosu", Add(num1, num2))
		case "-":
			fmt.Println("azaitu", Subtract(num1, num2))
		case "*":
			fmt.Println("Kobeitu", Multiply(num1, num2))
		case "/":
			bolu, err := Divide(num1, num2)
			if err != nil {
				fmt.Println("kate", err)
			} else {
				fmt.Println("Bolu", bolu)
			}
		default:
			fmt.Println("Kate not found")
		case "^":
			powResult, err := power(num1, num2)
			if err != nil {
				fmt.Println("Қате:", err)
			} else {
				fmt.Println("Дәреже нәтижесі:", powResult)
			}

		}
	}

}

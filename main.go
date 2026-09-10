package main

import (
	"fmt"

	"github.com/code-makersb/go-calculator/calculator"
)

func main() {
	for {
		var num1, num2 int
		var amal string
		fmt.Println("\n--- Калькулятор іске қосылды ---")

		fmt.Print("Бірінші санды енгіз (шығу үшін 99 жаз): ")
		_, err := fmt.Scan(&num1)
		if num1 == 99 {
			fmt.Println("Бағдарлама аяқталды!")
			break
		}
		if err != nil {
			fmt.Println("Қате: тек сан енгізіңіз!")
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

		// Функцияларды 'calculator.' арқылы шақырамыз!
		switch amal {
		case "+":
			fmt.Println("Қосу нәтижесі:", calculator.Add(num1, num2))
		case "-":
			fmt.Println("Азайту нәтижесі:", calculator.Subtract(num1, num2))
		case "*":
			fmt.Println("Көбейту нәтижесі:", calculator.Multiply(num1, num2))
		case "/":
			bolu, err := calculator.Divide(num1, num2)
			if err != nil {
				fmt.Println("Қате:", err)
			} else {
				fmt.Println("Бөлу нәтижесі:", bolu)
			}
		case "^":
			powResult, err := calculator.Power(num1, num2)
			if err != nil {
				fmt.Println("Қате:", err)
			} else {
				fmt.Println("Дәреже нәтижесі:", powResult)
			}
		default:
			fmt.Println("Қате: белгісіз амал енгізілді!")
		}
	}
}

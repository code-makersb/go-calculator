package main

import "fmt"

func main() {
	for {
		var num1 int
		var operator string
		var num2 int

		fmt.Println("\n--- Калькулятор іске қосылды ---")
		fmt.Print("Бірінші санды енгіз (шығу үшін 99 жаз): ")
		fmt.Scan(&num1)

		if num1 == 99 {
			fmt.Println("Бағдарлама аяқталды. Қош сау болыңыз!")
			break
		}

		fmt.Print("Математикалық амалды енгіз (+, -, *, /): ")
		fmt.Scan(&operator)

		fmt.Print("Екінші санды енгіз: ")
		fmt.Scan(&num2)

		switch operator {
		case "+":
			fmt.Println("Нәтиже (қосу):", num1+num2)
		case "-":
			fmt.Println("Нәтиже (азайту):", num1-num2)
		case "*":
			fmt.Println("Нәтиже (көбейту):", num1*num2)
		case "/":

			if num2 == 0 {
				fmt.Println("Қате: Санды нөлге бөлуге болмайды!")
				continue
			}
			fmt.Println("Нәтиже (бөлу):", num1/num2)
		default:

			fmt.Println("Қате: Белгісіз математикалық амал енгізілді!")
		}
	}
}

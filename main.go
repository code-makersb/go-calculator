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

		// Ескі if-else блоктарының орнына SWITCH-CASE қолданамыз
		switch operator {
		case "+":
			fmt.Println("Нәтиже (қосу):", num1+num2)
		case "-":
			fmt.Println("Нәтиже (азайту):", num1-num2)
		case "*":
			fmt.Println("Нәтиже (көбейту):", num1*num2)
		case "/":
			// Қауіпсіздік тексерісін case ішінде қалдырамыз
			if num2 == 0 {
				fmt.Println("Қате: Санды нөлге бөлуге болмайды!")
				continue // Бұл continue сырттағы for циклін басына қайтарады
			}
			fmt.Println("Нәтиже (бөлу):", num1/num2)
		default:
			// Пайдаланушы қате таңба енгізсе, осы блок іске қосылады
			fmt.Println("Қате: Белгісіз математикалық амал енгізілді!")
		}
	}
}

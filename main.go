package main

import "fmt"

func main() {
	for {
		fmt.Println("Dude its calculator")
		fmt.Println("give me number ")
		var a int
		fmt.Scan(&a)
		if a == 99 {
			fmt.Println("God damn, 99")
			break
		}
		fmt.Println("write +,-,*,/")
		var b string
		fmt.Scan(&b)

		fmt.Println("write number")
		var c int
		fmt.Scan(&c)

		if b == "+" {
			fmt.Println("kosu", a+c)
		} else if b == "-" {
			fmt.Println("minus", a-c)
		} else if b == "*" {
			fmt.Println("multiply", a*c)
		} else if b == "/" {
			if c == 0 {
				fmt.Println("god damn bro its wrong")
				continue
			}
			fmt.Println("bolu", a/c)

		}
	}
}

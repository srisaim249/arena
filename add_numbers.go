package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1, num2 := "", ""
	fmt.Println("this is the program to add 2 numbers")
	fmt.Printf("Enter a First Number ")
	fmt.Scanln(&num1)
	v1, err1 := strconv.ParseInt(num1, 10, 64)
	if err1 != nil {
		fmt.Println(num1+" is not a number, enter a valid number", err1)
		return
	}
	fmt.Printf("Enter a Second  Number ")
	fmt.Scanln(&num2)
	v2, err2 := strconv.ParseInt(num2, 10, 64)
	if err2 != nil {
		fmt.Println(num2+" is not a number, enter a valid number", err2)
		return
	}
	fmt.Println("The sum of ", v1, " and ", v2, " is ", v1+v2)
}

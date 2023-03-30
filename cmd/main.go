package main

import (
	"fmt"
	"module_34.6/pkg/calc"
)

func main() {

	str, err := calc.Calculate("./input.txt", "./output.txt")
	if err != nil {
		fmt.Println("Нет записей")
	}
	fmt.Println(string(str))

}

package l1packages

import (
	"fmt"
	"gocourse/l1packages/sum"

	"github.com/fatih/color"
)

func Main() {
	var number1, number2, result float64

	fmt.Println("Enter two numbers:")
	// fmt.Scanln(&number1, &number2)
	number1, number2 = 5, 6

	result = sum.Sum(number1, number2)

	c := color.New(color.FgCyan)
	c.Println(result)
}

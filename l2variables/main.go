package l2variables

import "fmt"

func Main() {
	var variable1 string = "Variable 1"
	variable2 := "Variable 2" // Type inference.

	fmt.Println(variable1, variable2)

	var (
		variable3 int = 3
		variable4 int = 4
	)

	variable3 = variable3 + variable4
	variable4 = variable3 - variable4
	variable3 = variable3 - variable4

	fmt.Println(variable4, variable3)

	const variable5 = "Variable 5"
	fmt.Println(variable5)
}

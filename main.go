package main

import (
	"calculator-yl/core"
)

func main() {
	// expr := "((2 * 9) - 3) - 5 / 2"
	// result, err := core.CalculateExpression(expr)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// } else {
	// 	fmt.Println("Result:", result)
	// }
	core.StartServer()
}

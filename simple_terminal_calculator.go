package main
import "fmt"
func main() {
	var num1 float32
	var num2 float32
	var operator string
	var res float32
	fmt.Print("enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("enter second number: ")
	fmt.Scanln(&num2)
	fmt.Print("what do you want to do? ( * / + - ) ")
	fmt.Scanln(&operator)
	switch operator {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Print("ERROR: zero not valid")
			fmt.Scanln(&num2)
		}
		res = num1 / num2
	default:
		fmt.Print("ERROR: enter valid operator ( / * + -)")
		fmt.Scanln(&operator)
		return
	}
	fmt.Println("\n", num1, operator, num2, "=", res)
}

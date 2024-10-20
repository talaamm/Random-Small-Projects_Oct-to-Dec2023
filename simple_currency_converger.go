package main
import (
	"fmt"
)
//math.Round
func main() {
	dollar := 4.15
	dinar := 5.61
	euro := 4.23
	var klb, klb3 string
	var money float64
	fmt.Print("you want to convert from which currency to shekel(for the opposite enter shekel): ")
	fmt.Scanln(&klb)
	fmt.Print("enter your the amount of money: ")
	fmt.Scanln(&money)
	if klb == "dollar" {
		res := (money * dollar)
		fmt.Println(money, klb, "is", res, "shekel\n")
	} else if klb == "euro" {
		res := money * euro
		fmt.Println(money, klb, "is", res, "shekel\n")
	} else if klb == "dinar" {
		res := money * dinar
		fmt.Println(money, klb, "is", res, "shekel\n")
	} else if klb == "shekel" {
		fmt.Print("what is the currency? ")
		fmt.Scanln(&klb3)
		if klb3 == "dollar" {
			res := money / dollar
			fmt.Println(money, klb, "is", res, klb3, "\n")
		} else if klb3 == "euro" {
			res := money / euro
			fmt.Println(money, klb, "is", res, klb3, "\n")
		} else if klb3 == "dinar" {
			res := money / dinar
			fmt.Println(money, klb, "is", res, klb3, "\n")
		}
	}
}

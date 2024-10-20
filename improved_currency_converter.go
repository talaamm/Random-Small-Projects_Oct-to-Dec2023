package main
import "fmt"
var currmap = map[string]float64{"USD": 3.98, "EURO": 4.24, "JOD": 5.62, "SHEKEL": 1}
func CurrtoShekel(curr string, money float64) float64 {
	for i, r := range currmap {
		if curr == i {
			return money * r
		}
	}
	return -1
}
func ShekeltoCurr(curr string, money float64) float64 {
	for i, r := range currmap {
		if curr == i {
			return money / r
		}
	}
	return -1
}
func main() {
	var myCurr, convert string
	var yourMoney float64
	fmt.Print("\nEnter your amount of money: ")
	fmt.Scanln(&yourMoney)
	fmt.Println("\nEnter your money currency")
	fmt.Print("(USD / SHEKEL / EURO / JOD): ")
	fmt.Scanln(&myCurr)
	fmt.Println("You want to convert it to which currnecy? ")
	fmt.Print("(USD / SHEKEL / EURO / JOD):")
	fmt.Scanln(&convert)
	if myCurr == convert {
		fmt.Println(yourMoney, myCurr, "is", yourMoney, convert)
	
	} else if convert == "SHEKEL" {
		moneyafter := CurrtoShekel(myCurr, yourMoney)
		if moneyafter == -1 {
			fmt.Println("OOPS an errror occured try again!")
			return
		}
		fmt.Println(yourMoney, myCurr, "is", moneyafter, convert)
	
	} else if myCurr == "SHEKEL" {
		moneyafter := ShekeltoCurr(convert, yourMoney)
		if moneyafter == -1 {
			fmt.Println("OOPS an ERROR occured try again!")
			return
		}
		fmt.Println(yourMoney, myCurr, "is", moneyafter, convert)
	
	
	} else {
		pt1 := CurrtoShekel(myCurr, yourMoney)
		pt2 := ShekeltoCurr(convert, pt1)
		if pt2 == -1 || pt1 == -1 {
			fmt.Println("OOPS an errror occured try again!")
			return
		}
		fmt.Println(yourMoney, myCurr, "is", pt2, convert)
	}
}

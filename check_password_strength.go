package main
import (
	"fmt"
	"regexp"
	"unicode"
)
func main() {
	Lengthword := 8
	upper := false
	lower := false
	number := false
	space := false
	symbolSTR := "[!@#$%^&*()]"
	// Get the password input
	var userPassword string
	fmt.Println("Enter your password")
	fmt.Scan(&userPassword)
	// Check password length
	if len(userPassword) < Lengthword { // convert from string to rune to count it right
		fmt.Println("The password is not strong. make it longer", Lengthword, "chartres.")
		return // stop the opretion
	}
	// see if there upper case
	for _, i := range userPassword {
		if unicode.IsUpper(i) {
			upper = true
		}
		if unicode.IsLower(i) {
			lower = true
		}
		if unicode.IsDigit(i) {
			number = true
		}
		complier, _ := regexp.MatchString(symbolSTR, string(i)) // if there are symbole(#$) and space give true
		if complier {
			space = true
		}
	}
	if !upper {
		fmt.Println("Missing uppercase letter!")
	}
	if !lower {
		fmt.Println("Missing lower letter!")
	}
	if !number {
		fmt.Println("Missing number!")
	}
	if !space {
		fmt.Println("Missing special letter!")
	}
	if upper && lower && number && space {
		fmt.Println("The password is strong.")
	} else {
		fmt.Println("\nThe password is not strong.")
	}
}

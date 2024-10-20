
package main
import (
	"fmt"
	"unicode"
)
func main() {
	var password string
	fmt.Print("Enter your Password: ")
	fmt.Scanln(&password)
	var uppercase, lowercase, length, special, number bool
	if len(password) >= 8 {
		length = true
	}
	for _, r := range password {
		if unicode.IsUpper(r) {
			uppercase = true
		}
		if unicode.IsLower(r) {
			lowercase = true
		}
		if unicode.IsDigit(r) {
			number = true
		}
		if unicode.IsPunct(r) {
			special = true
		}
	}
	fmt.Println("The password contains Uppercase:           ", uppercase)
	fmt.Println("The password contains Lowercase:           ", lowercase)
	fmt.Println("The password contains special charachter:  ", special)
	fmt.Println("The password contains number:              ", number)
	fmt.Println("The password contains is long:             ", length)
	if uppercase && lowercase && length && special && number {
		fmt.Println("\n\tPassword is STRONG!")
	} else {
		fmt.Println("\n\tPassword is WEAK!")
	}
}

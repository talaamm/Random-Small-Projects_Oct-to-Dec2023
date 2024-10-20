package main
import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)
const charset = "!@#$%^&*()_-.<>?/.,||\\\"'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
func randomString(length int) string {
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomString)
}
func guess(password string) string {
	charArr := []rune(charset)
	passArr := []rune(password)
	passHacked := ""
	for p := 0; p < len(passArr); p++ {
		for r := 0; r < len(charArr); r++ {
			fmt.Println("Hacking Your Password Please wait...\n\nYour Password is: ", passHacked, charArr[r])
			clearScreen()
			if passArr[p] == charArr[r] {
				passHacked += string(charArr[r])
				break
			}
		}
	}
	return passHacked
}
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func main() {
	var yourPass string
	fmt.Print("enter your password: ")
	fmt.Scanln(&yourPass)
	fmt.Println("\nYour Password is:   ", guess(yourPass))
}

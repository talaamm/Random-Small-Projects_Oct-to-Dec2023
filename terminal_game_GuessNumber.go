package main
import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)
var counterWin int
func randomGame() {
	ClearTerminal()
	var num, max, Lives int = -1, 10, 5 //num -1 since its never gonna be correct (to avoid errors when x = 0 )
	var yn, forerror string
	fmt.Println("\nWELCOME TO GUESS THE NUMBER WITH THE MONAMIES!")
	x := rand.Intn(max + 1)
	fmt.Printf("\nThe computer generated a random number in range of (0-%v)\nGuess the number!\t  Lives:", max)
	for i := Lives; i > 0; i-- {
		fmt.Print("❤️ ")
	}
	fmt.Println()
	//fmt.Println("for test random integer is: ", x)
	for num != x {
		_, err := fmt.Scan(&num)
		fmt.Scanln(&forerror)
		if err != nil {
			fmt.Println("ERROR: enter only integers and try again!")
			break
		}
		if num > x {
			ClearTerminal()
			fmt.Printf("\nWELCOME TO GUESS THE NUMBER WITH THE MONAMIES!\n\nThe computer generated a rnadom number in range of (0-%v)\n", max)
			fmt.Print("try smaller number\t  Lives:")
			Lives--
			if Lives == 0 {
				fmt.Print("💔")
			}
			for i := Lives; i > 0; i-- {
				fmt.Print("❤️ ")
			}
			fmt.Println()
		}
		if num < x {
			ClearTerminal()
			fmt.Printf("\nWELCOME TO GUESS THE NUMBER WITH THE MONAMIES!\n\nThe computer generated a rnadom number in range of (0-%v)\n", max)
			fmt.Print("try bigger number\t  Lives:")
			Lives--
			if Lives == 0 {
				fmt.Print("💔")
			}
			for i := Lives; i > 0; i-- {
				fmt.Print("❤️ ")
			}
			fmt.Println()
		}
		if num == x && Lives > 0 {
			fmt.Println("***WINNER WINNER CHICKEN DINNER!***")
			counterWin++
		}
		if Lives == 0 {
			fmt.Print("OOPS you are out of lives! 😟 \n")
			break
		}
	}
	fmt.Println("\nwanna play again? (y/n) ")
	fmt.Scan(&yn)
	for yn != "" && yn != "y" && yn != "n" {
		fmt.Println("ERROR: invalid input: answer with y for yes or n for no")
		fmt.Scan(&yn)
	}
	if yn == "y" {
		randomGame()
	} else if yn == "n" {
		fmt.Printf("\n***GOODBYE! game ended successfully!***\n***You GEUSSED %v number/s CORRECTLY!***\n\n   ***the monamies were here!***", counterWin)
	}
}
func ClearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func main() {
	randomGame()
}

	package main
import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)
var counter int
var failed, countLives1 int = 0, 5
func randMain5() {
	//in THIS randMain the num of lives is the difference between failed tries and successful ones
	clearTerminal()
	var num, max int = -1, 0 //num -1 since its never gonna be correct to avoid 0 status
	var yn string
	fmt.Println("\nWELCOME TO GUESS THE NUMBER WITH THE MONAMIES!\nThis time your lives count will start form 5\nThen it increases or decreses depending\non your failed tires!\nGOOD LUCK\n")
	fmt.Println("Enter a maximum number for the range of the random number!")
	_, err := fmt.Scan(&max)
	if err != nil {
		fmt.Println("ERROR: enter only integers and try again!")
		if counter != 0 {
			fmt.Printf("\n***You GEUSSED %v number/s CORRECTLY!***\n\n", counter)
		}
		return
	}
	x := rand.Intn(max + 1)
	clearTerminal()
	fmt.Printf("\nThe computer generated a rnadom number in range of (0-%v)\nGuess the number!\t  Lives:", max)
	for i := countLives1; i > 0; i-- {
		fmt.Print("❤️ ")
	}
	fmt.Println()
	//fmt.Println("for test and debug random integer is: ", x)
	for num != x {
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("ERROR: enter only integers and try again!")
			if counter != 0 {
				fmt.Printf("\n***You GEUSSED %v number/s CORRECTLY!***\n\n", counter)
			}
			return
		}
		if num > x {
			clearTerminal()
			fmt.Printf("\nThe computer generated a rnadom number in range of (0-%v)\n", max)
			fmt.Print("try smaller number\t  Lives:")
			countLives1--
			failed++
			if countLives1 == 0 {
				fmt.Print("💔")
			}
			for i := countLives1; i > 0; i-- {
				fmt.Print("❤️ ")
			}
			fmt.Println()
		}
		if num < x {
			clearTerminal()
			fmt.Printf("\nThe computer generated a rnadom number in range of (0-%v)\n", max)
			fmt.Print("try bigger number\t  Lives:")
			countLives1--
			failed++
			if countLives1 == 0 {
				fmt.Print("💔")
			}
			for i := countLives1; i > 0; i-- {
				fmt.Print("❤️ ")
			}
			fmt.Println()
		}
		if num == x && countLives1 > 0 {
			fmt.Println("***WINNER WINNER CHICKEN DINNER!***")
			counter++
			countLives1++
		}
		if countLives1 == 0 {
			fmt.Print("OOPS you are out of lives! ☹️ \n")
			return
		}
	}
	fmt.Println("\nwanna play again? (y/n) ")
	fmt.Scan(&yn)
	for yn != "y" && yn != "n" {
		fmt.Println("ERROR: invalid input: answer with y for yes or n for no")
		fmt.Scan(&yn)
	}
	if yn == "y" {
		randMain5()
	} else if yn == "n" {
		fmt.Printf("\n***GOODBYE Game Ended Successfully!***\n***You GEUSSED %v number/s CORRECTLY!***\n\n", counter)
	}
}
func randMain4() {
	//THIS randMain will be with lives and tries next the num of lives is the difference between failed tries and successful ones
	clearTerminal()
	var num, max int = -1, 0 //num -1 since its never gonna be correct to avoid 0 status
	var yn string
	var countLives0 = 3
	fmt.Println("\nWELCOME TO GUESS THE NUMBER WITH THE MONAMIES!")
	fmt.Println("Enter a maximum number for the range of the random number!")
	_, err := fmt.Scan(&max)
	if err != nil {
		fmt.Println("ERROR: enter only integers and try again!")
		if counter != 0 {
			fmt.Printf("\n***You GEUSSED %v number/s CORRECTLY!***\n\n", counter)
		}
		return
	}
	x := rand.Intn(max + 1)
	clearTerminal()
	fmt.Printf("\nThe computer generated a rnadom number in range of (0-%v)\nGuess the number!\t  Lives:", max)
	for i := countLives0; i > 0; i-- {
		fmt.Print("❤️ ")
	}
	fmt.Println()
	//fmt.Println("for test and debug random integer is: ", x)
	for num != x {
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("ERROR: enter only integers and try again!")
			if counter != 0 {
				fmt.Printf("\n***You GEUSSED %v number/s CORRECTLY!***\n\n", counter)
			}
			return
		}
		if num > x {
			clearTerminal()
			fmt.Printf("\nThe computer generated a rnadom number in range of (0-%v)\n", max)
			fmt.Print("try smaller number\t  Lives:")
			countLives0--
			if countLives0 == 0 {
				fmt.Print("💔")
			}
			for i := countLives0; i > 0; i-- {
				fmt.Print("❤️ ")
			}
			fmt.Println()
		}
		if num < x {
			clearTerminal()
			fmt.Printf("\nThe computer generated a rnadom number in range of (0-%v)\n", max)
			fmt.Print("try bigger number\t  Lives:")
			countLives0--
			if countLives0 == 0 {
				fmt.Print("💔")
			}
			for i := countLives0; i > 0; i-- {
				fmt.Print("❤️ ")
			}
			fmt.Println()
		}
		if num == x && countLives0 > 0 {
			fmt.Println("***WINNER WINNER CHICKEN DINNER!***")
			counter++
		}
		if countLives0 == 0 {
			fmt.Print("OOPS you are out of lives! :'( \n")
			break
		}
	}
	fmt.Println("\nwanna play again? (y/n) ")
	fmt.Scan(&yn)
	for yn != "y" && yn != "n" {
		fmt.Println("ERROR: invalid input: answer with y for yes or n for no")
		fmt.Scan(&yn)
	}
	if yn == "y" {
		randMain4()
	} else if yn == "n" {
		fmt.Printf("\n***GOODBYE session ended successfully!***\n***You GEUSSED %v number/s CORRECTLY!***\n\n", counter)
	}
}
func clearTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func randMain3() {
	//gonna add clear the terminal next randMain will be with lives and tries
	clearTerminal()
	var num, max int
	var yn string
	fmt.Println("\nWELCOME TO GUESS THE NUMBER WITH THE MONAMIES!")
	fmt.Println("Enter a maximum number for the range of the random number!")
	_, err := fmt.Scan(&max)
	if err != nil {
		fmt.Println("ERROR: enter only integers and try again!")
		if counter != 0 {
			fmt.Printf("\n***You GEUSSED %v number/s CORRECTLY!***\n\n", counter)
		}
		return
	}
	x := rand.Intn(max + 1)
	clearTerminal()
	fmt.Printf("\nThe computer generated a rnadom number in range of (0-%v)\nGuess the number! ", max)
	//fmt.Println("for test and debug random integer is: ", x)
	for num != x {
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("ERROR: enter only integers and try again!")
			if counter != 0 {
				fmt.Printf("\n***You GEUSSED %v number/s CORRECTLY!***\n\n", counter)
			}
			return
		}
		if num > x {
			fmt.Println("try smaller number")
		} else if num < x {
			fmt.Println("try bigger number")
		} else if num == x {
			fmt.Println("***WINNER WINNER CHICKEN DINNER!***")
			counter++
		}
	}
	fmt.Println("\nwanna play again? (y/n) ")
	fmt.Scan(&yn)
	for yn != "y" && yn != "n" {
		fmt.Println("ERROR: invalid input: answer with y for yes or n for no")
		fmt.Scan(&yn)
	}
	if yn == "y" {
		randMain3()
	} else if yn == "n" {
		fmt.Printf("\n***GOODBYE session ended successfully!***\n***You GEUSSED %v number/s CORRECTLY!***\n\n", counter)
	}
}
func randMain2() {
	var num, max int
	var yn string
	fmt.Println("\n\nWELCOME TO GUESS THE NUMBER WITH THE MONAMIES!")
	fmt.Println("Enter a maximum number for the range of the random number!")
	_, err := fmt.Scan(&max)
	if err != nil {
		fmt.Println("ERROR: enter only integers and try again!")
		return
	}
	x := rand.Intn(max + 1)
	fmt.Printf("\nThe computer generated a number in range of (0-%v)\nGuess the number! ", max)
	fmt.Println("for test and debug random integer is: ", x)
	for num != x {
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("ERROR: enter only integers and try again!")
			return
		}
		if num > x {
			fmt.Println("try smaller number")
		} else if num < x {
			fmt.Println("try bigger number")
		} else if num == x {
			fmt.Println("***WINNER WINNER CHICKEN DINNER!***")
			counter++
		}
	}
	fmt.Println("\nwanna play again? (y/n) ")
	fmt.Scan(&yn)
	for yn != "y" && yn != "n" {
		fmt.Println("invalid input: answer with y for yes or n for no")
		fmt.Scan(&yn)
	}
	if yn == "y" {
		randMain2()
	} else if yn == "n" {
		fmt.Printf("\n***GOODBYE session ended successfully!***\n***You GEUSSED %v numbers CORRECTLY!***\n\n", counter)
	}
}
func randMain1() {
	x := rand.Intn(101)
	var num int
	var yn string
	fmt.Print("\n\nThe computer generated a number in range of (0-100)\nGuess the number! ")
	//fmt.Println("for test and debug random integer is: ", x)
	for num != x {
		fmt.Scan(&num)
		if num > x {
			fmt.Println("try smaller number")
		} else if num < x {
			fmt.Println("try bigger number")
		} else if num == x {
			fmt.Println("***WINNER WINNER CHICKEN DINNER!***")
			counter++
		}
	}
	fmt.Println("\nwanna play again? (y/n) ")
	fmt.Scan(&yn)
	for yn != "y" && yn != "n" {
		fmt.Println("invalid input: answer with y for yes or n for no")
		fmt.Scan(&yn)
	}
	if yn == "y" {
		randMain1()
	} else if yn == "n" {
		fmt.Printf("\n***GOODBYE session ended successfully!***\n***You GEUSSED %v numbers CORRECTLY!***\n\n", counter)
	}
}
func randMain0() {
	x := rand.Intn(100)
	var num int
	fmt.Print("guess the number\n")
	for num != x {
		fmt.Scan(&num)
		if num > x {
			fmt.Println("try smaller number")
		} else if num < x {
			fmt.Println("try bigger number")
		} else if num == x {
			fmt.Println("you are winner")
		}
	}
}
func main() {
	randMain5()
}

package main
import "fmt"
func main() {
	numOfLines := 0
	numOfPyramids := 0
	fmt.Print("enter the number of lines of a single pyramid: ")
	fmt.Scanln(&numOfLines)
	fmt.Print("enter the number of pyramids: ")
	fmt.Scanln(&numOfPyramids)
	//numOfLines := 6
	//numOfPyramids := 4
	for i := 1; i <= numOfLines; i++ {
		for yy := 1; yy <= numOfPyramids; yy++ {
			for y := numOfLines; y >= i-1; y-- {
				fmt.Print(" ")
			}
			for z := 1; z <= 2*i-1; z++ {
				fmt.Print("*")
			}
			if i == 1 && numOfLines == 5 {
				for y := numOfLines; y >= i-1; y-- {
					fmt.Print(" ")
				}
			}
			fmt.Print("\t")
		}
		fmt.Println()
	}
}

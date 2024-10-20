/*
	    1
	   232
	  34543
	 4567654
	567898765
*/
package main
import "fmt"
func main() {
	for lineCount := 1; lineCount <= 5; lineCount++ {
		b := true
		start := 1
		if lineCount == 1 {
			fmt.Println(lineCount)
		} else if lineCount > 1 && lineCount < 4 {
			if lineCount%2 == 1 {
				start = lineCount
			}
			for i := lineCount; (i <= start+2) && (i >= lineCount); {
				if i == start+2 {
					fmt.Print(i)
					b = false
					i--
				} else if b {
					fmt.Print(i)
					i++
				} else if b == false {
					fmt.Print(i)
					i--
				}
			}
			fmt.Println()
		} else if lineCount == 4 {
			for i := lineCount; (i <= lineCount+3) && (i >= lineCount); {
				if i == lineCount+3 {
					fmt.Print(i)
					b = false
					i--
				} else if b {
					fmt.Print(i)
					i++
				} else if b == false {
					fmt.Print(i)
					i--
				}
			}
			fmt.Println()
		} else if lineCount == 5 {
			for i := lineCount; (i <= lineCount+4) && (i >= lineCount); {
				if i == lineCount+4 {
					fmt.Print(i)
					b = false
					i--
				} else if b {
					fmt.Print(i)
					i++
				} else if b == false {
					fmt.Print(i)
					i--
				}
			}
			fmt.Println()
		}
	}
}

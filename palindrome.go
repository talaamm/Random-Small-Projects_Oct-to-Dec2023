package main
import "fmt"
func main() {
	str := "fe5ef"
	isP := true
	arr := []rune(str)
	if len(arr)%2 == 1 {
		indx := len(arr)/2 + 1
		str1 := []rune(arr[:indx])   // ret
		str2 := []rune(arr[indx+1:]) //ter
		for i, b := 0, len(str2)-1; i < len(str1) && b >= 0; i, b = i+1, b-1 {
			if str2[b] != str1[i] {
				isP = false
			}
		}
	} else {
		isP = false
	}
	if isP {
		fmt.Println(str, "is palindrome")
	} else {
		fmt.Println(str, "is NOT palindrome")
	}
}

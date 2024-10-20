package main
import (
	"fmt"
	"strconv"
)
func main() {
	var item_num = map[int]string{
		150: "tomato",
		151: "khiar",
		152: "potato",
		153: "beitinjan",
		154: "onion",
		155: "lemon",
		156: "kosa",
		157: "felfel",
		158: "banana",
		159: "apple",
		160: "orange",
	}
	var as3ar = map[string]float64{
		"tomato":    5,
		"khiar":     5,
		"potato":    3.3,
		"beitinjan": 3.3,
		"onion":     5,
		"lemon":     10,
		"kosa":      10,
		"felfel":    12,
		"banana":    5,
		"apple":     6.6,
		"orange":    3.3,
	}
	fmt.Println("\n   !!Welcome to dokan the monamies!!\n")
	for in, esem := range item_num {
		for i, oo := range as3ar {
			if esem == i {
				fmt.Println("#"+strconv.Itoa(in), " (1 kilo)", i, fmt.Sprintf("   %0.1f", oo)+"₪")
			}
		}
	}
	var inputNum int
	var quan float64
	var price []float64
	var fatora []string
	var yn string
	for yn != "n" {
		fmt.Print("\nWhat Would you like to buy? (Enter item number): ")
		fmt.Scanln(&inputNum)
		for r := range item_num {
			if inputNum == r {
				fmt.Println("#"+strconv.Itoa(r), " (1 kilo)", item_num[r], fmt.Sprintf("%0.1f", as3ar[item_num[r]])+"₪")
				fmt.Print("\nwhat is the quantity(bel kilo)? ")
				fmt.Scanln(&quan)
				fatora = append(fatora, fmt.Sprintf("%0.1f", quan)+" kilo "+item_num[r])
				price = append(price, quan*as3ar[item_num[r]])
			}
		}
		fmt.Println("Would you like to buy more (y/n)? ")
		fmt.Scanln(&yn)
	}
	fmt.Println("\nYou Ordered:", arrStoS(fatora))
	fmt.Println("\nYour total price is: ", sumArr(price), "₪")
}
func sumArr(array []float64) float64 {
	res := 0.0
	for _, num := range array {
		res += num
	}
	return res
}
func arrStoS(array []string) string {
	str := ""
	for _, s := range array {
		str += "\n" + s
	}
	return str
}

package main

import "fmt"

func main() {
	cardNumber := []int{4, 0, 2, 4, 0, 0, 0, 9, 0, 2, 2, 1, 4, 3}
	result := checkValidNum(cardNumber)
	answer := ""
	if result {
		answer = "card is valid "
	} else {
		answer = "card is invalid"
	}
	fmt.Println(answer)
}
func checkValidNum(cardNumber []int) (result bool) {
	total := 0
	for i := 0; i < len(cardNumber); i++ {
		if i%2 == 0 {
			cardNumber[i] = cardNumber[i] * 2
			if cardNumber[i] > 9 {
				cardNumber[i] = cardNumber[i]/10 + cardNumber[i]%10
			}
		}
		total += cardNumber[i]

	}
	if total%10 == 0 {
		return true
	}
	return false
}

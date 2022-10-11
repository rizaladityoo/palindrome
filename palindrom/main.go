package main

import "fmt"

func main() {
	kata := "madam"

	fmt.Println(isPalindrom(kata))
}

func isPalindrom(kata string) bool {
	var loop int
	if loop = len(kata) / 2; len(kata)%2 != 0 {
		loop = (len(kata) - 1) / 2
	}

	for i := 0; i < loop; i++ {
		if kata[i] != kata[len(kata)-i-1] {
			return false
		}
	}
	return true
}

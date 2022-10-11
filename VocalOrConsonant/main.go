package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var name string
	fmt.Print("masukan huruf : ")
	fmt.Scanf("%s", &name)

	name = strings.ToLower(name)

	if len(name) > 1 {
		fmt.Println("masukan huruf")
		os.Exit(3)
	}

	if name == "a" || name == "i" || name == "u" || name == "e" || name == "o" {
		fmt.Println("vocal")
	} else {
		fmt.Println("konsonan")
	}
}

// main.go
package main

import (
	"fmt"
)

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is Even")
	} else {
		fmt.Println("7 is Odd")
	}

	if num := 2; num < 10 {
		fmt.Println("Num is less then 10")
	}
}

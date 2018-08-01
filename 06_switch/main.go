// main.go
package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Before noon")
	default:
		fmt.Println("After Noon")
	}

	switch time.Now().Weekday() {
	case time.Wednesday, time.Thursday:
		fmt.Println("Wed / Thu")
	default:
		fmt.Println("Other weekday")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("boolean")
		case int:
			fmt.Println("integer")
		case string:
			fmt.Println("string")
		default:
			fmt.Printf("Unknown - %T \n", t)
		}
	}

	whatAmI(12)
	whatAmI(true)
	whatAmI("hello")
	whatAmI(3.4)
}

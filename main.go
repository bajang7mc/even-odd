package main

import "fmt"

func main() {

	times := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, time := range times {
		if time%2 == 0 {
			fmt.Println(time, "is Even")
		} else {
			fmt.Println(time, "is Odd")
		}
	}

}

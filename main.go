package main

import "fmt"

func main() {

	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter text: ")
	//text, _ := reader.ReadString('\n')
	//text = strings.Replace(text, "\n", "", -1)
	//times, _ := strconv.Atoi(text)

	//fmt.Println(times)
	//fmt.Println(text)

	//for i := 0; i <= times[]; i++ {
	//	if i%2 == 0 {
	//		fmt.Println(i, "is even")
	//	} else {
	//		fmt.Println(i, "is odd")
	//	}
	//}

	times := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, time := range times {
		if time%2 == 0 {
			fmt.Println(time, "is even")
		} else {
			fmt.Println(time, "is odd")
		}
	}

}

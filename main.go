package main

import "fmt"
import "os"
import "strconv"

func main() {
	if len(os.Args)<2 {
		fmt.Println("ERROR!")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 0 || n > 10 {
		fmt.Println("ERROR!")
		return
	}
	for i := 0; i < n; i++ {
		fmt.Println("adventurous-gerbil")
	}
}

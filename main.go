package main

import "bufio"
import "fmt"
import "os"

func main() {
	fmt.Printf("Enter 1 to use mergesort recursively or 2 to sort iteratively\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		choice := scanner.Text()
		fmt.Printf("you entered " + choice + "\n")
		if choice == "1" || choice == "2" {
			fmt.Printf("you made the right decision, lets quit\n")
			break
		}
	}
}

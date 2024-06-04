package main

import "fmt"

func main() {
	fizzBuzz := []string{
		"FizzBuzz",
		"Fizz",
		"Buzz",
	}

	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println(fizzBuzz[0])
		} else if i % 3 == 0 {
			fmt.Println(fizzBuzz[1])
		} else if i % 5 == 0 {
			fmt.Println(fizzBuzz[2])
		} else {
			fmt.Println(i)
		}
	}
}
package main

import "fmt"

func main() {
	//for every number  from 1 to 20
	for i := 1; i < 20; i++ {
		//if the number is divisible by 3 and 5 print fizz buzz
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz")
			//else if the number is divisible by 3 print fizz
		} else if i%3 == 0 {
			fmt.Println("Fizz")
			//else if the number is divisible by 5 print buzz
		} else if i%5 == 0 {
			fmt.Println("buzz")
			//else print the number
		} else {
			fmt.Println(i)
		}

	}

}

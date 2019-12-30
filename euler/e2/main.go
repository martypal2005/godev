package main

import "fmt"

func main() {

	frs := 1
	sec := 2

	sum := sec

	for cur := frs + sec; cur < 4000000; cur = frs + sec {

		frs = sec
		sec = cur

		if cur%2 == 0 {
			sum += cur
			fmt.Println(cur)
		}
	}

	fmt.Println("--")
	fmt.Println(sum)
}

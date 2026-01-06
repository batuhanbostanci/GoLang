package main

import "fmt"

func main() {

	fact := factorial(5)
	fmt.Println(fact)

	numbers := []int{1, 2, 3, 4, 5}

	sum := sumup(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(sum)

	anotherSum := sumup2(1, numbers...)
	fmt.Println(anotherSum)

}

func sumup2(n int, numbers ...int) int {
	sum := n

	for _, val := range numbers {
		sum += val
	}

	return sum
}

// Variadic function, can accept any number of arguments
func sumup(number ...int) int {
	sum := 0

	for _, val := range number {
		sum += val
	}

	return sum
}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

	// result := 1

	// for i := 1; i <= n; i++ {

	// 	result = result * i
	// }
	// return result
}

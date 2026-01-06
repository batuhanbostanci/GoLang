package main

import "fmt"

type transformFn func(int) int

// type anotherFn func(int, []string, map[string]int) (string, int)

func main() {
	numbers := []int{1, 2, 3, 4}

	//doubled := transformNumbers(&numbers, double)
	//tripled := transformNumbers(&numbers, triple)

	moreNumbers := []int{5, 1, 2}

	//fmt.Println(doubled)
	//fmt.Println(tripled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)

	transformedNumbbers := transformNumbers(&numbers, transformerFn1)
	transformedMoreNumbbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformedNumbbers)
	fmt.Println(transformedMoreNumbbers)

}

func transformNumbers(numbers *[]int, transform transformFn) []int {

	dNumbers := []int{}
	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))
	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {

	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}
func double(number int) int {
	return number * 2
}

func triple(number int) int {

	return number * 3
}

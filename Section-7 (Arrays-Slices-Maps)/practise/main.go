package main

import "fmt"

func main() {
	// 1) Create a new array with three hobbies
	hobbies := [3]string{"Cooking", "Photography", "Coding"}
	fmt.Println("Hobbies Array:", hobbies)

	// 2) Output more data
	fmt.Println("First Hobby:", hobbies[0])

	// Second and third element as a new list (slice)
	subHobbies := hobbies[1:3]
	fmt.Println("Second & Third:", subHobbies)

	// 3) Create a slice based on the original array
	// (Note: The prompt says "based on the first element", but usually we slice the array itself)
	// Task: contains 1st and 2nd element
	slice1 := hobbies[0:2]
	slice2 := hobbies[:2] // Shorthand way
	fmt.Println("Slice way 1:", slice1)
	fmt.Println("Slice way 2:", slice2)

	// 4) Re-slice the slice from (3)
	// Change it to contain 2nd and last (3rd) element of original array
	// Original hobbies: [0: Cooking, 1: Photography, 2: Coding]
	reSliced := slice1[1:3]
	fmt.Println("Re-sliced (2nd and 3rd):", reSliced)

	// 5) Create a "dynamic array" (Slice) for course goals
	goals := []string{"Learn Go Basics", "Build an API"}
	fmt.Println("Goals:", goals)

	// 6) Change 2nd goal and add a 3rd
	goals[1] = "Master Concurrency"
	goals = append(goals, "Deploy an App")
	fmt.Println("Updated Goals:", goals)

	// 7) Bonus: Product struct and dynamic list
	type Product struct {
		id    string
		title string
		price float64
	}

	products := []Product{
		{id: "p1", title: "Mechanical Keyboard", price: 99.99},
		{id: "p2", title: "Gaming Mouse", price: 49.99},
	}

	newProduct := Product{id: "p3", title: "UltraWide Monitor", price: 399.99}
	products = append(products, newProduct)

	fmt.Println("Product List:", products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

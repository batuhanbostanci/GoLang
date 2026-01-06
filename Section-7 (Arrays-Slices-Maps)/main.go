package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {

	//username := make([]string, 3,5)
	usernames := make([]string, 2)

	usernames[0] = "john"
	usernames[1] = "doe"

	usernames = append(usernames, "alice")
	usernames = append(usernames, "bob")

	fmt.Println(usernames)

	//courseRatings := make(map[string]float64, 3)

	courseRatings := make(floatMap)

	courseRatings["GoLang"] = 4.5
	courseRatings["Python"] = 4.7
	courseRatings["JavaScript"] = 4.3

	//fmt.Println(courseRatings)
	courseRatings.output()

	for i, j := range usernames {
		fmt.Println(i, j)
	}

	for i, j := range courseRatings {
		fmt.Println(i, j)
	}
}

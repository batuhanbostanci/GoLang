package main

import "fmt"

func main() {

	websites := map[string]string{
		"Google":        "https://www.google.com",
		"GitHub":        "https://www.github.com",
		"StackOverflow": "https://stackoverflow.com",
	}

	fmt.Println(websites)

	websites["Reddit"] = "https://www.reddit.com"

	fmt.Println(websites)

	fmt.Println(websites["Google"])

	delete(websites, "Reddit")
	fmt.Println(websites)
}

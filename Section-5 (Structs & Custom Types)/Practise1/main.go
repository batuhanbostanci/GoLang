// package main

// import "fmt"

// // Define a custom type 'str' based on the built-in 'string' type

// type str string

// func (text str) log() {

// 	fmt.Println(text)
// }

// func main() {
// 	var name str = "Batu"

// 	name.log()
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}
	fmt.Println("Note saved successfully.")
}

func getNoteData() (string, string) {

	title := getUserInput("Note Title:")

	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(promt string) string {
	// var input string
	// fmt.Print(promt)
	// Scanln only accepts 1 word inputs
	//fmt.Scanln(&input)

	fmt.Printf("%v ", promt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // For Windows compatibility

	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	return text
}

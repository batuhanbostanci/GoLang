package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/interfacesAndGeneric/note"
	"example.com/interfacesAndGeneric/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {

	printSomething(1)
	printSomething("Hello, Go Interfaces!")
	printSomething(true)

	title, content := getNoteData()
	todoText := getUserInput("To text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	err = outputData(todo)

	if err != nil {

	}

	err = outputData(userNote)

	if err != nil {
		return
	}
}

func printSomething(value interface{}) {

	typedVal, ok := value.(int)

	if ok {
		fmt.Println("Integer value:", typedVal)
		return
	}

	typedStr, ok := value.(string)
	if ok {
		fmt.Println("String value:", typedStr)
		return
	}

	//This is the one way of using it
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer value:", value)
	// case string:
	// 	fmt.Println("String value:", value)
	// case bool:
	// 	fmt.Println("Boolean value:", value)

	// }

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(s saver) error {

	err := s.Save()

	if err != nil {
		fmt.Println("Error saving data:", err)
		return err
	}
	fmt.Println("Data saved successfully.")
	return nil
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

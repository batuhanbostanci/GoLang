package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (*Todo, error) {

	if text == "" {
		return nil, errors.New("todo text is a required field")
	}

	return &Todo{
		Text: text,
	}, nil
}

func (n Todo) Save() error {

	fileName := "todo.json"

	json, error := json.Marshal(n)

	if error != nil {
		fmt.Println("Error while marshalling note to json:", error)
		return error
	}
	fmt.Print(fileName)
	return os.WriteFile(fileName, json, 0644)

}

func (n *Todo) Display() {

	fmt.Printf("Yor todo text is %v", n.Text)

}

package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n Note) Save() error {

	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, error := json.Marshal(n)

	if error != nil {
		fmt.Println("Error while marshalling note to json:", error)
		return error
	}
	fmt.Print(fileName)
	return os.WriteFile(fileName, json, 0644)

}

func New(title string, content string) (*Note, error) {

	if title == "" || content == "" {
		return nil, errors.New("note title and content are required fields")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n *Note) Display() {

	fmt.Printf("Yor note titled %vhas the following content:\n\n%v\n", n.Title, n.Content)

}

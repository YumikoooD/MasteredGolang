package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note/note"
	"example.com/note/todo/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		fmt.Println(err)
		return
	}
	printSomething(42)
	printSomething("I juste print something")
	printSomething(42.97)
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving failed.")
		return err
	}

	fmt.Println("Saving succeeded!")
	return nil
}

func printSomething(value any) {
	fmt.Println(value)
	switch value.(type) {
	case int:
		fmt.Println("This is an int")
	case string:
		fmt.Println("this is a string")
	case float64:
		fmt.Println("this is a float")
	}
}

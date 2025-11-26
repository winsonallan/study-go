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
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the node succeeded!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) (string) {
	fmt.Printf("%v ", prompt)
	
	// Creates a reader
	reader := bufio.NewReader(os.Stdin) //Stdin is the command line

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	// Removes line breaks that could be problematic in the user input
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
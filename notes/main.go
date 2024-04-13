package main

import (
	"bufio"
	"fmt"
	"notes/note"
	"notes/todo"
	"os"
	"strings"
)

type saver interface {
	Save() error 
}

// type displayer interface {
// 	Display()
// }

// type outputable interface {
// 	Save() error
// 	Display()
// }

type outputable interface {
	saver
	Display()
}

//You can also create an embedded interface
// You can specify what kind of values the interface takes in for a function int, string
// If there is only one method inside a interface the interface ame will be name of the method+er more tha one -able

func main() {
	printSomething(1)
	printSomething(1.3)
	printSomething("hi")


	title, content := getNoteData()
	todoText := getTodoData()

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
	outputData(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func printSomething(value interface{}) {
	intValue, ok := value.(int)

	if ok {
		fmt.Println("Integer: ", intValue)
		return
	}

	floatValue, ok := value.(float64)

	if ok {
		fmt.Println("Integer: ", floatValue)
		return
	}

	stringValue, ok := value.(string)

	if ok {
		fmt.Println("Integer: ", stringValue)
		return
	}

	// switch value.(type) {
	// 	case int:
	// 		fmt.Println("Integer: ", value)
	// 	case float64:
	// 		fmt.Println("Float: ", value)
	// 	case string:
	// 		fmt.Println("String: ", value)
	// 	// default:
	// 	// 	fmt.Println("Not supported: ", value)
	// }
	fmt.Println(value)
} // Used to input any kind of value

func outputData(data outputable) (error) {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("saving data failed")
		return err
	}
	fmt.Println("saving data succeeded")
	return nil
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n') // Its called Rune

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getTodoData() string {
	text := getUserInput("Todo text: ")
	return text
}

// For a lenghty string you can't use Scan we have to go with a different method

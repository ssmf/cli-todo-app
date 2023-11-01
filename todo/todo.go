package todo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gookit/color"
)

var Purple = color.Magenta.Render
var Cyan = color.Cyan.Render
var Red = color.Red.Render
var Green = color.Green.Render
var Reset = color.White.Render

type Task struct {
	Name string
	Done bool
}

type TaskList []Task

func (l *TaskList) WriteJson() {
	file, _ := json.MarshalIndent(*l, "", " ")
	os.WriteFile("data.json", file, 0644)
}

func (l *TaskList) AddTask(t string, d bool) {

	*l = append(*l, Task{t, d})

}

func (l *TaskList) RemoveTask(name string) {
	var indx int = -1
	for i, v := range *l {
		if v.Name == name {
			indx = i
		}
	}

	if indx != -1 {
		*l = append((*l)[:indx], (*l)[indx+1:]...)
	}
}

func (t *Task) DoTask() {
	t.Done = !t.Done
}

func (l *TaskList) DisplayList(b bool) {
	if b {
		for i := 0; i < 75; i++ {
			fmt.Printf("-")
		}
		fmt.Printf("\n")
		for i, v := range *l {
			fmt.Printf("[%d]: %s - %s\n", i+1, Purple(v.Name), Purple(IfDoneDisplay(v.Done)))
		}
		for i := 0; i < 75; i++ {
			fmt.Printf("-")
		}
		fmt.Printf("\n")
	}
}

func DisplayHelp() {
	fmt.Printf("\nHere is a high level explenation of how this app works: ")
	fmt.Printf("\n\nEvery part of this cli app is related to flags. For an example, if we want to:"+
		" \n\n1.Add a new task"+
		" \n2.Provide a name for it"+
		"\n\nWe'll use 2 flags: -add as well as -name here are the flags applied:"+
		"\n\n%s", Cyan("main.exe -add -name Read"))

	fmt.Printf("\n\nHere is a list of all commands available:")
	//All flags explained
	fmt.Printf("\n\n")
	for i := 0; i < 75; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n%s -> %s %s %s", Green("-add"), Cyan("Specifies if user wants to"), Green("add"), Cyan("a new task"))
	fmt.Printf("\n%s -> %s %s %s", Green("-rm"), Cyan("Specifies if user wants to"), Red("remove"), Cyan("an existing task"))
	fmt.Printf("\n%s %s -> %s", Green("-name"), Purple("[value]"), Cyan("Defines task name"))
	fmt.Printf("\n%s -> %s %s", Green("-done"), Cyan("Defines if task is done, if not used, value is"), Red("false"))
	fmt.Printf("\n%s -> %s", Green("-list"), Cyan("Displays current list of todos"))
	fmt.Printf("\n%s -> %s", Green("-help"), Cyan("Displays all commands and their roles"))
	fmt.Printf("\n")
	for i := 0; i < 75; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
}

func IfDoneDisplay(b bool) string {
	var txt string
	if b {
		txt = Green("Done!")
	} else {
		txt = Red("Not done yet!")
	}
	return txt
}

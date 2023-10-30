package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

var Purple = "\033[35m"
var Cyan = "\033[36m"
var Red = "\033[31m"
var Green = "\033[32m"
var Reset = "\033[0m"

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
		for i, v := range *l {
			fmt.Printf("\n[%d]: %s%s - %s \n\n%s", i+1, Purple, v.Name, IfDoneDisplay(v.Done), Reset)
		}
	}
}

func DisplayHelp() {
	fmt.Printf("Here is a high level explenation of how this app works: ")
	fmt.Printf("\n\nEvery part of this cli app is related to flags. For an example, if we want to:"+
		" \n1.Add a new task"+
		" \n2.Provide a name for it"+
		"\nWe'll use 2 flags: -add as well as -name here are the flags applied:"+
		"\n\n%smain.exe -add -name Read%s", Cyan, Reset)

	fmt.Printf("\n\nHere is a list of all commands available:")
	//All flags explained
	fmt.Printf("\n\n")
	for i := 0; i < 30; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n%s-add %s-> %sSpecifies if user wants to %sadd a new task%s", Green, Reset, Cyan, Green, Reset)
	fmt.Printf("\n%s-rm %s-> %sSpecifies if user wants to %sremove an existing task%s", Green, Reset, Cyan, Red, Reset)
	fmt.Printf("\n%s-name %s[value] %s-> %sDefines a task name%s", Green, Purple, Reset, Cyan, Reset)
	fmt.Printf("\n%s-done %s-> %sDefines if task is done, if not used value is %sfalse%s", Green, Reset, Cyan, Red, Reset)
	fmt.Printf("\n%s-list %s-> %sDisplays current list of todos%s", Green, Reset, Cyan, Reset)
	fmt.Printf("\n%s-help %s-> %sDisplays all commands and their roles%s", Green, Reset, Cyan, Reset)
	fmt.Printf("\n")
	for i := 0; i < 30; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
}

func IfDoneDisplay(b bool) string {
	var txt string
	if b {
		txt = Green + "Done!"
	} else {
		txt = Red + "Not done yet!"
	}
	return txt
}

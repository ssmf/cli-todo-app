package main

import (
	"encoding/json"
	"flag"
	"os"

	"example.com/packages/todo"
)

func main() {

	var MainTaskList todo.TaskList
	DataFile, _ := os.ReadFile("data.json")

	json.Unmarshal(DataFile, &MainTaskList)

	AddTask := flag.Bool("add", false, "Specifies if user wants to add a new task")
	RemoveTask := flag.Bool("remove", false, "Specifies if user wants to remove an existing task")
	TaskName := flag.String("name", "BoilerplateTask", "Define a task name")
	TaskDone := flag.Bool("done", false, "Define if task is done")
	ShowList := flag.Bool("list", false, "Show current list of todos")

	flag.Parse()

	if *AddTask && *RemoveTask {
		panic("You cannot remove and add tasks at the same time")
	} else if *AddTask {
		MainTaskList.AddTask(*TaskName, *TaskDone)
	} else if *RemoveTask {
		// MainTaskList
	}
	MainTaskList.DisplayList(*ShowList)

}

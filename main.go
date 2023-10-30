package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"example.com/packages/todo"
)

func main() {

	var MainTaskList todo.TaskList
	DataFile, _ := os.ReadFile("data.json")

	json.Unmarshal(DataFile, &MainTaskList)

	AddTask := flag.Bool("add", false, "Specifies if user wants to add a new task")
	RemoveTask := flag.Bool("rm", false, "Specifies if user wants to remove an existing task")
	TaskName := flag.String("name", "BoilerplateTask", "Defines a task name")
	TaskDone := flag.Bool("done", false, "Defines if task is done")
	ShowList := flag.Bool("list", false, "Show current list of todos")
	Help := flag.Bool("help", false, "Displays all commands and their roles")

	flag.Parse()

	if *Help && !*AddTask && !*RemoveTask && !*ShowList && !*TaskDone {
		fmt.Println("it works!")
		todo.DisplayHelp()
	}

	if *AddTask && *RemoveTask {
		panic("You cannot remove and add tasks at the same time")
	} else if *AddTask {
		MainTaskList.AddTask(*TaskName, *TaskDone)
	} else if *RemoveTask {
		MainTaskList.RemoveTask(*TaskName)
	}
	MainTaskList.DisplayList(*ShowList)

	if *TaskDone && !*RemoveTask && !*AddTask {
		for i, v := range MainTaskList {
			if v.Name == *TaskName {
				MainTaskList[i].DoTask()
			}
		}
	}
	MainTaskList.WriteJson()
}

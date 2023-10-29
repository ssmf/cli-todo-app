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

	TaskName := flag.String("taskname", "BoilerplateTask", "Define a task name")
	TaskDone := flag.Bool("done", false, "Define if task is done")
	ShowList := flag.Bool("list", false, "Show current list of todos")

	flag.Parse()

	fmt.Println(*TaskName)

	MainTaskList.AddTask(*TaskName, *TaskDone)
	MainTaskList.DisplayList(*ShowList)

}

package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

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
	l.WriteJson()

}

func (l *TaskList) RemoveTask(i int) {

}

func (t *Task) DoTask() {
	t.Done = !t.Done
}

func (l *TaskList) DisplayList(b bool) {
	if b {
		for i, v := range *l {
			fmt.Printf("[%d]: %s - %s \n", i+1, v.Name, IfDoneDisplay(v.Done))
		}
	}
}

func IfDoneDisplay(b bool) string {
	var txt string
	if b {
		txt = "Done!"
	} else {
		txt = "Not done yet!"
	}
	return txt
}

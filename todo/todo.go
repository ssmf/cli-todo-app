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

func (l *TaskList) AddTask(t string, d bool) {

	*l = append(*l, Task{t, d})
	file, _ := json.MarshalIndent(*l, "", " ")
	os.WriteFile("data.json", file, 0644)

}

func (t Task) DoTask() {
	t.Done = !t.Done
}

func (l *TaskList) DisplayList(b bool) {
	if b == true {
		for _, v := range *l {
			fmt.Println(v)
		}
	}
}

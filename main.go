package main

import (
	"errors"
	"fmt"
)

type TaskStatus string

const (
	StatusNotStarted TaskStatus = "not started"
	StatusRunning    TaskStatus = "running"
	StatusCompleted  TaskStatus = "completed"
	StatusFailed     TaskStatus = "failed"
)

type Task struct {
	Name   string
	Status TaskStatus
	Action func() error
}

func fetch() error {
	fmt.Println("fetching data...")
	return nil
}

func testError() error {
	fmt.Println("something happened")
	return nil
}

func parse() error {
	fmt.Println("parsing data...")
	return errors.New("invalid data")
}

func download() error {
	fmt.Println("downloading data...")
	return nil
}

type Workflow struct {
	Name  string
	Tasks []Task
}

func (w Workflow) Print() {
	fmt.Println(w.Name)
	fmt.Println(w.Tasks)
}

func (t *Task) Run() error {
	t.Status = StatusRunning
	fmt.Println("task " + t.Name + " has started")
	err := t.Action()
	if err != nil {
		t.Status = StatusFailed
		return fmt.Errorf("task %s failed: %s", t.Name, err)
	}
	t.Status = StatusCompleted
	fmt.Println("task " + t.Name + " has completed")
	return nil
}

func (w *Workflow) Run() error {
	for i := range w.Tasks {
		err := w.Tasks[i].Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	workflow := Workflow{
		Name: "media-ingestion",
		Tasks: []Task{
			{"fetch", StatusNotStarted, fetch},
			{"parse", StatusNotStarted, parse},
			{"download", StatusNotStarted, download},
		},
	}
	fmt.Println("workflow engine started")
	workflow.Print()
	err := workflow.Run()
	if err != nil {
		fmt.Println("workload failed: ", err)
	}
	workflow.Print()
}

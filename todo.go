package todo

import {
	"time"
	"errors"
}

type item struct{
	Task string
	Done bool
	CreatedAt time.Time	
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	var todo item := item {
		Task: task,
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	var ls Todos = *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {	
	var ls Todos = *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

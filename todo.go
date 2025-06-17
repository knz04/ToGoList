package todo

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
)

type item struct{
	Task string
	Done bool
	CreatedAt time.Time	
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	var todo item = item {
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
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}

func (t *Todos) Load(filename string) error {
	var file []byte
	var err error

	file, err = os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}
	
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	var data []byte
	var err error

	data, err = json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (t *Todos) Print() {
	symbols := tw.NewSymbolCustom("Animal").
		WithRow("~").
		WithColumn("|").
		WithTopLeft("🐾").
		WithTopMid("🦴").
		WithTopRight("🐾").
		WithMidLeft("🐟").
		WithCenter("🐱").
		WithMidRight("🐟").
		WithBottomLeft("🐾").
		WithBottomMid("🦴").
		WithBottomRight("🐾")

	table := tablewriter.NewTable(os.Stdout, tablewriter.WithRenderer(renderer.NewBlueprint(tw.Rendition{Symbols: symbols})))
	table.Header([]string{"No.", "Task", "Done", "Created At", "Completed At"})

	for i, item := range *t {
		done := "❌"
		if item.Done {
			done = "✅"
		}

		completedAt := "Not yet completed"
		if !item.CompletedAt.IsZero() {
			completedAt = item.CompletedAt.Format(time.RFC3339)
		}

		row := []string{
			strconv.Itoa(i + 1),
			item.Task,
			done,
			item.CreatedAt.Format(time.RFC3339),
			completedAt,
		}

		table.Append(row)
	}

	table.Render()
}

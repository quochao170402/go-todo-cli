package data

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TodoList []Task

func (tasks TodoList) Display() {
	if len(tasks) == 0 {
		fmt.Println("\033[33mNo tasks available.\033[0m") // Yellow message
		return
	}

	// Create tabwriter for alignment
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Print header
	fmt.Fprintln(w, "ID\tStatus\tTitle\tDescription\tCreated At\tUpdated At")

	// Print separator line
	fmt.Fprintln(w, strings.Repeat("-", 100))

	// Print each task
	for _, task := range tasks {
		status := "\033[31m❌\033[0m" // red for not done
		if task.Done {
			status = "\033[32m✅\033[0m" // green for done
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
			task.Id.String(),
			status,
			task.Title,
			task.Description,
			task.CreatedAt.Format("2006-01-02 15:04"),
			task.UpdatedAt.Format("2006-01-02 15:04"),
		)
	}

	w.Flush()
}

func (tasks *TodoList) AddTask(title, description string) {
	todo := Task{
		Id:          uuid.New(),
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	*tasks = append(*tasks, todo)
}

func (tasks *TodoList) GetTasks() []Task {
	return *tasks
}

func (tasks *TodoList) MarkTaskDone(id uuid.UUID) {
	for i, task := range *tasks {
		if task.Id == id {
			(*tasks)[i].Done = true
			(*tasks)[i].UpdatedAt = time.Now()
			break
		}
	}
}

func (tasks *TodoList) RemoveTask(id uuid.UUID) {
	for i, task := range *tasks {
		if task.Id == id {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			break
		}
	}
}

func (tasks *TodoList) UpdateTask(id uuid.UUID, title, description string) {
	for i, task := range *tasks {
		if task.Id == id {
			(*tasks)[i].Title = title
			(*tasks)[i].Description = description
			(*tasks)[i].UpdatedAt = time.Now()
			break
		}
	}
}

func (tasks *TodoList) GetTaskById(id uuid.UUID) (*Task, error) {
	for _, task := range *tasks {
		if task.Id == id {
			return &task, nil
		}
	}
	return nil, errors.New("Task " + id.String() + " not found!")
}

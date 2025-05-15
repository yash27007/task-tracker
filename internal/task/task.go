package task

import (
	"fmt"
	"time"
)

// Task struct already defined
type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Status      string    `json:"status"` // "todo", "in-progress", "done"
}

// AddTask adds a task with description
func AddTask(description string) {
	err := AddNewTask(description)
	if err != nil {
		fmt.Println("Error adding task:", err)
		return
	}
	fmt.Printf("Added the task: %s\n", description)
}

// UpdateTask modifies the description of a task
func UpdateTask(id int, newDescription string) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	updated := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now()
			updated = true
			break
		}
	}

	if !updated {
		fmt.Printf("Task with ID %d not found\n", id)
		return
	}

	err = SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving updated task:", err)
		return
	}

	fmt.Printf("Updated the task with the ID: %d as %s\n", id, newDescription)
}

// DeleteTask removes a task by ID
func DeleteTask(id int) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	newTasks := []Task{}
	found := false

	for _, task := range tasks {
		if task.ID == id {
			found = true
			continue
		}
		newTasks = append(newTasks, task)
	}

	if !found {
		fmt.Printf("Task with ID %d not found\n", id)
		return
	}

	err = SaveTasks(newTasks)
	if err != nil {
		fmt.Println("Error saving tasks after deletion:", err)
		return
	}

	fmt.Printf("Deleted the task with the ID: %d\n", id)
}

// MarkInProgress sets a task status to "in-progress"
func MarkInProgress(id int) {
	updateStatus(id, "in-progress")
}

// MarkDone sets a task status to "done"
func MarkDone(id int) {
	updateStatus(id, "done")
}

// updateStatus is a helper function to change status
func updateStatus(id int, newStatus string) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	updated := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = newStatus
			tasks[i].UpdatedAt = time.Now()
			updated = true
			break
		}
	}

	if !updated {
		fmt.Printf("Task with ID %d not found\n", id)
		return
	}

	err = SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving task status:", err)
		return
	}

	fmt.Printf("Updated the status of task %d to '%s'\n", id, newStatus)
}

// ListAllTasks prints all tasks from the file
func ListAllTasks() {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	for _, task := range tasks {
		fmt.Printf("ID: %d\nDescription: %s\nStatus: %s\nCreated: %s\nUpdated: %s\n\n",
			task.ID, task.Description, task.Status, task.CreatedAt.Format(time.RFC822), task.UpdatedAt.Format(time.RFC822))
	}
}

// ListTaskByStatus shows tasks filtered by status
func ListTaskByStatus(status string) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	count := 0
	for _, task := range tasks {
		if task.Status == status {
			fmt.Printf("ID: %d\nDescription: %s\nCreated: %s\nUpdated: %s\n\n",
				task.ID, task.Description, task.CreatedAt.Format(time.RFC822), task.UpdatedAt.Format(time.RFC822))
			count++
		}
	}

	if count == 0 {
		fmt.Printf("No tasks with status '%s'\n", status)
	}
}

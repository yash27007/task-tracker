package main

import (
	"fmt"
	"os"
	"strconv"
	"task-tracker/internal/task"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		printHelp()
		return
	}

	command := args[1]

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Usage: task-cli add \"task description\"")
			return
		}
		description := args[2]
		task.AddTask(description)

	case "update":
		if len(args) < 4 {
			fmt.Println("Usage: task-cli update [id] \"new description\"")
			return
		}
		id := getId(args)
		if id == -1 {
			return
		}

		newDesc := args[3]
		task.UpdateTask(id, newDesc)

	case "delete":
		if len(args) < 3 {
			fmt.Println("Usage: task-cli delete [id]")
			return
		}
		id := getId(args)
		if id == -1 {
			return
		}
		task.DeleteTask(id)

	case "mark-in-progress":
		if len(args) < 3 {
			fmt.Println("Usage: task-cli mark-in-progress 3")
			return
		}
		id := getId(args)
		if id == -1 {
			return
		}
		task.MarkInProgress(id)

	case "mark-done":
		if len(args) < 3 {
			fmt.Println("Usage: task-cli done 4")
			return
		}
		id := getId(args)
		if id == -1 {
			return
		}
		task.MarkDone(id)
	case "list":
		if len(args) == 2 {
			task.ListAllTasks()
		} else if args[2] == "in-progress" {
			task.ListTaskByStatus("in-progress")
		} else if args[2] == "done" {
			task.ListTaskByStatus("done")
		} else if args[2] == "todo" {
			task.ListTaskByStatus("todo")
		}else{
			printHelp()
			return;
		}

	default:
		fmt.Println("Unknown command:", command)
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  task-cli add \"description\"")
	fmt.Println("  task-cli update [id] \"new description\"")
	fmt.Println("  task-cli delete [id]")
	fmt.Println("  task-cli mark-in-progress [id]")
	fmt.Println("  task-cli mark-done [id]")
	fmt.Println("  task-cli list [todo|done|in-progress]")
	fmt.Println("  task-cli list -- this is to list all the tasks")
}

func getId(args []string) int {
	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid ID:", args[2])
		return -1
	}
	return id
}

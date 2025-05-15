# 🧾 Task Tracker (Built with Go)

**Task Tracker** is a simple, zero-dependency command-line application to help you manage your daily tasks efficiently. Built in **Go**, it stores your tasks locally in a JSON file and allows you to **add**, **update**, **delete**, and **track statuses** like *todo*, *in-progress*, and *done* — all from your terminal.

---

## 🚀 Features

* ✅ Add new tasks with descriptions
* ✏️ Update task descriptions
* ❌ Delete tasks
* 🔄 Mark tasks as `in-progress` or `done`
* 📋 List all tasks or filter by status: `todo`, `in-progress`, `done`
* 📁 Stores tasks persistently in a local `tasks.json` file
* 🧱 No external libraries or frameworks used — just the Go standard library

---

## 📦 Installation

1. Make sure you have **Go installed**: [Install Go](https://golang.org/doc/install)
2. Clone the repository:

```bash
git clone https://github.com/yash27007/task-tracker.git
cd task-tracker-cli
```

3. Build the CLI:

```bash
go build -o task-cli
```

This creates a binary named `task-cli`.

---

## 🛠 Usage

The CLI uses **positional arguments** to perform different operations.

### ➕ Add a Task

```bash
./task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

### ✏️ Update a Task

```bash
./task-cli update 1 "Buy groceries and cook dinner"
# Output: Task updated successfully
```

### ❌ Delete a Task

```bash
./task-cli delete 1
# Output: Task deleted successfully
```

### 🔄 Mark as In-Progress

```bash
./task-cli mark-in-progress 1
# Output: Task marked as in-progress
```

### ✅ Mark as Done

```bash
./task-cli mark-done 1
# Output: Task marked as done
```

### 📋 List All Tasks

```bash
./task-cli list
```

### 📌 List Tasks by Status

```bash
./task-cli list todo
./task-cli list in-progress
./task-cli list done
```

---

## 📂 Task Structure

Each task is stored in `tasks.json` and has the following structure:

```json
{
  "id": 1,
  "description": "Buy groceries",
  "status": "todo",
  "createdAt": "2025-05-15T14:34:00Z",
  "updatedAt": "2025-05-15T14:34:00Z"
}
```

---

## 📁 File Storage

* All tasks are stored in a JSON file named `tasks.json` in the **current working directory**.
* The file is automatically created if it does not exist.

---

## ⚙️ Error Handling

* Gracefully handles invalid commands, missing files, incorrect IDs, and malformed inputs.
* Provides user-friendly error messages and status updates.

---

## 🧪 Example Workflow

```bash
./task-cli add "Write documentation"
./task-cli add "Fix bug #102"
./task-cli list
./task-cli mark-in-progress 2
./task-cli mark-done 1
./task-cli list done
./task-cli delete 1
```

---

## 📚 Learning Outcomes

This project is ideal for Go beginners and helps you practice:

* Working with the filesystem
* Reading/writing JSON
* Creating command-line tools
* Handling CLI arguments
* Error handling and validation

---

## 🤝 Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

---

## 📝 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## 💡 Inspiration
https://roadmap.sh/projects/task-tracker
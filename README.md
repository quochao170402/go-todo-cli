# Go Todo List CLI

A simple command-line **Todo List application** written in Go.  
This project is designed for learning **Go basics, CLI development, and clean code practices** while building something practical.

---

## 📌 Purpose

The main goal of this project is to:

- Practice **Go programming fundamentals** (structs, slices, JSON, time handling)
- Learn **CLI argument parsing** and command handling
- Work with **UUIDs** for unique task identifiers
- Handle **file persistence** for storing tasks
- Build a **user-friendly CLI interface** with readable table output
- Explore **Go modules & project structure** for real-world applications

---

## 🧠 What You'll Learn

- **Structs & Methods** — defining and working with custom types like `Task` and `TodoList`
- **JSON Marshalling/Unmarshalling** — saving and loading data from a file
- **Command Parsing** — interpreting arguments to run different commands (`add`, `update`, `list`, etc.)
- **Formatting Output** — using Go's `text/tabwriter` for a nice console table
- **Aliases** — supporting both long and short CLI commands (e.g., `list` and `ls`)
- **Error Handling** — writing clear and safe error messages for invalid input

---

## 📦 Installation

```bash
git clone https://github.com/quochao170402/go-todo-cli.git
cd go-todo-cli
go mod tidy
```

## Example

```pwsh
go run . add "Setup Project" "Initialize Git repository and Go module"
go run . --a "Learn Golang" "Read Go basics and practice small programs"
go run . --a "Build CLI" "Implement add, list, update, and delete commands"
go run . --a "Test CLI" "Add sample tasks to test all functions"
go run . --a "Publish" "Push code to GitHub"
```

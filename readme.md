## CLI Todo – TUI-powered task viewer

A small cross-platform CLI app for viewing and exploring tasks with a modern TUI built on Bubble Tea. It ships with three commands: list tasks, view task details, and open a create form UI.

### Features

- **List tasks**: Interactive table with keyboard navigation
- **Task details**: Pretty, colorized detail view
- **Create form UI**: Inputs for Task Name, Status, Priority, Due Date, and Description
- **Fast, keyboard-first UX** powered by Bubble Tea, Bubbles, and Lip Gloss

### Tech stack

- Go `go.mod` module: `cli/todo`
- Cobra (CLI), Bubble Tea/Bubbles (TUI), Lip Gloss (styles)

### Requirements

- Go (version matching `go.mod`, currently `1.24.2`)
- Windows, macOS, or Linux terminal

### Install / Build

```bash
# from the project root
go mod download
go build -o todo.exe .   # on Windows
# or
go build -o todo .       # on macOS/Linux
```

You can also run without building:

```bash
go run . --help
```

### Usage

```bash
# Show help
todo.exe --help

# List all tasks in a TUI table
todo.exe list

# View details for a task by ID or title
todo.exe details <id-or-title>

# Open the create form UI (inputs only, no persistence yet)
todo.exe create
```

### Commands

- **list (l)**: Get the list of all tasks and browse them in a table.
- **details (d)**: View details about a single task by ID or title.
- **create (c)**: Open a form to enter Task Name, Status, Priority, Due Date, and Description.

### TUI navigation

- **Global**: `q`, `esc`, or `ctrl+c` to quit

- **List view**

  - Move selection: `up`/`down` or `k`/`j`
  - Open the selected task: `enter`

- **Details view**

  - Close: `q`, `esc`, or `ctrl+c`

- **Create view**
  - Move focus: `tab`, `shift+tab`, `up`, `down`
  - Fields: Task Name, Status, Priority, Due Date, Description
  - Close: `q`, `esc`, or `ctrl+c`

Note: The create form currently renders inputs and supports navigation; saving/persistence is not implemented yet.

### Data

For demo purposes, tasks are read from `cmd/tasks.json`. The list and details views consume these records. The create form is UI-only in this version.

### Development

```bash
# Run in watch mode (example using `reflex`, optional)
reflex -r '\\.go$' -s -- go run . list

# Linting (choose your preferred tool)
golangci-lint run
```

Project structure highlights:

- `cmd/`: Cobra commands and helpers (`list`, `details`, `create`)
- `tui/`: Bubble Tea models for List, Details, and Create views
- `models/`: Task model definitions

### License

MIT. See `LICENSE`.

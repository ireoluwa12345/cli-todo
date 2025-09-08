# Go CLI Todo List

This is a simple command-line Todo List application written in Go. It allows you to add, view, and manage your tasks, storing them in a CSV file in your home directory.

## Features

- **Add Tasks:** Quickly add new tasks to your todo list.
- **View Tasks:** Display all your tasks in a formatted table.
- **Persistent Storage:** Tasks are saved in a CSV file (`todo.csv`) in your home directory.
- **Colorful Output:** Uses colored output for better readability.
- **Timestamps:** Each task is saved with the time it was added.

## Getting Started

### Prerequisites

- Go 1.18 or higher installed on your system.

### Installation

1. Clone the repository:

   ```
   git clone https://github.com/yourusername/go-todo-cli.git
   cd go-todo-cli
   ```

2. Install dependencies:

   ```
   go get github.com/fatih/color
   go get github.com/rodaine/table
   ```

3. Build the application:
   ```
   go build -o todo
   ```

### Usage

- **Add a Task:**

  ```
  ./todo add "Buy groceries"
  ```

- **View Tasks:**

  ```
  ./todo list
  ```

- **Remove a Task:**
  ```
  ./todo remove 1
  ```

> The CSV file is automatically created in your home directory as `todo.csv`.

## Project Structure

- `main.go` - Main application logic.
- `todo.csv` - CSV file storing your tasks (auto-generated).

## Dependencies

- [fatih/color](https://github.com/fatih/color) - For colored terminal output.
- [rodaine/table](https://github.com/rodaine/table) - For pretty table formatting.

## License

This project is licensed under the MIT License.

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## Author

- [Your Name](https://github.com/yourusername)

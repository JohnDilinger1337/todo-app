# Todo App

A simple command-line todo application built with Go that organizes tasks into multiple categories.

## Features

- Organize todos into three categories
- Mark todos as done or incomplete
- Display todos in a formatted, easy-to-read layout

## Project Structure

```
todo-app/
├── main.go              # Main application entry point
├── go.mod              # Go module definition
├── models/
│   └── todo.go         # Todo data model
├── services/
│   └── todo.go         # Todo service logic
└── utility/
    ├── formatters.go   # Formatting utilities
    └── helpers.go      # Helper functions
```

## Getting Started

### Prerequisites

- Go 1.x or higher

### Installation

1. Clone the repository
2. Navigate to the project directory
3. Run the application:

```bash
go run main.go
```

## Usage

The application displays a formatted list of todos organized by category:

```
############# Welcome to my to-do list #############
Category1 {
1) Learn Go (pending)
2) Deploy to cloud (pending)
}
Category2 {
1) Build a web app (pending)
}
Category3 {
1) Fix bugs (pending)
}
```

## Development

To run the application in watch mode (auto-reload on file changes):

```bash
air
```

## License

MIT

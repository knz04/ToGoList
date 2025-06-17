# ToGoList 📝
ToGoList is a CLI to do list application written in Go.

## Features

* **Add** new tasks to your list.
* **List** all tasks in a stylish, custom-themed table.
* **Mark** tasks as complete.
* **Delete** tasks from the list.
* **Persistent Storage:** Automatically saves and loads your list from a `.todos.json` file in the same directory.

## Demonstration

Here is an example of the output using the custom animal-themed table style:

```
🐾~~~~~|~~~~~~~~~~~~~~~~~~|~~~~~~|~~~~~~~~~~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~~~~~~~~~🐾
| NO . |       TASK       | DONE |        CREATED AT         |        COMPLETED AT        |
🐟~~~~~|~~~~~~~~~~~~~~~~~~|~~~~~~|~~~~~~~~~~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~~~~~~~~~🐟
| 1    | Learn Go         | ❌   | 2025-06-17T10:00:00+07:00 | Not yet completed          |
| 2    | Build a CLI App  | ✅   | 2025-06-17T11:30:00+07:00 | 2025-06-17T15:45:00+07:00 |
| 3    | Walk the dog     | ❌   | 2025-06-17T18:00:00+07:00 | Not yet completed          |
🐾~~~~~|~~~~~~~~~~~~~~~~~~|~~~~~~|~~~~~~~~~~~~~~~~~~~~~~~~~~~|~~~~~~~~~~~~~~~~~~~~~~~~~~🐾
```

## Installation

To build the application, clone the repository and run the following command from the project root. This will create an executable file named `togo`.

```bash
go build -o togo ./cmd/togo
```

## Usage

All commands are passed as flags to the `togo` executable.

### List All To-Dos

To display all the tasks in your list.

```bash
./togo -list
```

### Add a New To-Do

Use the `-add` flag followed by the task description.

```bash
./togo -add "Write the project README"
```

### Complete a To-Do

Use the `-complete` flag followed by the task number from the list.

```bash
# Marks task number 2 as complete
./togo -complete 2
```

### Delete a To-Do

Use the `-del` flag followed by the task number to remove it from the list.

```bash
# Deletes task number 3
./togo -del 3
```

### Show Help

Use the `-h` or `-help` flag to display a list of all available commands and their descriptions.

```bash
./togo -h
```

## Libraries
https://github.com/olekukonko/tablewriter?tab=readme-ov-file

## Reference
https://www.youtube.com/watch?v=j1CXoOQXbco

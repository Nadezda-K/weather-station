# Weather Station

A Go learning project for parsing sensor messages and maintaining state in memory.

## Task

Build a program that remembers the latest readings from one weather station. Each partial update changes one sensor, while a query prints the complete known state.

## What the program does

- Reads commands and sensor updates from standard input.
- Accepts updates in the form `<id>,<value>`, such as `1,21.5`.
- Prints all nine readings when given `get`.
- Resets every reading to `NULL` when given `clear`.
- Stops when given `exit`.
- Stores missing readings and values that cannot be parsed as numbers as `NULL`.

| Sensor ID | Reading |
| --- | --- |
| 1 | Air temperature |
| 2 | Air pressure |
| 7 | Precipitation |
| 11 | Wind speed |
| 12 | Wind direction |
| 13 | Humidity |
| 14 | Dew point |
| 15 | Soil moisture |
| 22 | Cloud cover |

## Run

With Go installed, run from this folder:

```bash
go run main.go
```

Example input:

```text
1,21.5
13,65
get
clear
exit
```

Use the comma-separated format for every update: the current parser expects both an ID and a value. Readings last only until the program exits.

## Learning focus

Structs, pointers, state updates, string splitting, numeric parsing, switch statements, and formatted output.

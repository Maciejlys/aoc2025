# Advent of Code 2025 - Go Solutions

This repository contains my solutions for [Advent of Code 2025](https://adventofcode.com/2025) written in Go.

## Structure

```
.
├── main.go              # Main runner for all solutions
├── utils/               # Common utilities and helpers
│   ├── input.go        # File reading and parsing utilities
│   ├── math.go         # Mathematical helper functions
│   ├── grid.go         # 2D grid and point utilities
│   └── collections.go  # Data structures (Set, Queue, Stack)
├── day01/              # Solution for day 1
│   ├── day01.go        # Implementation
│   ├── day01_test.go   # Tests
│   ├── input.txt       # Puzzle input (gitignored)
│   └── README.md       # Notes and approach
└── ...                 # More days
```

## Getting Started

### Prerequisites

- Go 1.21 or later

### Installation

1. Clone the repository:
```bash
git clone https://github.com/Maciejlys/aoc2025.git
cd aoc2025
```

2. Download your puzzle input from [Advent of Code](https://adventofcode.com/2025) and save it to the appropriate day folder (e.g., `day01/input.txt`)

### Running Solutions

Run a specific day:
```bash
go run main.go -day 1
```

Run with a custom input file:
```bash
go run main.go -day 1 -input path/to/input.txt
```

Build the binary:
```bash
go build -o aoc2025
./aoc2025 -day 1
```

### Running Tests

Run all tests:
```bash
go test ./...
```

Run tests for a specific day:
```bash
go test ./day01
```

Run tests with verbose output:
```bash
go test -v ./...
```

## Utils Package

The `utils` package provides commonly used functions for Advent of Code:

### File I/O
- `ReadLines(filepath)` - Read file as slice of strings
- `ReadFile(filepath)` - Read entire file as string
- `ReadInts(filepath)` - Read file as slice of integers
- `SplitByEmptyLines(lines)` - Split input into groups by empty lines
- `ParseIntList(s, separator)` - Parse comma/space-separated integers

### Math
- `Abs(x)` - Absolute value
- `Min(a, b)` / `Max(a, b)` - Min/max of two integers
- `Sum(nums)` - Sum of slice
- `Product(nums)` - Product of slice
- `Atoi(s)` - String to int with panic on error
- `GCD(a, b)` / `LCM(a, b)` - Greatest common divisor / Least common multiple

### Grid & Geometry
- `Point{X, Y}` - 2D point with methods:
  - `Add(other)` / `Sub(other)` - Vector operations
  - `ManhattanDistance(other)` - Manhattan distance
  - `Neighbors4()` - 4 orthogonal neighbors
  - `Neighbors8()` - 8 neighbors including diagonals
- Direction constants: `Up`, `Down`, `Left`, `Right`

### Data Structures
- `Set[T]` - Generic set implementation
- `Queue[T]` - Generic FIFO queue
- `Stack[T]` - Generic LIFO stack

## Creating a New Day

To add a solution for a new day (e.g., day 2):

1. Create the directory:
```bash
mkdir day02
```

2. Copy the template files from day01:
```bash
cp day01/day01.go day02/day02.go
cp day01/day01_test.go day02/day02_test.go
cp day01/README.md day02/README.md
cp day01/input.txt.example day02/input.txt.example
```

3. Update package names and imports in the copied files

4. Add the case to `main.go`:
```go
case 2:
    day02.Run(inputPath)
```

5. Download your input to `day02/input.txt` and start solving!

## License

This project is open source and available under the MIT License.

## Acknowledgments

- [Advent of Code](https://adventofcode.com/) by Eric Wastl
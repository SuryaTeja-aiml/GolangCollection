# Go Programming Guides: Inputs and Packages

Welcome! This repository contains comprehensive, well-documented code examples and explanations for two fundamental concepts in Go (Golang):
1. **Handling User Inputs** (Basic and Special data structures)
2. **The `main` Package** and its role in Go executables

---

## 1. Handling Inputs (`inputs/` folder)

Go provides several built-in mechanisms to read user inputs from standard input (`os.Stdin`). We have implemented two distinct examples to cover different use cases:

### A. Basic Inputs (`inputs/basic.go`)
This file demonstrates how to read standard, primitive data types.
* **Integer input** using `fmt.Scan`.
* **Float input** using format specifiers with `fmt.Scanf`.
* **Single Word string input** using `fmt.Scan`.
* **Line-by-line / Full sentence reading** using `bufio.NewScanner` and `scanner.Scan()`.
* **Robust input parsing** using `strconv` to safely convert string inputs into target types, which is the standard way to handle real-world CLI application inputs safely.

**How to run:**
```bash
go run inputs/basic.go
```

### B. Special and Complex Inputs (`inputs/special.go`)
In many competitive programming scenarios, data structures questions, or advanced CLI tools, you need to parse more complex formats:
* **Array / Slice of defined size**: Parsing $N$ elements from a single space-separated line.
* **Specific parts of an array**: Slicing an array using boundary conditions (e.g., getting elements from index $i$ to $j$).
* **Set Simulation**: Since Go doesn't have a built-in Set data structure, we simulate a Set using a `map[string]struct{}`.
* **Specific parts / Filtered Set**: Filtering set elements based on specific criteria (e.g., matching a given string prefix).

**How to run:**
```bash
go run inputs/special.go
```

---

## 2. The `main` Package (`packages/` folder)

In Go, package organization is key. The `main` package has a highly specific, unique role.

### Main Package Demo (`packages/main_pkg_demo.go`)
This file details all key aspects and use cases of the `main` package:
1. **Defining an Executable**: Any file starting with `package main` instructs the Go compiler to build an executable binary rather than a library package.
2. **The `main()` function**: The central entrypoint of any Go executable. It accepts no parameters and returns nothing.
3. **The `init()` function**:
   * Runs *before* the `main()` function is called.
   * Useful for initializing configuration, starting background connections, or checking environment variables.
   * You can define multiple `init()` functions; they execute sequentially in the order of definition.
4. **Command Line Arguments (`os.Args`)**: Accessing arguments supplied to your executable at runtime.
5. **Program Exiting (`os.Exit`)**: Terminating the program with a specific exit status code.

**How to run:**
```bash
go run packages/main_pkg_demo.go argument1 argument2
```

---

## Summary of Project Directory Structure

```
.
├── inputs/
│   ├── basic.go         # Standard data types (int, float, string) input examples
│   └── special.go       # Slices, subsets, and Set simulation inputs
└── packages/
    └── main_pkg_demo.go # Complete guide & code on the special 'main' package
```

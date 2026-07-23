package main

import (
	"fmt"
	"os"
)

/*
Use Case 1: Package main is special in Go.
In Go, the `package main` statement tells the Go compiler that this is an executable program rather than a shared library.
An executable must have a `main` package and a `main` function as the entry point.

Use Case 2: The `main()` function.
The `main` function is the entry point of the executable. It takes no arguments and returns no values.
Once the `main` function finishes executing, the program exits.

Use Case 3: The `init()` function.
Go supports an `init` function in any package, including the `main` package.
The `init()` function runs BEFORE the `main()` function is called.
It is commonly used to:
  - Initialize package-level variables.
  - Check/verify environment setup before execution begins.
  - Perform one-time setup tasks.
*/

// Package-level variable initialized manually or within init()
var programConfig string

// init() is automatically executed by Go run/build before main()
func init() {
	fmt.Println("[Init Phase] Initializing program configuration...")
	programConfig = "Default Config Value"

	// You can also access command-line args or environment variables in init()
	if len(os.Args) > 1 {
		fmt.Printf("[Init Phase] Received %d command line arguments.\n", len(os.Args)-1)
	}
}

// Another init() function. Go allows multiple init() functions, which execute in the order of appearance.
func init() {
	fmt.Println("[Init Phase 2] Running secondary init tasks...")
}

func main() {
	fmt.Println("\n--- Use Case 1: Entry Point Execution ---")
	fmt.Println("This is the main() function running. Since this file is in 'package main', it is compiled as an executable.")

	fmt.Println("\n--- Use Case 2: Package-Level Variables initialized by init() ---")
	fmt.Printf("Program Config value: %s\n", programConfig)

	fmt.Println("\n--- Use Case 3: Interacting with Command Line Arguments ---")
	// os.Args[0] is the path of the program itself.
	// os.Args[1:] are the custom arguments provided to the program.
	fmt.Printf("Program Name: %s\n", os.Args[0])
	if len(os.Args) > 1 {
		fmt.Println("Arguments supplied to main package:")
		for i, arg := range os.Args[1:] {
			fmt.Printf("  Arg %d: %s\n", i+1, arg)
		}
	} else {
		fmt.Println("No arguments supplied. Run with: go run packages/main_pkg_demo.go <arg1> <arg2>")
	}

	fmt.Println("\n--- Use Case 4: Graceful/Explicit Exits ---")
	fmt.Println("By default, returning from main() exits with status code 0.")
	fmt.Println("To exit with a non-zero status code, use os.Exit(code).")
	// Uncommenting below would immediately terminate the program with status 1
	// os.Exit(1)
}

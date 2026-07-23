package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// This program demonstrates how to read basic data types (int, float, string, full line) in Go.
func main() {
	fmt.Println("--- Go Basic Input Demonstration ---")

	// 1. Reading an Integer using fmt.Scan
	// fmt.Scan reads space-separated values.
	var integerVal int
	fmt.Print("Enter an integer: ")
	_, err := fmt.Scan(&integerVal)
	if err != nil {
		fmt.Printf("Error reading integer: %v\n", err)
	} else {
		fmt.Printf("You entered integer: %d\n", integerVal)
	}

	// 2. Reading a Float using fmt.Scanf
	// fmt.Scanf parses input according to a format specifier.
	var floatVal float64
	fmt.Print("Enter a float: ")
	_, err = fmt.Scanf("%f\n", &floatVal) // Read float and consume the newline
	if err != nil {
		// If previous Scan left a newline, Scanf might fail or behave unexpectedly.
		// Let's clear the buffer or handle it. Here we try to handle it.
		fmt.Printf("Error reading float (or skipped): %v\n", err)
	} else {
		fmt.Printf("You entered float: %.4f\n", floatVal)
	}

	// 3. Reading a Single Word (String) using fmt.Scan
	var word string
	fmt.Print("Enter a single word: ")
	_, err = fmt.Scan(&word)
	if err != nil {
		fmt.Printf("Error reading word: %v\n", err)
	} else {
		fmt.Printf("You entered word: %s\n", word)
	}

	// Clear remaining input buffer before using bufio
	// Often mixing fmt.Scan and bufio can cause bufio to read the leftover newline.
	// To safely demonstrate bufio, we set up a scanner.
	fmt.Println("\n--- Reading Full Line using bufio.Scanner ---")
	fmt.Print("Enter a full line of text: ")

	// Create a new scanner for standard input
	scanner := bufio.NewScanner(os.Stdin)
	// Consuming the rest of the line if there was any leftover newline
	if scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("You entered line: %q\n", line)
	}

	// 4. Detailed demonstration of parsing from a full line (robust approach)
	fmt.Println("\n--- Robust Parsing from Line Input ---")
	fmt.Print("Enter an integer and a float separated by space: ")
	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			parsedInt, err1 := strconv.Atoi(parts[0])
			parsedFloat, err2 := strconv.ParseFloat(parts[1], 64)
			if err1 == nil && err2 == nil {
				fmt.Printf("Parsed Int: %d, Parsed Float: %f\n", parsedInt, parsedFloat)
			} else {
				fmt.Printf("Error parsing: int_err=%v, float_err=%v\n", err1, err2)
			}
		} else {
			fmt.Println("Expected at least two space-separated values.")
		}
	}
}

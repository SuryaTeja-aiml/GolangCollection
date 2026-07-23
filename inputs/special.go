package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// This program demonstrates how to read and manipulate complex/special inputs in Go:
// 1. Array/Slice of a specific size
// 2. Reading specific parts of an array (e.g., indices i to j)
// 3. Simulating a Set (handling unique elements only)
// 4. Reading a specific range/part of a Set or filtering set elements
func main() {
	fmt.Println("--- Go Special Input/Data Structure Demonstration ---")
	scanner := bufio.NewScanner(os.Stdin)

	// -------------------------------------------------------------
	// 1. Reading an Array/Slice of a specific size (N elements)
	// -------------------------------------------------------------
	fmt.Println("\n1. [Slice Input] Enter number of elements followed by the elements (space-separated):")
	fmt.Print("Example: 5 10 20 30 40 50\nInput: ")

	var n int
	var elements []int

	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) > 0 {
			var err error
			n, err = strconv.Atoi(parts[0])
			if err != nil {
				fmt.Println("Invalid size:", err)
				return
			}

			// Parse the rest of the inputs up to N elements
			for i := 1; i < len(parts) && len(elements) < n; i++ {
				val, err := strconv.Atoi(parts[i])
				if err == nil {
					elements = append(elements, val)
				}
			}
		}
	}
	fmt.Printf("Parsed Slice (size %d): %v\n", n, elements)

	// -------------------------------------------------------------
	// 2. Reading a specific part of an Array (slicing/bounds)
	// -------------------------------------------------------------
	fmt.Println("\n2. [Slice Slicing] Extract a specific sub-array/slice (e.g., from index i to j - exclusive):")
	fmt.Print("Enter start and end index (e.g., '1 4' for elements 1 to 3): ")
	if scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) >= 2 {
			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				// Check bounds
				if start >= 0 && end <= len(elements) && start <= end {
					subSlice := elements[start:end]
					fmt.Printf("Sub-slice elements[%d:%d]: %v\n", start, end, subSlice)
				} else {
					fmt.Printf("Index out of bounds or invalid range. Current slice length is %d.\n", len(elements))
				}
			}
		}
	}

	// -------------------------------------------------------------
	// 3. Simulating a Set (Uniqueness) from Input
	// Go does not have a built-in Set, so we use a map[T]bool or map[T]struct{}.
	// -------------------------------------------------------------
	fmt.Println("\n3. [Set Input] Enter space-separated elements (duplicates will be filtered out):")
	fmt.Print("Example: apple banana apple orange banana\nInput: ")

	stringSet := make(map[string]struct{})
	var uniqueList []string // to keep insertion order if needed

	if scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		for _, item := range parts {
			if _, exists := stringSet[item]; !exists {
				stringSet[item] = struct{}{}
				uniqueList = append(uniqueList, item)
			}
		}
	}

	fmt.Println("Unique Elements (Set Keys):")
	for key := range stringSet {
		fmt.Printf(" - %s\n", key)
	}
	fmt.Printf("Unique elements in insertion order: %v\n", uniqueList)

	// -------------------------------------------------------------
	// 4. Getting specific parts/subsets of a Set
	// E.g., getting only elements starting with a prefix, or a range in order
	// -------------------------------------------------------------
	fmt.Println("\n4. [Special Part of Set] Filter the Set (e.g., items starting with a specific prefix):")
	fmt.Print("Enter a prefix to filter: ")
	if scanner.Scan() {
		prefix := strings.TrimSpace(scanner.Text())
		filtered := []string{}
		for item := range stringSet {
			if strings.HasPrefix(item, prefix) {
				filtered = append(filtered, item)
			}
		}
		fmt.Printf("Filtered Set elements starting with %q: %v\n", prefix, filtered)
	}
}

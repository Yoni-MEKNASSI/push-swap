// checker lit des instructions sur l'entrée standard, les exécute et affiche OK ou KO.
package main

import (
	"bufio"
	"fmt"
	"os"

	"push-swap/internal/stack"
)

func fail() {
	fmt.Fprintln(os.Stderr, "Error")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	nums, ok := stack.Parse(os.Args[1:])
	if !ok {
		fail()
	}
	if len(nums) == 0 {
		return
	}

	s := &stack.Stacks{A: nums}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if line := scanner.Text(); line != "" && !s.Apply(line) {
			fail()
		}
	}

	if err := scanner.Err(); err != nil {
		fail()
	}
	if s.IsSorted() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

package app

import "fmt"

// Run is the entrypoint for hoke. We delegate logic here
// so that it can be tested.
func Run(args []string) int {
	fmt.Println("I got my teeth back.")
	return 0
}

package greet

import "fmt"

// Message builds a greeting string for the given name.
func Message(name string) string {
	if name == "" {
		name = "Gopher"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

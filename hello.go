package hello

import (
	"strings"
	)

// We are declaring the Say function that takes a name string and returns a string
func Say(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}

	return "Hello, " + strings.Join(names, ", ") + "!"
}



// We are declaring the Say function that takes a name string and returns a string
// func Say(name string) string {
// 	return fmt.Sprintf("Hello, %s!", name)
// }
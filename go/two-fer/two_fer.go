// Package twofer allows people to share
package twofer

import "fmt"

// ShareWith provides the name of the lucky person being shared with, otherwise you.
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

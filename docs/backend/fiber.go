//go:build ignore

package main

import (
	"fmt"

	"github.com/charmbracelet/glamour"
)

func main() {
	in := `# Fiber

`

	out, _ := glamour.Render(in, "dark")
	fmt.Print(out)
}

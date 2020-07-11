package main

import (
	"fmt"
	"os"

	fa "github.com/ClarkGuan/fuckandroid"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please insert sub-command!")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "workspace":
		makeWorkspace(os.Args[2:])
	}
}

func makeWorkspace(args []string) {
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Sub-command `workspace` need a `name`")
		os.Exit(1)
	}
	name := args[0]
	dir := "."
	if len(args) > 1 {
		dir = args[1]
	}
	if err := fa.MakeWorkspace(name, dir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

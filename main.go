package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr,
			"Usage: %s source-hg-dir destinate-git-dir\n",
			os.Args[0])
		os.Exit(1)
	}

	if err := Trace(os.Args[1], os.Args[2]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}

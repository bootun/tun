package main

import (
	"fmt"
	"os"

	"github.com/bootun/tun/repl"
)

const VERSION = "0.0.1-preview"

func main() {
	fmt.Printf("Tun %s\n", VERSION)
	repl.Start(os.Stdin, os.Stdout)
}

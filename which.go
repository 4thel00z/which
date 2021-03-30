package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("usage: which <executable-name>")
		os.Exit(1)
	}
	path, err := exec.LookPath(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
	fmt.Print(path)
}

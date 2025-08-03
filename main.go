package main

import (
	"flag"
	"fmt"
	"github.com/mgill25/monkey-go/repl"
	"os"
	"os/user"
	"path"
	"path/filepath"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fileFlagPtr := flag.String("file", "", "Supply a file")
	vmFlagPtr := flag.Bool("vm", true, "VM mode or not")
	flag.Parse()

	if *fileFlagPtr == "" && *vmFlagPtr == false {
		fmt.Printf("Hello %s! This is the new Monkey programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.StartRepl(os.Stdin, os.Stdout, false)
	} else if *fileFlagPtr != "" && *vmFlagPtr == false {
		var filePath string
		if !path.IsAbs(*fileFlagPtr) {
			filePath = filepath.Join(".", *fileFlagPtr)
		} else {
			filePath = *fileFlagPtr
		}
		// read file line by line
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		repl.EvalFile(file, os.Stdout)
	} else if *vmFlagPtr {
		repl.StartRepl(os.Stdin, os.Stdout, true)
	}
}

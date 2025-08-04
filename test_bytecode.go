package main

import (
	"fmt"
	"github.com/mgill25/monkey-go/compiler"
	"github.com/mgill25/monkey-go/lexer"
	"github.com/mgill25/monkey-go/parser"
)

func main() {
	input := "1 + 2"
	
	l := lexer.New(input)
	p := parser.NewParser(l)
	program := p.ParseProgram()
	
	comp := compiler.New()
	err := comp.Compile(program)
	if err != nil {
		fmt.Printf("compilation failed: %s", err)
		return
	}
	
	err = comp.WriteBytecodeToFile("bytecode_output.txt")
	if err != nil {
		fmt.Printf("failed to write bytecode: %s", err)
		return
	}
	
	fmt.Println("Bytecode written to bytecode_output.txt")
}
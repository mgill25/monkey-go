package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/mgill25/monkey-go/ast"
	"github.com/mgill25/monkey-go/evaluator"
	"github.com/mgill25/monkey-go/lexer"
	"github.com/mgill25/monkey-go/object"
	"github.com/mgill25/monkey-go/parser"
)

const PROMPT = "#> "
const MONKEY_FACE = `
            __,__
        .--.  .-"     "-.  .--.
       / .. \/  .-. .-.  \/ .. \
      | |  '|  /   Y   \  |'  | |
      | \   \  \ 0 | 0 /  /   / |
       \ '- ,\.-"""""""-./, -' /
        ''-' /_   ^ ^   _\ '-''
            |  \._   _./  |
            \   \ '~' /   /
             '._ '-=-' _.'
               '-----'
`

func StartRepl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.NewParser(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

// EvalFile reads the entire file, parses it, and evaluates each statement
// individually, printing the result of each expression like the REPL does
func EvalFile(in io.Reader, out io.Writer) {
	content, err := io.ReadAll(in)
	if err != nil {
		fmt.Fprintf(out, "Error reading file: %v\n", err)
		return
	}

	env := object.NewEnvironment()
	l := lexer.New(string(content))
	p := parser.NewParser(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	for _, stmt := range program.Statements {
		evaluated := evaluator.Eval(stmt, env)
		if evaluated != nil {
			// Only print results for expression statements, not let statements
			switch stmt.(type) {
			case *ast.ExpressionStatement:
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\n")
			}
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

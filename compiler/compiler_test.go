package compiler

import (
	"fmt"
	"github.com/mgill25/monkey-go/ast"
	"github.com/mgill25/monkey-go/code"
	"github.com/mgill25/monkey-go/lexer"
	"github.com/mgill25/monkey-go/object"
	"github.com/mgill25/monkey-go/parser"
	"testing"
)

type compilerTestCase struct {
	input                string
	expectedConstants    []any
	expectedInstructions []code.Instructions
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "1 + 2",
			expectedConstants: []any{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
			},
		},
	}

	runCompilerTests(t, tests)
}

func runCompilerTests(t *testing.T, tests []compilerTestCase) {
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)
		compiler := New()
		err := compiler.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}
		bytecode := compiler.Bytecode()
		err = testInstructions(tt.expectedInstructions, bytecode.Instructions)
		if err != nil {
			t.Fatalf("testInstructions failed: %s", err)
		}

		err = testConstants(tt.expectedConstants, bytecode.Constants)
		if err != nil {
			t.Fatalf("testConstants failed: %s", err)
		}

	}
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.NewParser(l)
	return p.ParseProgram()
}

/*
expected: []code.Instructions
actual: code.Instruction
Why is actual flatted where as expected is not?
expected is a slice of slice of bytes -> so we need to flatten it.

Curious - why not produce []code.Instructions as the Bytecode output?
  - Likely because actual bytecode is just a flatted instruction set!
*/
func testInstructions(expected []code.Instructions, actual code.Instructions) error {
	expectedConcat := concatInstructions(expected)
	if len(actual) != len(expectedConcat) {
		return fmt.Errorf("wrong instruction length\nwant=%q\ngot=%q", expectedConcat, actual)
	}
	for i, ins := range expectedConcat {
		// assert that all instructions are exactly equal
		if actual[i] != ins {
			return fmt.Errorf("wrong instruction at %d.\nwant=%q\ngot=%q", i, ins, actual[i])
		}
	}
	return nil
}

func concatInstructions(instructions []code.Instructions) code.Instructions {
	out := code.Instructions{} // which is basically []byte
	for _, ins := range instructions {
		out = append(out, ins...)
	}
	return out
}

func testConstants(expected []any, actual []object.Object) error {
	if len(expected) != len(actual) {
		return fmt.Errorf("wrong number of constants\nwant=%d\ngot=%d", len(expected), len(actual))
	}
	for i, constant := range expected {
		switch constant.(type) {
		case int:
			err := testIntegerObject(int64(constant.(int)), actual[i])
			if err != nil {
				return fmt.Errorf("constant %d - testIntegerObject failed: %s", i, err)
			}
		}
	}
	return nil
}

func testIntegerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("object is not Integer. got=%T(%+v)", actual, actual)
	}
	if result.Value != expected {
		return fmt.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
	}
	return nil
}

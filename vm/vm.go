package vm

import (
	"fmt"
	"github.com/mgill25/monkey-go/code"
	"github.com/mgill25/monkey-go/compiler"
	"github.com/mgill25/monkey-go/object"
)

const StackSize = 2048

type VM struct {
	constants    []object.Object
	instructions code.Instructions
	stack        []object.Object
	sp           int // always points to the next value. Top of stack is stack[sp-1]

}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,
		stack:        make([]object.Object, StackSize),
		sp:           0,
	}
}

func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

func (vm *VM) Run() error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := code.Opcode(vm.instructions[ip])
		fmt.Printf("vm: fetched op=%+v\n", op)
		switch op {
		case code.OpConstant:
			fmt.Printf("vm: OpConstant => decode the operands\n")
			constIndex := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2
			// push the decoded operands on to the stack
			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}
		case code.OpAdd:
			fmt.Printf("vm: OpAdd\n")
			right := vm.pop()
			left := vm.pop()
			leftVal := left.(*object.Integer).Value
			rightVal := right.(*object.Integer).Value
			result := leftVal + rightVal
			vm.push(&object.Integer{Value: result})
		}
	}
	return nil
}

func (vm *VM) push(o object.Object) error {
	fmt.Printf("vm.push(): will push o=%+v to the stack\n", o)
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}
	vm.stack[vm.sp] = o
	vm.sp++
	return nil
}

func (vm *VM) pop() object.Object {
	o := vm.stack[vm.sp-1]
	vm.sp--
	fmt.Printf("vm.pop(): just popped %+v from the stack\n", o)
	return o
}

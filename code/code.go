package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/*
  - Bytecode is made up of instructions.
  - Instructions themselves are a *series of bytes*.
  - A single instruction consists of an Opcode (1 byte) + Operands (>= 1 bytes)
PS: We will not define an `Instruction` type because in Go it is easier to work with []byte.
*/

type Opcode byte
type Instructions []byte

const (
	OpConstant Opcode = iota // iota generates increasing byte values for us.
	OpAdd
)

/*
Now we need to ensure that we can somehow represent the fact that OpConstant has 1 operand.
So we define a "Definition" data structure which can encode this fact.
*/

type Definition struct {
	Name          string
	OperandWidths []int // number of bytes _each operand_ makes up.
}

// PS: we will only go up to uint16 (Max 2 bytes to represent the integer space - keeping instructions small for now)
var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
	OpAdd:      {"OpAdd", []int{}},
}

func Lookup(op byte) (*Definition, error) {
	opcode := Opcode(op)
	def, ok := definitions[opcode]
	if !ok {
		return nil, fmt.Errorf("opcode %d not defined", opcode)
	}
	return def, nil
}

/*
Make makes bytecode instructions.
Make ignores error handling for now.
*/
func Make(op Opcode, operands ...int) []byte {
	fmt.Printf("Make(): looking up op=%v in definitions\n", op)
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	// Now we need to count the length of the entire instruction:
	// which is 1 byte for opcode + however many are defined for its associated operands.
	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}
	fmt.Printf("Make(): instructionLen=%d, will allocate to build instruction\n", instructionLen)
	// So now we have the byte representation of the op + length of instructions.
	// We can allocate a byte slice in memory that represents the instruction code.
	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	// We now need to encode all of the operands.
	// We have made a decision to use Big-Endian order to encode operands.
	offset := 1
	for i, o := range operands {
		// Get the width of every operand.
		width := def.OperandWidths[i]
		switch width {
		case 2:
			fmt.Printf("Make(): Encoding operands in big-endian order\n")
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}
	fmt.Printf("Make(): returning instruction %v\n", instruction)
	return instruction
}

func (ins Instructions) String() string {
	var out bytes.Buffer
	i := 0
	for i < len(ins) {
		def, err := Lookup(ins[i])
		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}
		operands, read := ReadOperands(def, ins[i+1:])
		fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(def, operands))
		i += 1 + read
	}
	return out.String()
}

func (ins Instructions) fmtInstruction(def *Definition, operands []int) string {
	operandCount := len(def.OperandWidths)
	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand len %d does not match defined %d\n",
			len(operands), operandCount)
	}
	switch operandCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	}
	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
}

// ReadOperands is the counterpart of Make()
// Make encodes the operands of a bytecode instruction.
// ReadOperands decodes the instructions
func ReadOperands(def *Definition, ins Instructions) (operands []int, offset int) {
	// Use the definition of the opcode to find operand width
	// and allocate enough space to hold them.
	operands = make([]int, len(def.OperandWidths))
	offset = 0
	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(ins))
		}
		offset += width
	}
	return
}

// Also called in the VM during instruction fetch-decode-execute cycle.
func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}

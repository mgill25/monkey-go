# Virtual Machine Executes Bytecode

Bytecode:
	- Opcodes contained in each instruction are 1 byte in size.

opcode: "operator" part of an instruction. aka op.
example: PUSH
but "PUSH" looks like a 4-char string. But in a proper VM, it'd just be a mnemonic
to an actual 1-byte wide opcode.

operands (aka arguments or parameters) are also contained in bytecode.

they are placed alongside each other.

Instruction[PUSH, 505]
Instruction[PUSH, 205]
Instruction[ADD]

But the Operands are not necessarily 1-byte wide.
For example: an integer operand greater than 255
	- would take more than 1 byte to represent it.

Also, bytecode is a **binary format**. Not necessarily as readable as above.
You can't read it as a text file.
PUSH might stand for number 0, POP might refer to number 23 etc...

The Operands are also encoded.
In case an operand needs multiple bytes to be accurately represented - the endianness plays a big role!
	- Little Endian: LSB comes first
	- Big Endian: LSB comes last
Example:
	Say we decide: PUSH = 1, ADD = 2 and the order is Big Endian
	Bytecode will look like so:
		 (PUSH)        (505)
		00000001 00000001 11111001
This is the job of an assembler.
	- Taking a human-readable representation of bytecode and turning it into binary data.
assembly language is the readable version of Bytecode!
It contains mnemonics and readable operands.
The reverse is done by disassemblers.

Bytecode is domain-specific language for a domain-specific machine.
JVM bytecode has specific instructions like:
	- invokeinterface
	- getstatic
	- new
Ruby's bytecode has: 
	- putself
	- send
	- putobject
Lua's bytecode has:
	- dedicated instructions for accessing and manipulating tables and tuples.

All of these will not be available in the instruction set of a general purpose x86-64 CPU.

This ability to specialise is one of the biggest reasons for building a VM.

## What to build first? compiler or VM?
- both at the same time in this book.

## Hello, bytecode

Goal: Compile and Execute: `1 + 2` expression.
	lexer -> parser -> compiler -> vm
Data Structures:
	String -> Tokens -> AST -> Bytecode -> Objects

We will build a *Stack Machine*.

Why? Because stack machines are easier to understand than register machines.

So, for a stack machine to do `1 + 2`, we just need to do:
	PUSH 1
	PUSH 2
	ADD

## The Idea of `Constants`

- We can just implement `PUSH 1` as an instruction, where `1` is the actual integer we are pushing on to the stack. BUT...
- What if we later want to push **other things** on to the stack? Like String literals etc. Putting them into _bytecode_
is technically possible but it would cause a lot of **bloat** - since bytecode size would increase enormously.

Hence, we basically use _constant expressions_ aka _constants_.
These refer to *expressions whose value doesn't change* and can be **determined at compile time**.

This means we don't need to run a program to know what those expressions evaluate to. "20" will remain 20 - don't need to compile to know that.

### Constant Pool
- Basically a section of the Bytecode where these constants are stored.
- In the Instructions themselves - we just use the Index of the Constants instead of the actual values
- So, `PUSH 0` would mean _push the value stored at `constantPool[0]`_

- When we come across an integer in our code, we will evaluate it and keep track of `*object.Integer` by:
    - storing it in memory
    - assigning it a number
    - refer to this instruction by that number in the bytecode instruction.
    - hand over the constants by putting them in a data structure: `constant pool` and use the number as index to retrieve it.

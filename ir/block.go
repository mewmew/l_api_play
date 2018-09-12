package ir

import "github.com/llir/l/ir/types"

// BasicBlock is an LLVM IR basic block.
type BasicBlock struct {
	// Name of local variable associated with the basic block.
	LocalName string
	// Instructions of the basic block.
	Insts []Instruction
	// Terminator of the basic block.
	Term Terminator
}

// NewBlock returns a new basic block based on the given label name. An empty
// label name indicates an unnamed basic block.
func NewBlock(name string) *BasicBlock {
	return &BasicBlock{LocalName: name}
}

// Type returns the type of the basic block.
func (block *BasicBlock) Type() types.Type {
	return types.Label
}

// Ident returns the identifier associated with the basic block.
func (block *BasicBlock) Ident() string {
	panic("not yet implemented")
}

// Name returns the name of the basic block.
func (block *BasicBlock) Name() string {
	return block.LocalName
}

// SetName sets the name of the basic block.
func (block *BasicBlock) SetName(name string) {
	block.LocalName = name
}

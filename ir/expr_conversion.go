package ir

import (
	"fmt"

	"github.com/llir/l/ir/types"
)

// --- [ Conversion expressions ] ----------------------------------------------

// ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprTrunc is an LLVM IR trunc expression.
type ExprTrunc struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewTruncExpr returns a new trunc expression based on the given source value
// and target type.
func NewTruncExpr(from Constant, to types.Type) *ExprTrunc {
	return &ExprTrunc{From: from, To: to}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprTrunc) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprTrunc) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprTrunc) Ident() string {
	// "trunc" "(" Type Constant "to" Type ")"
	return fmt.Sprintf("trunc (%v to %v)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprTrunc) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprZExt is an LLVM IR zext expression.
type ExprZExt struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewZExtExpr returns a new zext expression based on the given source value and
// target type.
func NewZExtExpr(from Constant, to types.Type) *ExprZExt {
	return &ExprZExt{From: from, To: to}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprZExt) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprZExt) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprZExt) Ident() string {
	// "zext" "(" Type Constant "to" Type ")"
	return fmt.Sprintf("zext (%v to %v)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprZExt) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprSExt is an LLVM IR sext expression.
type ExprSExt struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewSExtExpr returns a new sext expression based on the given source value and
// target type.
func NewSExtExpr(from Constant, to types.Type) *ExprSExt {
	return &ExprSExt{From: from, To: to}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprSExt) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprSExt) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprSExt) Ident() string {
	// "sext" "(" Type Constant "to" Type ")"
	return fmt.Sprintf("sext (%v to %v)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprSExt) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFPTrunc is an LLVM IR fptrunc expression.
type ExprFPTrunc struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewFPTruncExpr returns a new fptrunc expression based on the given source
// value and target type.
func NewFPTruncExpr(from Constant, to types.Type) *ExprFPTrunc {
	return &ExprFPTrunc{From: from, To: to}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFPTrunc) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFPTrunc) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFPTrunc) Ident() string {
	// "fptrunc" "(" Type Constant "to" Type ")"
	return fmt.Sprintf("fptrunc (%v to %v)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFPTrunc) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFPExt is an LLVM IR fpext expression.
type ExprFPExt struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewFPExtExpr returns a new fpext expression based on the given source value
// and target type.
func NewFPExtExpr(from Constant, to types.Type) *ExprFPExt {
	return &ExprFPExt{From: from, To: to}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFPExt) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFPExt) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFPExt) Ident() string {
	// "fpext" "(" Type Constant "to" Type ")"
	return fmt.Sprintf("fpext (%v to %v)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFPExt) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFPToUI is an LLVM IR fptoui expression.
type ExprFPToUI struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewFPToUIExpr returns a new fptoui expression based on the given source value
// and target type.
func NewFPToUIExpr(from Constant, to types.Type) *ExprFPToUI {
	return &ExprFPToUI{From: from, To: to}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFPToUI) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFPToUI) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFPToUI) Ident() string {
	// "fptoui" "(" Type Constant "to" Type ")"
	return fmt.Sprintf("fptoui (%v to %v)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFPToUI) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFPToSI is an LLVM IR fptosi expression.
type ExprFPToSI struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewFPToSIExpr returns a new fptosi expression based on the given source value
// and target type.
func NewFPToSIExpr(from Constant, to types.Type) *ExprFPToSI {
	return &ExprFPToSI{From: from, To: to}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFPToSI) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFPToSI) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFPToSI) Ident() string {
	// "fptosi" "(" Type Constant "to" Type ")"
	return fmt.Sprintf("fptosi (%v to %v)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFPToSI) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprUIToFP is an LLVM IR uitofp expression.
type ExprUIToFP struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewUIToFPExpr returns a new uitofp expression based on the given source value
// and target type.
func NewUIToFPExpr(from Constant, to types.Type) *ExprUIToFP {
	return &ExprUIToFP{From: from, To: to}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprUIToFP) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprUIToFP) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprUIToFP) Ident() string {
	// "uitofp" "(" Type Constant "to" Type ")"
	return fmt.Sprintf("uitofp (%v to %v)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprUIToFP) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprSIToFP is an LLVM IR sitofp expression.
type ExprSIToFP struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewSIToFPExpr returns a new sitofp expression based on the given source value
// and target type.
func NewSIToFPExpr(from Constant, to types.Type) *ExprSIToFP {
	return &ExprSIToFP{From: from, To: to}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprSIToFP) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprSIToFP) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprSIToFP) Ident() string {
	// "sitofp" "(" Type Constant "to" Type ")"
	return fmt.Sprintf("sitofp (%v to %v)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprSIToFP) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprPtrToInt is an LLVM IR ptrtoint expression.
type ExprPtrToInt struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewPtrToIntExpr returns a new ptrtoint expression based on the given source
// value and target type.
func NewPtrToIntExpr(from Constant, to types.Type) *ExprPtrToInt {
	return &ExprPtrToInt{From: from, To: to}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprPtrToInt) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprPtrToInt) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprPtrToInt) Ident() string {
	// "ptrtoint" "(" Type Constant "to" Type ")"
	return fmt.Sprintf("ptrtoint (%v to %v)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprPtrToInt) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprIntToPtr is an LLVM IR inttoptr expression.
type ExprIntToPtr struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewIntToPtrExpr returns a new inttoptr expression based on the given source
// value and target type.
func NewIntToPtrExpr(from Constant, to types.Type) *ExprIntToPtr {
	return &ExprIntToPtr{From: from, To: to}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprIntToPtr) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprIntToPtr) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprIntToPtr) Ident() string {
	// "inttoptr" "(" Type Constant "to" Type ")"
	return fmt.Sprintf("inttoptr (%v to %v)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprIntToPtr) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprBitCast is an LLVM IR bitcast expression.
type ExprBitCast struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewBitCastExpr returns a new bitcast expression based on the given source
// value and target type.
func NewBitCastExpr(from Constant, to types.Type) *ExprBitCast {
	return &ExprBitCast{From: from, To: to}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprBitCast) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprBitCast) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprBitCast) Ident() string {
	// "bitcast" "(" Type Constant "to" Type ")"
	return fmt.Sprintf("bitcast (%v to %v)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprBitCast) Simplify() Constant {
	panic("not yet implemented")
}

// ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprAddrSpaceCast is an LLVM IR addrspacecast expression.
type ExprAddrSpaceCast struct {
	// Value before conversion.
	From Constant
	// Type after conversion.
	To types.Type
}

// NewAddrSpaceCastExpr returns a new addrspacecast expression based on the
// given source value and target type.
func NewAddrSpaceCastExpr(from Constant, to types.Type) *ExprAddrSpaceCast {
	return &ExprAddrSpaceCast{From: from, To: to}
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprAddrSpaceCast) String() string {
	return fmt.Sprintf("%v %v", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprAddrSpaceCast) Type() types.Type {
	return e.To
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprAddrSpaceCast) Ident() string {
	// "addrspacecast" "(" Type Constant "to" Type ")"
	return fmt.Sprintf("addrspacecast (%v to %v)", e.From, e.To)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprAddrSpaceCast) Simplify() Constant {
	panic("not yet implemented")
}

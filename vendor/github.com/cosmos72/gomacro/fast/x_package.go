// this file was generated by gomacro command: import _i "github.com/cosmos72/gomacro/fast"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package fast

import (
	r "reflect"

	"github.com/cosmos72/gomacro/imports"
)

// reflection: allow interpreted code to import "github.com/cosmos72/gomacro/fast"
func init() {
	imports.Packages["github.com/cosmos72/gomacro/fast"] = imports.Package{
		Binds: map[string]r.Value{
			"ConstBind":           r.ValueOf(ConstBind),
			"ConstBindDescriptor": r.ValueOf(ConstBindDescriptor),
			"EFlag4Value":         r.ValueOf(EFlag4Value),
			"EIsNil":              r.ValueOf(EIsNil),
			"EIsTypeAssert":       r.ValueOf(EIsTypeAssert),
			"FuncBind":            r.ValueOf(FuncBind),
			"IntBind":             r.ValueOf(IntBind),
			"MakeEFlag":           r.ValueOf(MakeEFlag),
			"New":                 r.ValueOf(New),
			"NewComp":             r.ValueOf(NewComp),
			"NewEnv":              r.ValueOf(NewEnv),
			"NewInnerInterp":      r.ValueOf(NewInnerInterp),
			"NoIndex":             r.ValueOf(NoIndex),
			"OptDefaults":         r.ValueOf(COptDefaults),
			"OptKeepUntyped":      r.ValueOf(COptKeepUntyped),
			"PlaceAddress":        r.ValueOf(PlaceAddress),
			"PlaceSettable":       r.ValueOf(PlaceSettable),
			"VarBind":             r.ValueOf(VarBind),
		}, Types: map[string]r.Type{
			"Assign":             r.TypeOf((*Assign)(nil)).Elem(),
			"Bind":               r.TypeOf((*Bind)(nil)).Elem(),
			"BindClass":          r.TypeOf((*BindClass)(nil)).Elem(),
			"BindDescriptor":     r.TypeOf((*BindDescriptor)(nil)).Elem(),
			"Builtin":            r.TypeOf((*Builtin)(nil)).Elem(),
			"Call":               r.TypeOf((*Call)(nil)).Elem(),
			"Code":               r.TypeOf((*Code)(nil)).Elem(),
			"Comp":               r.TypeOf((*Comp)(nil)).Elem(),
			"CompGlobals":        r.TypeOf((*CompGlobals)(nil)).Elem(),
			"CompileOptions":     r.TypeOf((*CompileOptions)(nil)).Elem(),
			"EFlags":             r.TypeOf((*EFlags)(nil)).Elem(),
			"Env":                r.TypeOf((*Env)(nil)).Elem(),
			"Expr":               r.TypeOf((*Expr)(nil)).Elem(),
			"FuncInfo":           r.TypeOf((*FuncInfo)(nil)).Elem(),
			"Function":           r.TypeOf((*Function)(nil)).Elem(),
			"I":                  r.TypeOf((*I)(nil)).Elem(),
			"Import":             r.TypeOf((*Import)(nil)).Elem(),
			"Interp":             r.TypeOf((*Interp)(nil)).Elem(),
			"Lit":                r.TypeOf((*Lit)(nil)).Elem(),
			"LoopInfo":           r.TypeOf((*LoopInfo)(nil)).Elem(),
			"Macro":              r.TypeOf((*Macro)(nil)).Elem(),
			"Place":              r.TypeOf((*Place)(nil)).Elem(),
			"PlaceOption":        r.TypeOf((*PlaceOption)(nil)).Elem(),
			"Stmt":               r.TypeOf((*Stmt)(nil)).Elem(),
			"Symbol":             r.TypeOf((*Symbol)(nil)).Elem(),
			"Run":                r.TypeOf((*Run)(nil)).Elem(),
			"TypeAssertionError": r.TypeOf((*TypeAssertionError)(nil)).Elem(),
			"UntypedLit":         r.TypeOf((*UntypedLit)(nil)).Elem(),
			"Var":                r.TypeOf((*Var)(nil)).Elem(),
		}, Untypeds: map[string]string{}, Wrappers: map[string][]string{
			"Bind":        []string{"ConstTo", "DefaultType", "ReflectValue", "Untyped", "UntypedKind"},
			"Comp":        []string{"CollectAst", "CollectNode", "CollectPackageImportsWithRename", "Copy", "Debugf", "Error", "Errorf", "Fprintf", "Gensym", "GensymAnonymous", "GensymPrivate", "IncLine", "IncLineBytes", "LookupPackage", "ParseBytes", "Position", "Print", "ReadMultiline", "Sprintf", "ToString", "TypeOfBool", "TypeOfBuiltin", "TypeOfComplex128", "TypeOfComplex64", "TypeOfError", "TypeOfFloat32", "TypeOfFloat64", "TypeOfFunction", "TypeOfImport", "TypeOfInt", "TypeOfInt16", "TypeOfInt32", "TypeOfInt64", "TypeOfInt8", "TypeOfInterface", "TypeOfMacro", "TypeOfString", "TypeOfUint", "TypeOfUint16", "TypeOfUint32", "TypeOfUint64", "TypeOfUint8", "TypeOfUintptr", "TypeOfUntypedLit", "UnloadPackage", "WarnExtraValues", "Warnf", "WriteDeclsToFile", "WriteDeclsToStream"},
			"CompGlobals": []string{"CollectAst", "CollectNode", "CollectPackageImportsWithRename", "Copy", "Debugf", "Error", "Errorf", "Fprintf", "Gensym", "GensymAnonymous", "GensymPrivate", "ImportPackage", "IncLine", "IncLineBytes", "LookupPackage", "ParseBytes", "Position", "Print", "ReadMultiline", "Sprintf", "ToString", "UnloadPackage", "WarnExtraValues", "Warnf", "WriteDeclsToFile", "WriteDeclsToStream"},
			"Expr":        []string{"IsNil", "ReflectValue", "Untyped", "UntypedKind"},
			"Place":       []string{"Address", "AsPlace", "AsSymbol"},
			"Run":         []string{"CollectAst", "CollectNode", "CollectPackageImportsWithRename", "Copy", "Debugf", "Error", "Errorf", "Fprintf", "Gensym", "GensymAnonymous", "GensymPrivate", "ImportPackage", "IncLine", "IncLineBytes", "LookupPackage", "ParseBytes", "Position", "Print", "ReadMultiline", "Sprintf", "ToString", "UnloadPackage", "WarnExtraValues", "Warnf", "WriteDeclsToFile", "WriteDeclsToStream"},
			"Symbol":      []string{"AsSymbol", "Const", "ConstTo", "ConstValue", "DefaultType", "ReflectValue", "String", "Untyped", "UntypedKind"},
		},
	}
}
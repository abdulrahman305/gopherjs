package compiler

// decls.go contains logic responsible for compiling top-level declarations,
// such as imports, types, functions, etc.

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"sort"
	"strings"

	"github.com/gopherjs/gopherjs/compiler/analysis"
	"github.com/gopherjs/gopherjs/compiler/internal/dce"
	"github.com/gopherjs/gopherjs/compiler/internal/symbol"
	"github.com/gopherjs/gopherjs/compiler/internal/typeparams"
	"github.com/gopherjs/gopherjs/compiler/typesutil"
)

// Decl represents a package-level symbol (e.g. a function, variable or type).
//
// It contains code generated by the compiler for this specific symbol, which is
// grouped by the execution stage it belongs to in the JavaScript runtime.
type Decl struct {
	// The package- or receiver-type-qualified name of function or method obj.
	// See go/types.Func.FullName().
	FullName string
	// A logical equivalent of a symbol name in an object file in the traditional
	// Go compiler/linker toolchain. Used by GopherJS to support go:linkname
	// directives. Must be set for decls that are supported by go:linkname
	// implementation.
	LinkingName symbol.Name
	// A list of package-level JavaScript variable names this symbol needs to declare.
	Vars []string
	// A JS expression by which the object represented by this decl may be
	// referenced within the package context. Empty if the decl represents no such
	// object.
	RefExpr string
	// NamedRecvType is method named recv declare.
	NamedRecvType string
	// JavaScript code that declares basic information about a symbol. For a type
	// it configures basic information about the type and its identity. For a function
	// or method it contains its compiled body.
	DeclCode []byte
	// JavaScript code that initializes reflection metadata about type's method list.
	MethodListCode []byte
	// JavaScript code that initializes the rest of reflection metadata about a type
	// (e.g. struct fields, array type sizes, element types, etc.).
	TypeInitCode []byte
	// JavaScript code that needs to be executed during the package init phase to
	// set the symbol up (e.g. initialize package-level variable value).
	InitCode []byte
	// dce stores the information for dead-code elimination.
	dce dce.Info
	// Set to true if a function performs a blocking operation (I/O or
	// synchronization). The compiler will have to generate function code such
	// that it can be resumed after a blocking operation completes without
	// blocking the main thread in the meantime.
	Blocking bool
}

// minify returns a copy of Decl with unnecessary whitespace removed from the
// JS code.
func (d Decl) minify() Decl {
	d.DeclCode = removeWhitespace(d.DeclCode, true)
	d.MethodListCode = removeWhitespace(d.MethodListCode, true)
	d.TypeInitCode = removeWhitespace(d.TypeInitCode, true)
	d.InitCode = removeWhitespace(d.InitCode, true)
	return d
}

// Dce gets the information for dead-code elimination.
func (d *Decl) Dce() *dce.Info {
	return &d.dce
}

// topLevelObjects extracts package-level variables, functions and named types
// from the package AST.
func (fc *funcContext) topLevelObjects(srcs sources) (vars []*types.Var, functions []*ast.FuncDecl, typeNames typesutil.TypeNames) {
	if !fc.isRoot() {
		panic(bailout(fmt.Errorf("functionContext.discoverObjects() must be only called on the package-level context")))
	}

	for _, file := range srcs.Files {
		for _, decl := range file.Decls {
			switch d := decl.(type) {
			case *ast.FuncDecl:
				sig := fc.pkgCtx.Defs[d.Name].(*types.Func).Type().(*types.Signature)
				if sig.Recv() == nil {
					fc.objectName(fc.pkgCtx.Defs[d.Name]) // register toplevel name
				}
				if !isBlank(d.Name) {
					functions = append(functions, d)
				}
			case *ast.GenDecl:
				switch d.Tok {
				case token.TYPE:
					for _, spec := range d.Specs {
						o := fc.pkgCtx.Defs[spec.(*ast.TypeSpec).Name].(*types.TypeName)
						typeNames.Add(o)
						fc.objectName(o) // register toplevel name
					}
				case token.VAR:
					for _, spec := range d.Specs {
						for _, name := range spec.(*ast.ValueSpec).Names {
							if !isBlank(name) {
								o := fc.pkgCtx.Defs[name].(*types.Var)
								vars = append(vars, o)
								fc.objectName(o) // register toplevel name
							}
						}
					}
				case token.CONST:
					// skip, constants are inlined
				}
			}
		}
	}

	return vars, functions, typeNames
}

// importDecls processes import declarations.
//
// For each imported package:
//   - A new package-level variable is reserved to refer to symbols from that
//     package.
//   - A Decl instance is generated to be included in the Archive.
//
// Lists of imported package paths and corresponding Decls is returned to the caller.
func (fc *funcContext) importDecls() (importedPaths []string, importDecls []*Decl) {
	if !fc.isRoot() {
		panic(bailout(fmt.Errorf("functionContext.importDecls() must be only called on the package-level context")))
	}

	imports := []*types.Package{}
	for _, pkg := range fc.pkgCtx.Pkg.Imports() {
		if pkg == types.Unsafe {
			// Prior to Go 1.9, unsafe import was excluded by Imports() method,
			// but now we do it here to maintain previous behavior.
			continue
		}
		imports = append(imports, pkg)
	}

	// Deterministic processing order.
	sort.Slice(imports, func(i, j int) bool { return imports[i].Path() < imports[j].Path() })

	for _, pkg := range imports {
		importedPaths = append(importedPaths, pkg.Path())
		importDecls = append(importDecls, fc.newImportDecl(pkg))
	}

	return importedPaths, importDecls
}

// newImportDecl registers the imported package and returns a Decl instance for it.
func (fc *funcContext) newImportDecl(importedPkg *types.Package) *Decl {
	pkgVar := fc.importedPkgVar(importedPkg)
	d := &Decl{
		Vars:     []string{pkgVar},
		DeclCode: []byte(fmt.Sprintf("\t%s = $packages[\"%s\"];\n", pkgVar, importedPkg.Path())),
		InitCode: fc.CatchOutput(1, func() { fc.translateStmt(fc.importInitializer(importedPkg.Path()), nil) }),
	}
	d.Dce().SetAsAlive()
	return d
}

// importInitializer calls the imported package $init() function to ensure it is
// initialized before any code in the importer package runs.
func (fc *funcContext) importInitializer(impPath string) ast.Stmt {
	pkgVar := fc.pkgCtx.pkgVars[impPath]
	id := fc.newIdent(fmt.Sprintf(`%s.$init`, pkgVar), types.NewSignatureType(nil, nil, nil, nil, nil, false))
	call := &ast.CallExpr{Fun: id}
	fc.Blocking[call] = true
	fc.Flattened[call] = true

	return &ast.ExprStmt{X: call}
}

// varDecls translates all package-level variables.
//
// `vars` argument must contain all package-level variables found in the package.
// The method returns corresponding Decls that declare and initialize the vars
// as appropriate. Decls are returned in order necessary to correctly initialize
// the variables, considering possible dependencies between them.
func (fc *funcContext) varDecls(vars []*types.Var) []*Decl {
	if !fc.isRoot() {
		panic(bailout(fmt.Errorf("functionContext.varDecls() must be only called on the package-level context")))
	}

	var varDecls []*Decl
	varsWithInit := fc.pkgCtx.VarsWithInitializers()

	initializers := []*types.Initializer{}

	// For implicitly-initialized vars we generate synthetic zero-value
	// initializers and then process them the same way as explicitly initialized.
	for _, o := range vars {
		if varsWithInit[o] {
			continue
		}
		initializer := &types.Initializer{
			Lhs: []*types.Var{o},
			Rhs: fc.zeroValue(o.Type()),
		}
		initializers = append(initializers, initializer)
	}

	// Add explicitly-initialized variables to the list. Implicitly-initialized
	// variables should be declared first in case explicit initializers depend on
	// them.
	initializers = append(initializers, fc.pkgCtx.InitOrder...)

	for _, init := range initializers {
		varDecls = append(varDecls, fc.newVarDecl(init))
	}

	return varDecls
}

// newVarDecl creates a new Decl describing a variable, given an explicit
// initializer.
func (fc *funcContext) newVarDecl(init *types.Initializer) *Decl {
	var d Decl

	assignLHS := []ast.Expr{}
	for _, o := range init.Lhs {
		assignLHS = append(assignLHS, fc.newIdentFor(o))

		// For non-exported package-level variables we need to declared a local JS
		// variable. Exported variables are represented as properties of the $pkg
		// JS object.
		if !o.Exported() {
			d.Vars = append(d.Vars, fc.objectName(o))
		}
		if fc.pkgCtx.HasPointer[o] && !o.Exported() {
			d.Vars = append(d.Vars, fc.varPtrName(o))
		}
	}

	fc.pkgCtx.CollectDCEDeps(&d, func() {
		fc.localVars = nil
		d.InitCode = fc.CatchOutput(1, func() {
			fc.translateStmt(&ast.AssignStmt{
				Lhs: assignLHS,
				Tok: token.DEFINE,
				Rhs: []ast.Expr{init.Rhs},
			}, nil)
		})

		// Initializer code may have introduced auxiliary variables (e.g. for
		// handling multi-assignment or blocking calls), add them to the decl too.
		d.Vars = append(d.Vars, fc.localVars...)
		fc.localVars = nil // Clean up after ourselves.
	})

	d.Dce().SetName(init.Lhs[0])
	if len(init.Lhs) != 1 || analysis.HasSideEffect(init.Rhs, fc.pkgCtx.Info.Info) {
		d.Dce().SetAsAlive()
	}
	return &d
}

// funcDecls translates all package-level function and methods.
//
// `functions` must contain all package-level function and method declarations
// found in the AST. The function returns Decls that define corresponding JS
// functions at runtime. For special functions like init() and main() decls will
// also contain code necessary to invoke them.
func (fc *funcContext) funcDecls(functions []*ast.FuncDecl) ([]*Decl, error) {
	var funcDecls []*Decl
	var mainFunc *types.Func
	for _, fun := range functions {
		o := fc.pkgCtx.Defs[fun.Name].(*types.Func)

		if fun.Recv == nil {
			// Auxiliary decl shared by all instances of the function that defines
			// package-level variable by which they all are referenced.
			varDecl := Decl{}
			varDecl.Dce().SetName(o)
			varDecl.Vars = []string{fc.objectName(o)}
			if o.Type().(*types.Signature).TypeParams().Len() != 0 {
				varDecl.DeclCode = fc.CatchOutput(0, func() {
					fc.Printf("%s = {};", fc.objectName(o))
				})
			}
			funcDecls = append(funcDecls, &varDecl)
		}

		for _, inst := range fc.knownInstances(o) {
			funcDecls = append(funcDecls, fc.newFuncDecl(fun, inst))

			if o.Name() == "main" {
				mainFunc = o // main() function candidate.
			}
		}
	}
	if fc.pkgCtx.isMain() {
		if mainFunc == nil {
			return nil, fmt.Errorf("missing main function")
		}
		// Add a special Decl for invoking main() function after the program has
		// been initialized. It must come after all other functions, especially all
		// init() functions, otherwise main() will be invoked too early.
		funcDecls = append(funcDecls, &Decl{
			InitCode: fc.CatchOutput(1, func() { fc.translateStmt(fc.callMainFunc(mainFunc), nil) }),
		})
	}
	return funcDecls, nil
}

// newFuncDecl returns a Decl that defines a package-level function or a method.
func (fc *funcContext) newFuncDecl(fun *ast.FuncDecl, inst typeparams.Instance) *Decl {
	o := fc.pkgCtx.Defs[fun.Name].(*types.Func)
	d := &Decl{
		FullName:    o.FullName(),
		Blocking:    fc.pkgCtx.IsBlocking(o),
		LinkingName: symbol.New(o),
	}
	d.Dce().SetName(o, inst.TArgs...)

	if typesutil.IsMethod(o) {
		recv := typesutil.RecvType(o.Type().(*types.Signature)).Obj()
		d.NamedRecvType = fc.objectName(recv)
	} else {
		d.RefExpr = fc.instName(inst)
		switch o.Name() {
		case "main":
			if fc.pkgCtx.isMain() { // Found main() function of the program.
				d.Dce().SetAsAlive() // Always reachable.
			}
		case "init":
			d.InitCode = fc.CatchOutput(1, func() { fc.translateStmt(fc.callInitFunc(o), nil) })
			d.Dce().SetAsAlive() // init() function is always reachable.
		}
	}

	fc.pkgCtx.CollectDCEDeps(d, func() {
		d.DeclCode = fc.namedFuncContext(inst).translateTopLevelFunction(fun)
	})
	return d
}

// callInitFunc returns an AST statement for calling the given instance of the
// package's init() function.
func (fc *funcContext) callInitFunc(init *types.Func) ast.Stmt {
	id := fc.newIdentFor(init)
	call := &ast.CallExpr{Fun: id}
	if fc.pkgCtx.IsBlocking(init) {
		fc.Blocking[call] = true
	}
	return &ast.ExprStmt{X: call}
}

// callMainFunc returns an AST statement for calling the main() function of the
// program, which should be included in the $init() function of the main package.
func (fc *funcContext) callMainFunc(main *types.Func) ast.Stmt {
	id := fc.newIdentFor(main)
	call := &ast.CallExpr{Fun: id}
	ifStmt := &ast.IfStmt{
		Cond: fc.newIdent("$pkg === $mainPkg", types.Typ[types.Bool]),
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{X: call},
				&ast.AssignStmt{
					Lhs: []ast.Expr{fc.newIdent("$mainFinished", types.Typ[types.Bool])},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{fc.newConst(types.Typ[types.Bool], constant.MakeBool(true))},
				},
			},
		},
	}
	if fc.pkgCtx.IsBlocking(main) {
		fc.Blocking[call] = true
		fc.Flattened[ifStmt] = true
	}

	return ifStmt
}

// namedTypeDecls returns Decls that define all names Go types.
//
// `typeNames` must contain all named types defined in the package, including
// those defined inside function bodies.
func (fc *funcContext) namedTypeDecls(typeNames typesutil.TypeNames) ([]*Decl, error) {
	if !fc.isRoot() {
		panic(bailout(fmt.Errorf("functionContext.namedTypeDecls() must be only called on the package-level context")))
	}

	var typeDecls []*Decl
	for _, o := range typeNames.Slice() {
		if o.IsAlias() {
			continue
		}

		typeDecls = append(typeDecls, fc.newNamedTypeVarDecl(o))

		for _, inst := range fc.knownInstances(o) {
			d, err := fc.newNamedTypeInstDecl(inst)
			if err != nil {
				return nil, err
			}
			typeDecls = append(typeDecls, d)
		}
	}

	return typeDecls, nil
}

// newNamedTypeVarDecl returns a Decl that defines a JS variable to store named
// type definition.
//
// For generic types, the variable is an object containing known instantiations
// of the type, keyed by the type argument combination. Otherwise it contains
// the type definition directly.
func (fc *funcContext) newNamedTypeVarDecl(obj *types.TypeName) *Decl {
	varDecl := &Decl{Vars: []string{fc.objectName(obj)}}
	if typeparams.HasTypeParams(obj.Type()) {
		varDecl.DeclCode = fc.CatchOutput(0, func() {
			fc.Printf("%s = {};", fc.objectName(obj))
		})
	}
	if isPkgLevel(obj) {
		varDecl.TypeInitCode = fc.CatchOutput(0, func() {
			fc.Printf("$pkg.%s = %s;", encodeIdent(obj.Name()), fc.objectName(obj))
		})
	}
	return varDecl
}

// newNamedTypeInstDecl returns a Decl that represents an instantiation of a
// named Go type.
func (fc *funcContext) newNamedTypeInstDecl(inst typeparams.Instance) (*Decl, error) {
	originType := inst.Object.Type().(*types.Named)

	fc.typeResolver = typeparams.NewResolver(fc.pkgCtx.typesCtx, typeparams.ToSlice(originType.TypeParams()), inst.TArgs)
	defer func() { fc.typeResolver = nil }()

	instanceType := originType
	if !inst.IsTrivial() {
		instantiated, err := types.Instantiate(fc.pkgCtx.typesCtx, originType, inst.TArgs, true)
		if err != nil {
			return nil, fmt.Errorf("failed to instantiate type %v with args %v: %w", originType, inst.TArgs, err)
		}
		instanceType = instantiated.(*types.Named)
	}

	underlying := instanceType.Underlying()
	d := &Decl{}
	d.Dce().SetName(inst.Object, inst.TArgs...)
	fc.pkgCtx.CollectDCEDeps(d, func() {
		// Code that declares a JS type (i.e. prototype) for each Go type.
		d.DeclCode = fc.CatchOutput(0, func() {
			size := int64(0)
			constructor := "null"

			switch t := underlying.(type) {
			case *types.Struct:
				constructor = fc.structConstructor(t)
			case *types.Basic, *types.Array, *types.Slice, *types.Chan, *types.Signature, *types.Interface, *types.Pointer, *types.Map:
				size = sizes32.Sizeof(t)
			}
			if tPointer, ok := underlying.(*types.Pointer); ok {
				if _, ok := tPointer.Elem().Underlying().(*types.Array); ok {
					// Array pointers have non-default constructors to support wrapping
					// of the native objects.
					constructor = "$arrayPtrCtor()"
				}
			}
			fc.Printf(`%s = $newType(%d, %s, %q, %t, "%s", %t, %s);`,
				fc.instName(inst), size, typeKind(originType), inst.TypeString(), inst.Object.Name() != "", inst.Object.Pkg().Path(), inst.Object.Exported(), constructor)
		})

		// Reflection metadata about methods the type has.
		d.MethodListCode = fc.CatchOutput(0, func() {
			if _, ok := underlying.(*types.Interface); ok {
				return
			}
			var methods []string
			var ptrMethods []string
			for i := 0; i < instanceType.NumMethods(); i++ {
				entry, isPtr := fc.methodListEntry(instanceType.Method(i))
				if isPtr {
					ptrMethods = append(ptrMethods, entry)
				} else {
					methods = append(methods, entry)
				}
			}
			if len(methods) > 0 {
				fc.Printf("%s.methods = [%s];", fc.instName(inst), strings.Join(methods, ", "))
			}
			if len(ptrMethods) > 0 {
				fc.Printf("%s.methods = [%s];", fc.typeName(types.NewPointer(instanceType)), strings.Join(ptrMethods, ", "))
			}
		})

		// Certain types need to run additional type-specific logic to fully
		// initialize themselves.
		switch t := underlying.(type) {
		case *types.Array, *types.Chan, *types.Interface, *types.Map, *types.Pointer, *types.Slice, *types.Signature, *types.Struct:
			d.TypeInitCode = fc.CatchOutput(0, func() {
				fc.Printf("%s.init(%s);", fc.instName(inst), fc.initArgs(t))
			})
		}
	})
	return d, nil
}

// structConstructor returns JS constructor function for a struct type.
func (fc *funcContext) structConstructor(t *types.Struct) string {
	constructor := &strings.Builder{}

	ctrArgs := make([]string, t.NumFields())
	for i := 0; i < t.NumFields(); i++ {
		ctrArgs[i] = fieldName(t, i) + "_"
	}

	fmt.Fprintf(constructor, "function(%s) {\n", strings.Join(ctrArgs, ", "))
	fmt.Fprintf(constructor, "\t\tthis.$val = this;\n")

	// If no arguments were passed, zero-initialize all fields.
	fmt.Fprintf(constructor, "\t\tif (arguments.length === 0) {\n")
	for i := 0; i < t.NumFields(); i++ {
		fmt.Fprintf(constructor, "\t\t\tthis.%s = %s;\n", fieldName(t, i), fc.translateExpr(fc.zeroValue(t.Field(i).Type())).String())
	}
	fmt.Fprintf(constructor, "\t\t\treturn;\n")
	fmt.Fprintf(constructor, "\t\t}\n")

	// Otherwise initialize fields with the provided values.
	for i := 0; i < t.NumFields(); i++ {
		fmt.Fprintf(constructor, "\t\tthis.%[1]s = %[1]s_;\n", fieldName(t, i))
	}
	fmt.Fprintf(constructor, "\t}")
	return constructor.String()
}

// methodListEntry returns a JS code fragment that describes the given method
// function for runtime reflection. It returns isPtr=true if the method belongs
// to the pointer-receiver method list.
func (fc *funcContext) methodListEntry(method *types.Func) (entry string, isPtr bool) {
	name := method.Name()
	if reservedKeywords[name] {
		name += "$"
	}
	pkgPath := ""
	if !method.Exported() {
		pkgPath = method.Pkg().Path()
	}
	t := method.Type().(*types.Signature)
	entry = fmt.Sprintf(`{prop: "%s", name: %s, pkg: "%s", typ: $funcType(%s)}`,
		name, encodeString(method.Name()), pkgPath, fc.initArgs(t))
	_, isPtr = t.Recv().Type().(*types.Pointer)
	return entry, isPtr
}

// anonTypeDecls returns a list of Decls corresponding to anonymous Go types
// encountered in the package.
//
// `anonTypes` must contain an ordered list of anonymous types with the
// identifiers that were auto-assigned to them. They must be sorted in the
// topological initialization order (e.g. `[]int` is before `struct{f []int}`).
//
// See also typesutil.AnonymousTypes.
func (fc *funcContext) anonTypeDecls(anonTypes []*types.TypeName) []*Decl {
	if !fc.isRoot() {
		panic(bailout(fmt.Errorf("functionContext.anonTypeDecls() must be only called on the package-level context")))
	}
	decls := []*Decl{}
	for _, t := range anonTypes {
		d := &Decl{
			Vars: []string{t.Name()},
		}
		d.Dce().SetName(t)
		fc.pkgCtx.CollectDCEDeps(d, func() {
			d.DeclCode = []byte(fmt.Sprintf("\t%s = $%sType(%s);\n", t.Name(), strings.ToLower(typeKind(t.Type())[5:]), fc.initArgs(t.Type())))
		})
		decls = append(decls, d)
	}
	return decls
}

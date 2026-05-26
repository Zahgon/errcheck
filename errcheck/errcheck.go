// Package errcheck is the library used to implement the errcheck command-line tool.
package errcheck

import (
	"errors"
	"go/ast"
	"go/token"
	"go/types"
	"regexp"

	"golang.org/x/tools/go/packages"
)

var errorType *types.Interface

func init() {
	errorType = types.Universe.Lookup("error").Type().Underlying().(*types.Interface)
}

var (
	// ErrNoGoFiles is returned when CheckPackage is run on a package with no Go source files.
	//
	// Deprecated: this error is no longer returned by errcheck.LoadPackages.
	ErrNoGoFiles = errors.New("package contains no go source files")
)

// UncheckedError indicates the position of an unchecked error return.
type UncheckedError struct {
	Pos          token.Position
	Line         string
	FuncName     string
	SelectorName string
}

// Result is returned from the CheckPackage function, and holds all the errors
// that were found to be unchecked in a package.
//
// Aggregation can be done using the Append method for users that want to
// combine results from multiple packages.
type Result struct {
	// UncheckedErrors is a list of all the unchecked errors in the package.
	// Printing an error reports its position within the file and the contents of the line.
	UncheckedErrors []UncheckedError
}

type byName []UncheckedError

// Less reports whether the element with index i should sort before the element with index j.
func (b byName) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (b byName) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (b byName) Len() int {
	_ = "STUB: not implemented"

	// Append appends errors to e. Append does not do any duplicate checking.
	return 0
}

func (r *Result) Append(other Result) { _ = "STUB: not implemented"; return }

// Unique returns the unique errors that have been accumulated. Duplicates may occur
// when a file containing an unchecked error belongs to > 1 package.
//
// The method receiver remains unmodified after the call to Unique.
func (r Result) Unique() Result { _ = "STUB: not implemented"; return *new(Result) }

// compact in-place

// Exclusions define symbols and language elements that will be not checked
type Exclusions struct {

	// Packages lists paths of excluded packages.
	Packages []string

	// SymbolRegexpsByPackage maps individual package paths to regular
	// expressions that match symbols to be excluded.
	//
	// Packages whose paths appear both here and in Packages list will
	// be excluded entirely.
	//
	// This is a legacy input that will be deprecated in errcheck version 2 and
	// should not be used.
	SymbolRegexpsByPackage map[string]*regexp.Regexp

	// Symbols lists patterns that exclude individual package symbols.
	//
	// For example:
	//
	//   "fmt.Errorf"              // function
	//   "fmt.Fprintf(os.Stderr)"  // function with set argument value
	//   "(hash.Hash).Write"       // method
	//
	Symbols []string

	// TestFiles excludes _test.go files.
	TestFiles bool

	// GeneratedFiles excludes generated source files.
	//
	// Source file is assumed to be generated if its contents
	// match the following regular expression:
	//
	//   ^// Code generated .* DO NOT EDIT\\.$
	//
	GeneratedFiles bool

	// BlankAssignments ignores assignments to blank identifier.
	BlankAssignments bool

	// TypeAssertions ignores unchecked type assertions.
	TypeAssertions bool
}

// Checker checks that you checked errors.
type Checker struct {
	// Exclusions defines code packages, symbols, and other elements that will not be checked.
	Exclusions Exclusions

	// Tags are a list of build tags to use.
	Tags []string

	// The mod flag for go build.
	Mod string
}

// loadPackages is used for testing.
var loadPackages = func(cfg *packages.Config, paths ...string) ([]*packages.Package, error) {
	return packages.Load(cfg, paths...)
}

// LoadPackages loads all the packages in all the paths provided. It uses the
// exclusions and build tags provided to by the user when loading the packages.
func (c *Checker) LoadPackages(paths ...string) ([]*packages.Package, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var generatedCodeRegexp = regexp.MustCompile("^// Code generated .* DO NOT EDIT\\.$")
var dotStar = regexp.MustCompile(".*")

func (c *Checker) shouldSkipFile(file *ast.File) bool { _ = "STUB: not implemented"; return false }

// CheckPackage checks packages for errors that have not been checked.
//
// It will exclude specific errors from analysis if the user has configured
// exclusions.
func (c *Checker) CheckPackage(pkg *packages.Package) Result {
	_ = "STUB: not implemented"
	return *new(Result)
}

// Apply SymbolRegexpsByPackage first so that if the same path appears in
// Packages, a more narrow regexp will be superseded by dotStar below.

// TODO warn if previous entry overwritten?

// TODO warn if previous entry overwritten?

// visitor implements the errcheck algorithm
type visitor struct {
	typesInfo *types.Info
	fset      *token.FileSet
	ignore    map[string]*regexp.Regexp
	blank     bool
	asserts   bool
	lines     map[string][]string
	exclude   map[string]bool

	errors []UncheckedError
}

// selectorAndFunc tries to get the selector and function from call expression.
// For example, given the call expression representing "a.b()", the selector
// is "a.b" and the function is "b" itself.
//
// The final return value will be true if it is able to do extract a selector
// from the call and look up the function object it refers to.
//
// If the call does not include a selector (like if it is a plain "f()" function call)
// then the final return value will be false.
func (v *visitor) selectorAndFunc(call *ast.CallExpr) (*ast.SelectorExpr, *types.Func, bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

// Shouldn't happen, but be paranoid

// fullName will return a package / receiver-type qualified name for a called function
// if the function is the result of a selector. Otherwise it will return
// the empty string.
//
// The name is fully qualified by the import path, possible type,
// function/method name and pointer receiver.
//
// For example,
//   - for "fmt.Printf(...)" it will return "fmt.Printf"
//   - for "base64.StdEncoding.Decode(...)" it will return "(*encoding/base64.Encoding).Decode"
//   - for "myFunc()" it will return ""
func (v *visitor) fullName(call *ast.CallExpr) string { _ = "STUB: not implemented"; return "" }

// TODO(dh): vendored packages will have /vendor/ in their name,
// thus not matching vendored standard library packages. If we
// want to support vendored stdlib packages, we need to implement
// FullName with our own logic.

func getSelectorName(sel *ast.SelectorExpr) string { _ = "STUB: not implemented"; return "" }

// selectorName will return a name for a called function
// if the function is the result of a selector. Otherwise it will return
// the empty string.
//
// The name is fully qualified by the import path, possible type,
// function/method name and pointer receiver.
//
// For example,
//   - for "fmt.Printf(...)" it will return "fmt.Printf"
//   - for "base64.StdEncoding.Decode(...)" it will return "base64.StdEncoding.Decode"
//   - for "myFunc()" it will return ""
func (v *visitor) selectorName(call *ast.CallExpr) string { _ = "STUB: not implemented"; return "" }

// namesForExcludeCheck will return a list of fully-qualified function names
// from a function call that can be used to check against the exclusion list.
//
// If a function call is against a local function (like "myFunc()") then no
// names are returned. If the function is package-qualified (like "fmt.Printf()")
// then just that function's fullName is returned.
//
// Otherwise, we walk through all the potentially embedded interfaces of the receiver
// to collect a list of type-qualified function names that we will check.
func (v *visitor) namesForExcludeCheck(call *ast.CallExpr) []string {
	_ = "STUB: not implemented"
	return nil
}

// This will be missing for functions without a receiver (like fmt.Printf),
// so just fall back to the function's fullName in that case.

// This will return with ok false if the function isn't defined
// on an interface, so just fall back to the fullName.

// Like in fullName, vendored packages will have /vendor/ in their name,
// thus not matching vendored standard library packages. If we
// want to support vendored stdlib packages, we need to implement
// additional logic here.

// isBufferType checks if the expression type is a known in-memory buffer type.
func (v *visitor) argName(expr ast.Expr) string {
	_ = "STUB: not implemented"
	// Special-case literal "os.Stdout" and "os.Stderr"
	return ""
}

func (v *visitor) excludeCall(call *ast.CallExpr) bool { _ = "STUB: not implemented"; return false }

func (v *visitor) ignoreCall(call *ast.CallExpr) bool { _ = "STUB: not implemented"; return false }

// Try to get an identifier.
// Currently only supports simple expressions:
//     1. f()
//     2. x.y.f()
//     3. x.y[T]() or x.y[T1, T2]()

// eg: *ast.SliceExpr

// If we got an identifier for the function, see if it is ignored

// baseCallExpr returns the underlying function expression for a call. Type
// arguments and parentheses wrap the selector/name in additional AST nodes, so
// matching call.Fun directly would miss forms like errors.AsType[T](err).
func baseCallExpr(fun ast.Expr) ast.Expr { _ = "STUB: not implemented"; return *new(ast.Expr) }

// nonVendoredPkgPath returns the unvendored version of the provided package
// path (or returns the provided path if it does not represent a vendored
// path).
func nonVendoredPkgPath(pkgPath string) string { _ = "STUB: not implemented"; return "" }

// errorsByArg returns a slice s such that
// len(s) == number of return types of call
// s[i] == true iff return type at position i from left is an error type
func (v *visitor) errorsByArg(call *ast.CallExpr) []bool { _ = "STUB: not implemented"; return nil }

// Single return

// Single return via pointer

// Multiple returns

// Single return

// Single return via pointer

func (v *visitor) callReturnsError(call *ast.CallExpr) bool {
	_ = "STUB: not implemented"
	return false
}

// isRecover returns true if the given CallExpr is a call to the built-in recover() function.
func (v *visitor) isRecover(call *ast.CallExpr) bool { _ = "STUB: not implemented"; return false }

// TODO (dtcaciuc) collect token.Pos and then convert them to UncheckedErrors
// after visitor is done running. This will allow to integrate more cleanly
// with analyzer so that we don't have to convert Position back to Pos.
func (v *visitor) addErrorAtPosition(position token.Pos, call *ast.CallExpr) {
	_ = "STUB: not implemented"
	return
}

func readfile(filename string) []string { _ = "STUB: not implemented"; return nil }

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	_ = "STUB: not implemented"
	return *new(ast.Visitor)
}

// ignore declarations w/o assignments

// checkAssignment checks the assignment statement and returns a boolean value
// indicating whether to continue checking the substructure in AssignStmt or not
func (v *visitor) checkAssignment(lhs, rhs []ast.Expr) (followed bool) {
	_ = "STUB: not implemented"

	// single value on rhs; check against lhs identifiers
	return false
}

// We shortcut calls to recover() because errorsByArg can't
// check its return types for errors since it returns interface{}.

// type switch

// assertion result not read

// assertion result ignored

// multiple value on rhs; in this case a call can't return
// multiple values. Assume len(lhs) == len(rhs)

// Shouldn't happen anyway, no multi assignment in type switches

func (v *visitor) checkAssertExpr(expr *ast.TypeAssertExpr) { _ = "STUB: not implemented"; return }

// type switch

func isErrorType(t types.Type) bool { _ = "STUB: not implemented"; return false }

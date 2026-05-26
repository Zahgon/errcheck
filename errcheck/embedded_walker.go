package errcheck

import (
	"go/types"
)

// walkThroughEmbeddedInterfaces returns a slice of Interfaces that
// we need to walk through in order to reach the actual definition,
// in an Interface, of the method selected by the given selection.
//
// false will be returned in the second return value if:
//   - the right side of the selection is not a function
//   - the actual definition of the function is not in an Interface
//
// The returned slice will contain all the interface types that need
// to be walked through to reach the actual definition.
//
// For example, say we have:
//
//	type Inner interface {Method()}
//	type Middle interface {Inner}
//	type Outer interface {Middle}
//	type T struct {Outer}
//	type U struct {T}
//	type V struct {U}
//
// And then the selector:
//
//	V.Method
//
// We'll return [Outer, Middle, Inner] by first walking through the embedded structs
// until we reach the Outer interface, then descending through the embedded interfaces
// until we find the one that actually explicitly defines Method.
func walkThroughEmbeddedInterfaces(sel *types.Selection) ([]types.Type, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Start off at the receiver.

// First, we can walk through any Struct fields provided
// by the selection Index() method. We ignore the last
// index because it would give the method itself.

// Now currentT is either a type implementing the actual function,
// an Invalid type (if the receiver is a package), or an interface.
//
// If it's not an Interface, then we're done, as this function
// only cares about Interface-defined functions.
//
// If it is an Interface, we potentially need to continue digging until
// we find the Interface that actually explicitly defines the function.

// The first interface we pass through is this one we've found. We return the possibly
// wrapping types.Named because it is more useful to work with for callers.

// If this interface itself explicitly defines the given method
// then we're done digging.

// Otherwise, we find which of the embedded interfaces _does_
// define the method, add it to our list, and loop.

// Returned a nil interface, we are done.

func getTypeAtFieldIndex(startingAt types.Type, fieldIndex int) types.Type {
	_ = "STUB: not implemented"
	return *new(types.Type)
}

// getEmbeddedInterfaceDefiningMethod searches through any embedded interfaces of the
// passed interface searching for one that defines the given function. If found, the
// types.Named wrapping that interface will be returned along with true in the second value.
//
// If no such embedded interface is found, nil and false are returned.
func getEmbeddedInterfaceDefiningMethod(interfaceT *types.Interface, fn *types.Func) (*types.Named, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func explicitlyDefinesMethod(interfaceT *types.Interface, fn *types.Func) bool {
	_ = "STUB: not implemented"
	return false
}

func definesMethod(interfaceT *types.Interface, fn *types.Func) bool {
	_ = "STUB: not implemented"
	return false
}

func maybeDereference(t types.Type) types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

func maybeUnname(t types.Type) types.Type { _ = "STUB: not implemented"; return *new(types.Type) }

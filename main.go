package main

import (
	"os"
	"regexp"

	"github.com/kisielk/errcheck/errcheck"
)

const (
	exitCodeOk int = iota
	exitUncheckedError
	exitFatalError
)

type ignoreFlag map[string]*regexp.Regexp

// global flags
var (
	abspath bool
	verbose bool
)

func (f ignoreFlag) String() string { _ = "STUB: not implemented"; return "" }

func (f ignoreFlag) Set(s string) error { _ = "STUB: not implemented"; return nil }

type tagsFlag []string

func (f *tagsFlag) String() string { _ = "STUB: not implemented"; return "" }

func (f *tagsFlag) Set(s string) error { _ = "STUB: not implemented"; return nil }

func reportResult(e errcheck.Result) { _ = "STUB: not implemented"; return }

func logf(msg string, args ...interface{}) { _ = "STUB: not implemented"; return }

func mainCmd(args []string) int { _ = "STUB: not implemented"; return 0 }

func checkPaths(c *errcheck.Checker, paths ...string) (errcheck.Result, error) {
	_ = "STUB: not implemented"
	return *new(errcheck.Result), nil
}

// Check for errors in the initial packages.

func parseFlags(checker *errcheck.Checker, args []string) ([]string, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func main() {
	os.Exit(mainCmd(os.Args))
}

package errcheck

import (
	"reflect"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:       "errcheck",
	Doc:        "check for unchecked errors",
	Run:        runAnalyzer,
	ResultType: reflect.TypeOf(Result{}),
}

var (
	argBlank       bool
	argAsserts     bool
	argExcludeFile string
	argExcludeOnly bool
)

func init() {
	Analyzer.Flags.BoolVar(&argBlank, "blank", false, "if true, check for errors assigned to blank identifier")
	Analyzer.Flags.BoolVar(&argAsserts, "assert", false, "if true, check for ignored type assertion results")
	Analyzer.Flags.StringVar(&argExcludeFile, "exclude", "", "Path to a file containing a list of functions to exclude from checking")
	Analyzer.Flags.BoolVar(&argExcludeOnly, "excludeonly", false, "Use only excludes from exclude file")
}

func runAnalyzer(pass *analysis.Pass) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// deprecated & not used

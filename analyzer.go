package mockerylint

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

// ...
func New() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "mockerylint",
		Doc:      "TODO",
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Run:      run,
	}
}

func run(pass *analysis.Pass) (any, error) {
	_ = pass
	return nil, nil
}

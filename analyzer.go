package mockerylint

import (
	"errors"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
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
	visitor, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, errors.New("missing inspect.Analyzer requirement")
	}

	filter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	visitor.Preorder(filter, func(n ast.Node) {
		visit(pass, n)
	})

	return nil, nil
}

func visit(pass *analysis.Pass, node ast.Node) {
	call, ok := node.(*ast.CallExpr)
	if !ok {
		return
	}

	selector, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}

	if selector.Sel.Name != "On" {
		return
	}

	if !isMock(selector.X, pass.TypesInfo) {
		return
	}

	pass.Reportf(node.Pos(), "use .EXPECT instead of .On")
}

func isMock(expr ast.Expr, info *types.Info) bool {
	ident, ok := expr.(*ast.Ident)
	if !ok {
		return false
	}

	ptr, ok := info.ObjectOf(ident).Type().(*types.Pointer)
	if !ok {
		return false
	}

	typ, ok := ptr.Elem().Underlying().(*types.Struct)
	if !ok {
		return false
	}

	for field := range typ.Fields() {
		if !field.Embedded() {
			continue
		}

		if field.Type().String() == "github.com/stretchr/testify/mock.Mock" {
			return true
		}
	}

	return false
}

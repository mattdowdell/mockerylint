package mockerylint_test

import (
	"go/ast"
	"go/token"
	"slices"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/mattdowdell/mockerylint"
)

var rules = map[string]bool{
	"usefactory":  true,
	"useexpecter": true,
	"usetimes":    true,
}

var methods = map[string][]string{
	"use .EXPECT instead of .On":                                        {"On"},
	".Test() can be removed when using mock factory":                    {"Test"},
	".AssertExpectations() can be removed when using mock factory":      {"AssertExpectations"},
	"expectation should call .Maybe(), .Once(), .Twice(), or .Times(N)": {"Return", "RunAndReturn"},
}

func TestAnalyzer(t *testing.T) {
	tests := map[string]struct {
		dir string
	}{
		"useexpecter": {
			dir: "./useexpecter",
		},
		"usefactory": {
			dir: "./usefactory",
		},
		"usetimes": {
			dir: "./usetimes",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			analyzer := mockerylint.New()
			testdata := analysistest.TestData()

			results := analysistest.Run(t, testdata, analyzer, tt.dir)

			checkResults(t, results)
		})
	}
}

// checkResults asserts the properties of a diagnostic that analysistest cannot see. Its
// want comments are matched on the line and the message alone, leaving the category, the
// position within the line, and the suppression of generated files unchecked.
func checkResults(t *testing.T, results []*analysistest.Result) {
	t.Helper()

	var generated int

	for _, result := range results {
		pkg := result.Action.Package
		suppressed := generatedFiles(pkg.Fset, pkg.Syntax)
		generated += len(suppressed)

		for _, diag := range result.Action.Diagnostics {
			pos := pkg.Fset.Position(diag.Pos)

			if !rules[diag.Category] {
				t.Errorf("%s: %q has category %q, want a rule name", pos, diag.Message, diag.Category)
			}

			if suppressed[pos.Filename] {
				t.Errorf("%s: %q reported against generated code", pos, diag.Message)
			}

			checkPosition(t, pos, pkg.Syntax, &diag)
		}
	}

	// Suppression is only being exercised while there is something to suppress, and the
	// check above would pass unnoticed if the generated files lost their header.
	if generated == 0 {
		t.Error("no generated files found, so nothing exercises suppressing them")
	}
}

// checkPosition asserts that a diagnostic about a method call points at the method
// rather than at the receiver its chain starts from.
func checkPosition(t *testing.T, pos token.Position, files []*ast.File, diag *analysis.Diagnostic) {
	t.Helper()

	want, ok := methods[diag.Message]
	if !ok {
		return
	}

	got := identAt(files, diag.Pos)

	if !slices.Contains(want, got) {
		t.Errorf("%s: %q points at %q, want one of %q", pos, diag.Message, got, want)
	}
}

// generatedFiles returns the names of the files carrying a generated code header, which
// the analyzer is expected to skip.
func generatedFiles(fset *token.FileSet, files []*ast.File) map[string]bool {
	names := make(map[string]bool)

	for _, file := range files {
		if ast.IsGenerated(file) {
			names[fset.Position(file.FileStart).Filename] = true
		}
	}

	return names
}

// identAt returns the name of the identifier starting at pos, or an empty string when
// the position does not start one.
func identAt(files []*ast.File, pos token.Pos) string {
	for _, file := range files {
		if pos < file.FileStart || pos >= file.FileEnd {
			continue
		}

		var name string

		ast.Inspect(file, func(node ast.Node) bool {
			if ident, ok := node.(*ast.Ident); ok && ident.Pos() == pos {
				name = ident.Name
			}

			return name == ""
		})

		return name
	}

	return ""
}

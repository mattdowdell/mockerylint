package mockerylint_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/mattdowdell/mockerylint"
)

func TestAnalyzer(t *testing.T) {
	tests := map[string]struct {
		dir string
	}{}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			analyzer := mockerylint.New()
			testdata := analysistest.TestData()

			analysistest.RunWithSuggestedFixes(t, testdata, analyzer, tt.dir)
		})
	}
}

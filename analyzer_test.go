package mockerylint_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/mattdowdell/mockerylint"
)

func TestAnalyzer(t *testing.T) {
	tests := map[string]struct {
		dir   string
		fixes bool
	}{
		"useexpecter": {
			dir:   "./useexpecter",
			fixes: true,
		},
		"usefactory": {
			dir: "./usefactory",
		},
		"usetimes": {
			dir: "./usetimes",
		},
		"noanything": {
			dir: "./noanything",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			analyzer := mockerylint.New(nil)
			testdata := analysistest.TestData()

			if tt.fixes {
				analysistest.RunWithSuggestedFixes(t, testdata, analyzer, tt.dir)
			} else {
				analysistest.Run(t, testdata, analyzer, tt.dir)
			}
		})
	}
}

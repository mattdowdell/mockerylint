package mockerylint_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/mattdowdell/mockerylint"
)

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

			analysistest.Run(t, testdata, analyzer, tt.dir)
		})
	}
}

// The rules the analyzer applies, which are the names of the flags that turn them off
// and the categories carried by the diagnostics each one produces.
const (
	useFactory  = "usefactory"
	useExpecter = "useexpecter"
	useTimes    = "usetimes"
)

var (
	rules = []string{useFactory, useExpecter, useTimes}
	dirs  = []string{"./useexpecter", "./usefactory", "./usetimes"}
)

func TestAnalyzerRuleFlags(t *testing.T) {
	analyzer := mockerylint.New()

	for _, rule := range rules {
		t.Run(rule, func(t *testing.T) {
			found := analyzer.Flags.Lookup(rule)
			if found == nil {
				t.Fatalf("no %s flag registered", rule)
			}

			// A rule that did not default to true would be left out of the output of an
			// analyzer that was never configured.
			if found.DefValue != "true" {
				t.Errorf("%s defaults to %s, want true", rule, found.DefValue)
			}
		})
	}
}

func TestAnalyzerDisable(t *testing.T) {
	for _, rule := range rules {
		t.Run(rule, func(t *testing.T) {
			counts := countDisabled(t, rule)

			if counts[rule] != 0 {
				t.Errorf("disabled %s reported %d diagnostics, want 0", rule, counts[rule])
			}

			// An analyzer reporting nothing at all would satisfy the check above, so the
			// rules left enabled are required to still report.
			for _, other := range rules {
				if other != rule && counts[other] == 0 {
					t.Errorf("%s reported nothing while %s was disabled", other, rule)
				}
			}
		})
	}
}

// countDisabled runs the analyzer over all the testdata with one rule turned off, and
// returns how many diagnostics each rule produced.
func countDisabled(t *testing.T, rule string) map[string]int {
	t.Helper()

	analyzer := mockerylint.New()

	if err := analyzer.Flags.Set(rule, "false"); err != nil {
		t.Fatalf("disabling %s: %v", rule, err)
	}

	// The want comments describe the full set of rules, so a run with one turned off
	// cannot match them. The diagnostics are counted directly instead.
	results := analysistest.Run(discard{}, analysistest.TestData(), analyzer, dirs...)
	counts := make(map[string]int)

	for _, result := range results {
		for _, diag := range result.Action.Diagnostics {
			counts[diag.Category]++
		}
	}

	return counts
}

// discard swallows the failures analysistest reports, for a run whose diagnostics are
// checked directly rather than against the want comments of the testdata.
type discard struct{}

func (discard) Errorf(string, ...any) {}

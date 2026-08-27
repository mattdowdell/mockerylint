package main

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/mattdowdell/mockerylint"
)

func main() {
	singlechecker.Main(newAnalyzer())
}

// newAnalyzer creates the analyzer and gives each rule a flag named after it, so any rule
// can be turned off on its own with e.g. -usetimes=false.
//
// The flags are declared on the analyzer rather than parsed here, as singlechecker parses
// the command line itself. The analyzer reads the rules as it runs, so the values parsed
// into them take effect.
func newAnalyzer() *analysis.Analyzer {
	opts := mockerylint.DefaultOptions()
	analyzer := mockerylint.New(opts)

	analyzer.Flags.BoolVar(
		&opts.UseFactory,
		mockerylint.RuleUseFactory,
		opts.UseFactory,
		"apply the "+mockerylint.RuleUseFactory+" rule",
	)
	analyzer.Flags.BoolVar(
		&opts.UseExpecter,
		mockerylint.RuleUseExpecter,
		opts.UseExpecter,
		"apply the "+mockerylint.RuleUseExpecter+" rule",
	)
	analyzer.Flags.BoolVar(
		&opts.UseTimes,
		mockerylint.RuleUseTimes,
		opts.UseTimes,
		"apply the "+mockerylint.RuleUseTimes+" rule",
	)
	analyzer.Flags.BoolVar(
		&opts.NoAnything,
		mockerylint.RuleNoAnything,
		opts.NoAnything,
		"apply the "+mockerylint.RuleNoAnything+" rule",
	)

	return analyzer
}

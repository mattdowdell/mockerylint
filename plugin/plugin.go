package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/mattdowdell/mockerylint"
)

//nolint:gochecknoinits // required for golangci-lint plugin
func init() {
	register.Plugin("mockerylint", New)
}

type Plugin struct {
	opts mockerylint.Options
}

func New(settings any) (register.LinterPlugin, error) {
	// The configuration type will be map[string]any or []interface, it depends on your configuration.
	// You can use https://github.com/go-viper/mapstructure to convert map to struct.

	opts, err := register.DecodeSettings[mockerylint.Options](settings)
	if err != nil {
		return nil, err
	}

	return &Plugin{opts}, nil
}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		mockerylint.New(&p.opts),
	}, nil
}

func (p *Plugin) GetLoadMode() string {
	// NOTE: the mode can be `register.LoadModeSyntax` or `register.LoadModeTypesInfo`.
	// - `register.LoadModeSyntax`: if the linter doesn't use types information.
	// - `register.LoadModeTypesInfo`: if the linter uses types information.

	return register.LoadModeTypesInfo
}

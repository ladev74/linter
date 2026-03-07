package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/ladev74/linter/internal/analyzer"
	"github.com/ladev74/linter/internal/config"
)

func init() {
	register.Plugin(analyzer.Name, New)
}

type plugin struct {
	cfg analyzer.Config
}

var _ register.LinterPlugin = (*plugin)(nil)

func New(settings any) (register.LinterPlugin, error) {
	cfg, err := register.DecodeSettings[analyzer.Config](settings)
	if err != nil {
		return nil, err
	}

	if cfg.ConfigPath != "" {
		fileCfg, err := config.New(cfg.ConfigPath)
		if err != nil {
			return nil, err
		}

		cfg.Rules = fileCfg.Analyzer.Rules
	}

	return &plugin{
		cfg: cfg,
	}, nil
}

func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	a := analyzer.New(&p.cfg)

	return []*analysis.Analyzer{a}, nil
}

func (p *plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

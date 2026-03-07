package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/ladev74/linter/internal/analyzer/rules"
)

func TestIsSupportedLoggerPath(t *testing.T) {
	if !isSupportedLoggerPath(zapPath) {
		t.Fatal("expected zap to be supported")
	}

	if !isSupportedLoggerPath(slogPath) {
		t.Fatal("expected slog to be supported")
	}

	if isSupportedLoggerPath("fmt") {
		t.Fatal("fmt should not be supported")
	}
}

func TestAnalyzer(t *testing.T) {
	cfg := Config{
		Rules: rules.Config{
			LowercaseStart: rules.LowercaseStart{
				Enabled: true,
			},
			EnglishOnly: rules.EnglishOnly{
				Enabled: true,
			},
			DisallowSpecialCharacters: rules.DisallowSpecialCharacters{
				Enabled: true,
			},
			DisallowSensitiveData: rules.DisallowSensitiveData{
				Enabled: true,
				Patterns: []string{
					"pass",
					"api_key",
				},
			},
		}}

	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, New(&cfg),
		"english_only",
		"lowercase_start",
		"disallow_special_characters",
		"disallow_sensitive_data",
		//"./...",
	)
}

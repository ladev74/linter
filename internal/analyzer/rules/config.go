package rules

type Config struct {
	LowercaseStart            LowercaseStart            `yaml:"lowercase_start" json:"lowercase_start"`
	EnglishOnly               EnglishOnly               `yaml:"english_only" json:"english_only"`
	DisallowSpecialCharacters DisallowSpecialCharacters `yaml:"disallow_special_characters" json:"disallow_special_characters"`
	DisallowSensitiveData     DisallowSensitiveData     `yaml:"disallow_sensitive_data" json:"disallow_sensitive_data"`
}

type LowercaseStart struct {
	Enabled bool `yaml:"enabled" json:"enabled"`
}
type EnglishOnly struct {
	Enabled bool `yaml:"enabled" json:"enabled"`
}

type DisallowSpecialCharacters struct {
	Enabled bool `yaml:"enabled" json:"enabled"`
}

type DisallowSensitiveData struct {
	Enabled  bool     `yaml:"enabled" json:"enabled"`
	Patterns []string `yaml:"patterns" json:"patterns"`
}

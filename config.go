package openhab

import "github.com/jinzhu/copier"

// Client for openHAB. It's using openHAB REST API internally.
type Config struct {
	Enabled  bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	APIToken string `json:"apiToken,omitempty" yaml:"apiToken,omitempty"`
	API      string `json:"api,omitempty" yaml:"api,omitempty"`
	User     string `json:"user,omitempty" yaml:"user,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
}

// Clone return copy
func (t *Config) Clone() *Config {
	c := &Config{}
	copier.Copy(&c, &t)
	return c
}

func ExampleConfig() *Config {
	return &Config{
		Enabled:  true,
		APIToken: "APIToken",
		API:      "https://...",
	}
}

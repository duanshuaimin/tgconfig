package example2

import (
	telegraf "github.com/influxdata/tgconfig"
	"github.com/influxdata/tgconfig/plugins/inputs"
)

// Config contains configuration for Example2.  It's structure
// must match the data in the configuration file or source.
type Config struct {
	Value string `toml:"value"`
}

// Example2 is an example input plugin.
type Example2 struct {
	Value string
}

// New creates an Example2 from an Config.
func New(config *Config) ([]telegraf.Input, error) {
	return []telegraf.Input{&Example2{config.Value}}, nil
}

func (p *Example2) Gather() error {
	return nil
}

func init() {
	inputs.Add("example2", New)
}

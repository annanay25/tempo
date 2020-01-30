package frigg

import (
	"github.com/spf13/viper"
)

// Config holds the configuration for redbull.
type Config struct {
	Backend string
}

// InitFromViper initializes the options struct with values from Viper
func (c *Config) InitFromViper(v *viper.Viper) {
	c.Backend = v.GetString("backend")
}
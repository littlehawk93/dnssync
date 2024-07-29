package config

import "github.com/littlehawk93/dnssync/provider/cloudflare"

// Configuration configuration for various providers
type Configuration struct {
	Cloudflare *cloudflare.Provider `mapstructure:"cloudflare"`
}

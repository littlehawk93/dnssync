package config

import (
	"strings"

	"github.com/littlehawk93/dnssync/provider"
	"github.com/littlehawk93/dnssync/provider/cloudflare"
	"github.com/littlehawk93/dnssync/provider/namesilo"
)

// Configuration configuration for various providers
type Configuration struct {
	Cloudflare *cloudflare.Provider `mapstructure:"cloudflare"`
	NameSilo   *namesilo.Provider   `mapstructure:"namesilo"`
}

func (me Configuration) GetMatchingProvider(name string) provider.Provider {

	name = strings.TrimSpace(strings.ToLower(name))

	providers := me.GetProviders()

	for _, p := range providers {

		if p != nil && p.GetName() == name {
			return p
		}
	}
	return nil
}

// GetProviders get a set of all providers in this configuration
func (me Configuration) GetProviders() []provider.Provider {

	return []provider.Provider{
		me.Cloudflare,
		me.NameSilo,
	}
}

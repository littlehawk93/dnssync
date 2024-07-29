package provider

import "net"

// Provider defines a client for a DNS provider
type Provider interface {
	UpdateIP(ip net.IP, name string, ttl int, force bool) error
	GetName() string
}

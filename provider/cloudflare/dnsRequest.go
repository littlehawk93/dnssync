package cloudflare

// dnsRequest a DNS record request to be sent to the Cloudflare API
type dnsRequest struct {
	ID      string   `json:"id"`
	Content string   `json:"content"`
	Name    string   `json:"name"`
	Proxied bool     `json:"proxied"`
	Type    string   `json:"type"`
	Comment string   `json:"comment"`
	TTL     int      `json:"ttl"`
	Tags    []string `json:"tags"`
}

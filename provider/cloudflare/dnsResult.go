package cloudflare

import "time"

// dnsResult Cloudflare DNS API result
type dnsResult struct {
	ID         string                 `json:"id"`
	Content    string                 `json:"content"`
	Name       string                 `json:"name"`
	Proxied    bool                   `json:"proxied"`
	Proxiable  bool                   `json:"proxiable"`
	Type       string                 `json:"type"`
	Comment    string                 `json:"comment"`
	CreatedOn  time.Time              `json:"created_on"`
	ModifiedOn time.Time              `json:"modified_on"`
	TTL        int                    `json:"ttl"`
	Meta       map[string]interface{} `json:"meta"`
	Tags       []string               `json:"tags"`
}

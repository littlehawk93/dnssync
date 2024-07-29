package cloudflare

// response an API response from CloudFlare
type response[T any] struct {
	Errors   []interface{}
	Messages []interface{}
	Success  bool
	Result   T
}

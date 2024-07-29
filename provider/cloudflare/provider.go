package cloudflare

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
)

const (
	apiHost              = "api.cloudflare.com"
	apiDnsListPathFormat = "/client/v4/zones/%s/dns_records"
	apiDnsPathFormat     = "/client/v4/zones/%s/dns_records/%s"
)

type Provider struct {
	ZoneID string `mapstructure:"zone_id"`
	Token  string `mapstructure:"token"`
}

// GetName returns this provider's name
func (me *Provider) GetName() string {
	return "cloudflare"
}

// UpdateIP update the IP address for a particular DNS record
func (me *Provider) UpdateIP(ip net.IP, name string, ttl int, force bool) error {

	request := dnsRequest{
		Name:    name,
		Content: ip.String(),
		TTL:     ttl,
		Type:    "A",
		Proxied: false,
		Comment: "",
		Tags:    nil,
	}

	dnsResults, err := me.sendDNSListRequest(me.ZoneID)

	if err != nil {
		return err
	}

	for _, result := range dnsResults.Result {
		if result.Name == name {
			if !force && result.Content == ip.String() {
				return nil
			}
			_, err := me.sendDNSRequest(result.ID, request)
			return err
		}
	}
	return errors.New("no matching DNS record")
}

func (me *Provider) sendDNSListRequest(zoneID string) (response[[]dnsResult], error) {

	var result response[[]dnsResult]

	u := url.URL{
		Scheme: "https",
		Host:   apiHost,
		Path:   fmt.Sprintf(apiDnsListPathFormat, zoneID),
	}

	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return result, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", me.Token))

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return result, err
	}

	if resp.StatusCode != 200 {
		return result, fmt.Errorf("api error: %s", string(respBytes))
	}

	if err = json.Unmarshal(respBytes, &result); err != nil {
		return result, err
	}
	return result, nil
}

func (me *Provider) sendDNSRequest(dnsID string, dns dnsRequest) (response[dnsResult], error) {

	var result response[dnsResult]

	u := url.URL{
		Scheme: "https",
		Host:   apiHost,
		Path:   fmt.Sprintf(apiDnsPathFormat, me.ZoneID, dnsID),
	}

	jsonBytes, _ := json.Marshal(&dns)
	reqBody := bytes.NewBuffer(jsonBytes)

	req, err := http.NewRequest("PUT", u.String(), reqBody)

	if err != nil {
		return result, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", me.Token))

	client := http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return result, err
	}

	if resp.StatusCode != 200 {
		return result, fmt.Errorf("api error: %s", string(respBytes))
	}

	if err = json.Unmarshal(respBytes, &result); err != nil {
		return result, err
	}
	return result, nil
}

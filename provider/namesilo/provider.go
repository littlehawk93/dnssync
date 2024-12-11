package namesilo

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
)

const (
	APIHost     string = "www.namesilo.com"
	APIVersion  int    = 1
	APIDataType string = "json"
)

type Provider struct {
	Key string `mapstructure:"key"`
}

// GetName returns this provider's name
func (me *Provider) GetName() string {
	return "namesilo"
}

// UpdateIP update the IP address for a particular DNS record
func (me *Provider) UpdateIP(ip net.IP, name string, ttl int, force bool) error {

	response, err := me.getDnsValues(name)

	if err != nil {
		return err
	}

	for _, record := range response.Reply.ResourceRecords {
		if record.Type == "A" && (force || record.Value != ip.String()) {
			if err := me.updateDnsRecord(record.Host, record.RecordID, ip, ttl); err != nil {
				return err
			}
		}
	}
	return nil
}

func (me *Provider) getDnsValues(domain string) (dnsListResponse, error) {

	var result dnsListResponse

	req, err := me.createApiRequest("dnsListRecords", map[string]string{"domain": domain})

	if err != nil {
		return result, err
	}

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

func (me *Provider) updateDnsRecord(domain, recordId string, ip net.IP, ttl int) error {

	req, err := me.createApiRequest("dnsUpdateRecord", map[string]string{
		"domain":  domain,
		"rrid":    recordId,
		"host":    domain,
		"rrvalue": ip.String(),
		"rrttl":   string(ttl),
	})

	if err != nil {
		return err
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("api error: %s", string(respBytes))
	}
	return nil
}

func (me *Provider) createApiRequest(operation string, query map[string]string) (*http.Request, error) {

	u := &url.URL{
		Scheme: "https",
		Host:   APIHost,
		Path:   fmt.Sprintf("/api/%s", operation),
	}

	q := u.Query()

	q.Add("version", string(APIVersion))
	q.Add("type", APIDataType)
	q.Add("key", me.Key)

	for k, v := range query {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	return http.NewRequest("GET", u.String(), nil)
}

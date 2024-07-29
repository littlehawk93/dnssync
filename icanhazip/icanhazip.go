package icanhazip

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiHost string = "icanhazip.com"
)

// Return this machine's public facing IP address
func GetIP() (net.IP, error) {

	u := url.URL{
		Scheme: "http",
		Host:   apiHost,
		Path:   "",
	}

	resp, err := http.Get(u.String())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	ipStr := strings.TrimSpace(string(bytes))

	ip := net.ParseIP(ipStr)

	if ip == nil {
		return nil, fmt.Errorf("'%s' is not a valid IP address", ipStr)
	}
	return ip, nil
}

package ip

import "github.com/go-zoox/fetch"

// IPRemote is the remote ip echo site.
const IPRemote = "https://ip.zcorky.com"

// GetPublicIP gets the public ip.
func GetPublicIP() (string, error) {
	response, err := fetch.Get(IPRemote)
	if err != nil {
		return "", err
	}

	return response.String(), nil
}

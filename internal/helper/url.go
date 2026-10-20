package helper

import (
	"errors"
	"net"
	"net/url"
	"strings"
)

var (
	ErrURLRequired   = errors.New("URL is required")
	ErrInvalidURL    = errors.New("invalid URL")
	ErrInvalidScheme = errors.New("invalid Scheme URL")
)

func ValidateURL(raw string) (string, error) {
	raw = strings.TrimSpace(raw)

	if raw == "" {
		return "", ErrURLRequired
	}

	u, err := url.Parse(raw)
	if err != nil {
		return "", ErrInvalidURL
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return "", ErrInvalidScheme
	}

	if u.Host == "" || u.Hostname() == "" {
		return "", ErrInvalidURL
	}

	if u.User != nil {
		return "", ErrInvalidURL
	}

	// Reject malformed ports and invalid port ranges.
	if port := u.Port(); port != "" {
		if portNumber, err := net.LookupPort("tcp", port); err != nil || portNumber < 1 || portNumber > 65535 {
			return "", ErrInvalidURL
		}
	}

	return u.String(), nil
}

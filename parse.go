package cidr

import (
	"net"
)

// An ParsedCIDR contains
// FirstIP the first ip of the parsed range
// LasttIP the last ip of the parsed range
// IPNet represents an IP network
// IsIPv4 check for IPv4
// IsIPv6 check for IPv6
type ParsedCIDR struct {
	FirstIP net.IP
	LastIP  net.IP
	IPNet   *net.IPNet
	IsIPv4  bool
	IsIPv6  bool
}

// ParseCIDR return ParsedCIDR for the provided IP Range
func ParseCIDR(s string) (*ParsedCIDR, error) {
	firstIP, ipNet, err := net.ParseCIDR(s)
	if err != nil {
		return nil, err
	}

	var v4, v6 bool

	if firstIP.To4() != nil {
		v4 = true
	} else {
		v6 = true
	}

	parsed := ParsedCIDR{
		FirstIP: firstIP,
		IPNet:   ipNet,
		IsIPv4:  v4,
		IsIPv6:  v6,
	}

	if parsed.IsIPv4 {
		parsed.LastIP = parsed.getLastIPv4()
	} else {
		parsed.LastIP = parsed.getLastIPv6()
	}

	return &parsed, nil
}

package cidr

import (
	"encoding/binary"
	"net"
)

func (p *ParsedCIDR) getLastIPv4() net.IP {
	if p.IsIPv4 != true {
		return nil
	}

	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, p.LastIPv4())

	return ip
}

// FirstIPv4 return Decimal FirstIP
func (p *ParsedCIDR) FirstIPv4() uint32 {
	if p.IsIPv4 != true {
		return 0
	}

	ip := p.FirstIP.To4()
	var firstIP uint32
	firstIP += uint32(ip[0]) << 24
	firstIP += uint32(ip[1]) << 16
	firstIP += uint32(ip[2]) << 8
	firstIP += uint32(ip[3])

	return firstIP
}

// HostCountIPv4 return number of IPs on the parsed range
func (p *ParsedCIDR) HostCountIPv4() uint32 {
	ones, bits := p.IPNet.Mask.Size()
	return uint32(1 << (bits - ones))
}

// LastIPv4 return Decimal LastIP
func (p *ParsedCIDR) LastIPv4() uint32 {
	if p.IsIPv4 != true {
		return 0
	}

	return p.FirstIPv4() + p.HostCountIPv4() - 1
}

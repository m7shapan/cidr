package cidr

import (
	"math/big"
	"net"
)

func (p *ParsedCIDR) getLastIPv6() net.IP {
	if p.IsIPv6 != true {
		return nil
	}

	ip := make(net.IP, 16)
	ip = p.LastIPv6().Bytes()
	return ip
}

// FirstIPv6 return Decimal FirstIP
func (p *ParsedCIDR) FirstIPv6() *big.Int {
	if p.IsIPv6 != true {
		return nil
	}

	IPInt := big.NewInt(0)
	return IPInt.SetBytes(p.FirstIP.To16())
}

// HostCountIPv6 return number of IPs on the parsed range
func (p *ParsedCIDR) HostCountIPv6() *big.Int {
	ones, bits := p.IPNet.Mask.Size()
	var max = big.NewInt(1)

	return max.Lsh(max, uint(bits-ones))
}

// LastIPv6 return Decimal LastIP
func (p *ParsedCIDR) LastIPv6() *big.Int {
	if p.IsIPv6 != true {
		return nil
	}

	IPInt := p.FirstIPv6()
	return IPInt.Add(IPInt, big.NewInt(0).Sub(p.HostCountIPv6(), big.NewInt(1)))
}

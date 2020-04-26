package cidr

import (
	"math/big"
	"net"
)

// IPv4tod convert IPv4 to decimal
// in case IP is not IPv4 will return 0
func IPv4tod(ip net.IP) uint32 {
	if ip.To4() == nil {
		return 0
	}

	ip = ip.To4()
	var intIP uint32
	intIP += uint32(ip[0]) << 24
	intIP += uint32(ip[1]) << 16
	intIP += uint32(ip[2]) << 8
	intIP += uint32(ip[3])

	return intIP
}

// IPv6tod convert IPv6 to decimal
// in case IP is not IPv6 will return nil
func IPv6tod(ip net.IP) *big.Int {
	if ip.To4() != nil {
		return nil
	}
	return big.NewInt(0).SetBytes(ip.To16())
}

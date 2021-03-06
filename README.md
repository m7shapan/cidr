# CIDR (Classless inter-domain routing)

This Package converts IP CIDR to range and return First IP, Last IP, First IP decimal, Last IP decimal and Total Host count for IPv4 and IPv6

## Install
`go get -u github.com/m7shapan/cidr`

## How to use
 See [the GoDoc](https://pkg.go.dev/github.com/m7shapan/cidr)

```go
package main

import (
	"fmt"
	"log"

	"github.com/m7shapan/cidr"
)

func main() {
	p, err := cidr.ParseCIDR("1.0.0.0/24")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("First IP:", p.FirstIP)
	if p.IsIPv4 {
		fmt.Println("First IP (Decimal):", p.FirstIPv4())
	} else {
		fmt.Println("First IP (Decimal):", p.FirstIPv6())
	}

	fmt.Println("Last IP:", p.LastIP)

	if p.IsIPv4 {
		fmt.Println("Last IP (Decimal):", p.LastIPv4())
	} else {
		fmt.Println("Last IP (Decimal):", p.LastIPv6())
	}

	if p.IsIPv4 {
		fmt.Println("Total Host:", p.HostCountIPv4())
	} else {
		fmt.Println("Total Host:", p.HostCountIPv6())
	}

	// First IP: 1.0.0.0
	// First IP (Decimal): 16777216
	// Last IP: 1.0.0.255
	// Last IP (Decimal): 16777471
	// Total Host: 256


	ip := net.ParseIP("1.0.0.0")
	fmt.Println(IPv4tod(ip)) // 16777216

	ip := net.ParseIP("2001:4860:4860::8888")
	fmt.Println(IPv6tod(ip)) // 42541956123769884636017138956568135816

	var i uint32 = 16777216
	fmt.Println(DtoIPv4(i)) // 1.0.0.0

	b := new(big.Int)
	b.SetString("42541956123769884636017138956568135816", 10)
	fmt.Println(DtoIPv6(b)) // 2001:4860:4860::8888

}
```

## Package usage for SEO
- get ip from ip range
- convert ip range to ip
- golang convert ip range to decimal
- How to get first/last IP address of CIDR
- Convert IP version 6 address to integer or decimal number
- IPv4 to IP Decimal Conversion
- ipv4 to decimal golang
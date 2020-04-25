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

}
```
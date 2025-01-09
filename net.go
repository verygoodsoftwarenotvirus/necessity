package necessity

import (
	"net"
)

func MustResolveIPAddr(network, address string) *net.IPAddr {
	ipAddr, err := net.ResolveIPAddr(network, address)
	if err != nil {
		panic(err)
	}

	return ipAddr
}

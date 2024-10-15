package addr

import (
	"crypto/tls"
	"errors"
	"net"
	"os"
	"strconv"
	"strings"
)

type Address struct {
	// Host domain name or ipv4, ipv6 address
	Host string
	// Port service port
	Port int
}

func (addr Address) Addr() string {
	return net.JoinHostPort(addr.Host, strconv.Itoa(addr.Port))
}

// ParseAddressesFromEnv reads a comma-separated list of addresses from an environment variable
// and returns a slice of Address structs. If the environment variable is not set or
// contains invalid addresses, it returns an empty slice.
func ParseAddressesFromEnv(key string) ([]*Address, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return nil, nil
	}

	return ParseAddressesFromStr(aux)
}

// ParseAddressFromEnv reads a single address from an environment variable
// and returns an Address struct. If the environment variable is not set or
// contains an invalid address, it returns nil.
func ParseAddressFromEnv(key string) (*Address, error) {
	aux := os.Getenv(key)
	if aux == "" {
		return nil, nil
	}

	return ParseAddressFromStr(aux)
}

// ParseAddressesFromStr parses a comma-separated list of "host:port" strings into a slice of Address structs.
// If any of the addresses is invalid, it will be skipped and an error returned.
func ParseAddressesFromStr(addressesStr string) ([]*Address, error) {
	addressStrings := strings.Split(addressesStr, ",")
	var addresses []*Address
	for _, addressStr := range addressStrings {
		address, err := ParseAddressFromStr(addressStr)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}
	return addresses, nil
}

// ParseAddressFromStr parses a string in the format "host:port" into an Address struct.
// If the format is invalid or the port is not a valid integer, it returns an error.
func ParseAddressFromStr(addressStr string) (*Address, error) {
	host, portStr, err := net.SplitHostPort(addressStr)
	if err != nil {
		return nil, errors.New("invalid address format: " + addressStr)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, errors.New("invalid port number: " + portStr)
	}

	return &Address{Host: host, Port: port}, nil
}

func Dial(addr Address, tlsConfig *tls.Config) (net.Conn, error) {
	if tlsConfig == nil {
		return net.Dial("tcp", addr.Addr())
	} else {
		return tls.Dial("tcp", addr.Addr(), tlsConfig)
	}
}

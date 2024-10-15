package addr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAddressFromStr(t *testing.T) {
	tests := []struct {
		input          string
		expectedHost   string
		expectedPort   int
		expectingError bool
	}{
		// Test for valid IPv4 address
		{"192.168.1.1:8080", "192.168.1.1", 8080, false},
		// Test for invalid IPv4 address (missing port)
		{"192.168.1.1", "", 0, true},
		// Test for valid domain
		{"example.com:80", "example.com", 80, false},
		// Test for invalid domain (missing port)
		{"example.com", "", 0, true},
		// Test for valid IPv6 address
		{"[::1]:8080", "::1", 8080, false},
		// Test for valid IPv6 with full address
		{"[2001:db8::ff00:42:8329]:443", "2001:db8::ff00:42:8329", 443, false},
		// Test for invalid IPv6 address (missing port)
		{"[2001:db8::ff00:42:8329]", "", 0, true},
		// Test for invalid address format
		{"invalid:address", "", 0, true},
	}

	for _, test := range tests {
		addr, err := ParseAddressFromStr(test.input)
		if test.expectingError {
			assert.Error(t, err, "Expected error but got none for input: %s", test.input)
		} else {
			assert.NoError(t, err, "Unexpected error for input: %s", test.input)
			assert.Equal(t, test.expectedHost, addr.Host)
			assert.Equal(t, test.expectedPort, addr.Port)
		}
	}
}

package addr

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestParseAddressesFromEnvWhenEnvIsSet(t *testing.T) {
	key := "TEST_ENV_ADDRESSES"
	err := os.Setenv(key, "zookeeper-0:2181,zookeeper-1:2181")
	require.NoError(t, err)

	expected := []*Address{
		{Host: "zookeeper-0", Port: 2181},
		{Host: "zookeeper-1", Port: 2181},
	}

	result, err := ParseAddressesFromEnv(key)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	_ = os.Unsetenv(key)
}

func TestParseAddressesFromEnvWhenEnvIsNotSet(t *testing.T) {
	key := "TEST_ENV_ADDRESSES"

	result, err := ParseAddressesFromEnv(key)

	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestParseAddressesFromEnvWhenEnvHasInvalidFormat(t *testing.T) {
	key := "TEST_ENV_ADDRESSES"
	err := os.Setenv(key, "invalidAddress")
	require.NoError(t, err)

	// Expect an error when the format is invalid
	result, err := ParseAddressesFromEnv(key)

	assert.Error(t, err)
	assert.Nil(t, result)

	_ = os.Unsetenv(key)
}

func TestParseAddressFromEnvWhenEnvIsSet(t *testing.T) {
	key := "TEST_ENV_ADDRESS"
	err := os.Setenv(key, "zookeeper-0:2181")
	require.NoError(t, err)

	expected := &Address{Host: "zookeeper-0", Port: 2181}

	result, err := ParseAddressFromEnv(key)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	_ = os.Unsetenv(key)
}

func TestParseAddressFromEnvWhenEnvIsNotSet(t *testing.T) {
	key := "TEST_ENV_ADDRESS"

	result, err := ParseAddressFromEnv(key)

	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestParseAddressFromEnvWhenEnvHasInvalidFormat(t *testing.T) {
	key := "TEST_ENV_ADDRESS"
	err := os.Setenv(key, "invalidAddress")
	require.NoError(t, err)

	// Expect an error when the format is invalid
	result, err := ParseAddressFromEnv(key)

	assert.Error(t, err)
	assert.Empty(t, result)

	_ = os.Unsetenv(key)
}

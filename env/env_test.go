package env

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGetOrError(t *testing.T) {
	key := "SOME_ENV_VAR"
	value := "some value"

	t.Run("load environment variable", func(t *testing.T) {
		_ = os.Setenv(key, value)
		v, err := GetOrError(key)
		require.NoError(t, err)
		assert.Equal(t, value, v)
	})

	t.Run("environment variable not defined", func(t *testing.T) {
		_ = os.Unsetenv(key)
		_, err := GetOrError(key)
		assert.Error(t, err)
	})
}

func TestGetOrElse(t *testing.T) {

}

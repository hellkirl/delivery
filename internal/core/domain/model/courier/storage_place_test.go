package courier_test

import (
	"delivery/internal/core/domain/model/courier"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStoragePlace(t *testing.T) {
	t.Run("New Storage Place | Correct Arguments", func(t *testing.T) {
		sp, err := courier.NewStoragePlace("test", 12)
		require.NoError(t, err)
		require.NotNil(t, sp)
	})

	t.Run("New Storage Place | Incorrect Arguments", func(t *testing.T) {
		sp, err := courier.NewStoragePlace("test", -1)
		require.Error(t, err)
		require.Nil(t, sp)
	})

	t.Run("Storage Place | Can Store", func(t *testing.T) {
		
	})
}

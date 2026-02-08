package kernel_test

import (
	"delivery/internal/core/domain/model/kernel"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLocation(t *testing.T) {
	t.Run("Correct Coordinates", func(t *testing.T) {
		loc, err := kernel.NewLocation(10, 10)
		require.NoError(t, err)
		require.NotNil(t, loc)
	})

	t.Run("Incorrect X Coordinate", func(t *testing.T) {
		loc, err := kernel.NewLocation(15, 2)
		require.Error(t, err)
		require.Nil(t, loc)
	})

	t.Run("Incorrect Y Coordinate", func(t *testing.T) {
		loc, err := kernel.NewLocation(2, 15)
		require.Error(t, err)
		require.Nil(t, loc)
	})

	t.Run("Incorrect Coordinates", func(t *testing.T) {
		loc, err := kernel.NewLocation(15, 15)
		require.Error(t, err)
		require.Nil(t, loc)
	})

	t.Run("Equal Locations", func(t *testing.T) {
		loc1, err := kernel.NewLocation(10, 10)
		require.NoError(t, err)
		loc2, err := kernel.NewLocation(10, 10)
		require.NoError(t, err)
		require.True(t, loc1.Equals(loc2))
	})

	t.Run("Unequal Locations", func(t *testing.T) {
		loc1, err := kernel.NewLocation(5, 7)
		require.NoError(t, err)
		loc2, err := kernel.NewLocation(10, 10)
		require.NoError(t, err)
		require.False(t, loc1.Equals(loc2))
	})

	t.Run("Calculate distance", func(t *testing.T) {
		loc1, err := kernel.NewLocation(4, 9)
		require.NoError(t, err)
		loc2, err := kernel.NewLocation(2, 6)
		require.NoError(t, err)

		var correctDistance uint8 = 5
		calculatedDistance := loc1.CalculateDistance(loc2)
		require.Equal(t, correctDistance, calculatedDistance)
	})
}

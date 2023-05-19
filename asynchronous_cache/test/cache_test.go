package test

import (
	"asynchronous_cache/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Cache(t *testing.T) {
	t.Parallel()

	cache := storage.NewCache()

	t.Run("correctly storage value", func(t *testing.T) {
		t.Parallel()
		key := "someKey"
		value := "someValue"

		err := cache.Set(key, value)
		assert.NoError(t, err)

		total := cache.TotalCount()
		assert.Equal(t, total, int64(1))

		storageValue, err := cache.Get(key)
		assert.NoError(t, err)

		assert.Equal(t, value, storageValue)
	})

	t.Run("correctly updated value", func(t *testing.T) {
		t.Parallel()
		key := "someKey"
		value := "someValue"

		err := cache.Set(key, value)
		assert.NoError(t, err)

		storageValue, err := cache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, value, storageValue)

		newValue := "someValue2"
		err = cache.Set(key, newValue)
		assert.NoError(t, err)

		newStorageValue, err := cache.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, newValue, newStorageValue)

		total := cache.TotalCount()
		assert.Equal(t, total, int64(1))
	})

	t.Run("no data races", func(t *testing.T) {
		t.Parallel()

		parallelFactor := 100_000
		emulateLoad(t, cache, parallelFactor)
	})
}

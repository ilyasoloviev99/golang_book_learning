package test

import (
	"asynchronous_cache/helpers"
	"asynchronous_cache/storage"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func emulateLoad(t *testing.T, c storage.CacheWithMetrics, parallelFactor int) {
	wg := sync.WaitGroup{}

	for i := 0; i < parallelFactor; i++ {
		key := fmt.Sprintf("%d-key", i)
		value := fmt.Sprintf("%d-value", i)

		wg.Add(1)
		go func(k string) {
			err := c.Set(k, value)
			assert.NoError(t, err)
			wg.Done()
		}(key)

		wg.Add(1)

		go func(k, v string) {
			storageValue, err := c.Get(k)
			if !errors.Is(err, helpers.ErrNotFound) {
				assert.Equal(t, v, storageValue)
			}
			wg.Done()
		}(key, value)

		wg.Add(1)
		go func(k string) {
			err := c.Delete(k)
			assert.NoError(t, err)
			wg.Done()
		}(key)
	}

	wg.Wait()
}

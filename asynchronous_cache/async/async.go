package async

import (
	"asynchronous_cache/helpers"
	"asynchronous_cache/storage"
	"context"
)

type Cache struct {
	c *storage.Cache
}

func InitAsyncCache() *Cache {
	return &Cache{
		c: storage.NewCache(),
	}
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	ch := make(chan string)
	go func() {
		defer close(ch)
		v, ok := c.c.Get(key)
		if ok {
			ch <- v
		}
	}()

	select {
	case <-ctx.Done():
		return "", helpers.ErrTimeout
	case x, ok := <-ch:
		if ok {
			return x, nil
		}
		return "", helpers.ErrNotFound
	}
}

func (c *Cache) Add(ctx context.Context, key, value string) error {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		c.c.Set(key, value)
	}()

	select {
	case <-ctx.Done():
		return helpers.ErrTimeout
	case <-ch:
		return nil
	}
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		c.c.Delete(key)
	}()

	select {
	case <-ctx.Done():
		return helpers.ErrTimeout
	case <-ch:
		return nil
	}
}

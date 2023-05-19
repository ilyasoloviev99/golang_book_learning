package storage

import (
	"sync"
	"time"
)

var timeout = time.Millisecond * 2

type Cache struct {
	messages map[string]string
	mu       *sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		messages: make(map[string]string),
		mu:       new(sync.RWMutex),
	}
}

func (c *Cache) Set(key, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	time.Sleep(timeout)
	c.messages[key] = value

	return nil
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	time.Sleep(timeout)

	messages, ok := c.messages[key]

	return messages, ok
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	time.Sleep(timeout)

	delete(c.messages, key)
}

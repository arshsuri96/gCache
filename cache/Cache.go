package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	lock sync.RWMutex
	data map[string][]byte
}

func New() *Cache {
	return &Cache{
		data: make(map[string][]byte),
	}
}

func (c *Cache) Set(key, value []byte, ttl time.Duration) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data[string(key)] = value

	return nil
}

func (c *Cache) Has(key []byte) bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	_, ok := c.data[string(key)]

	return ok
}

func (c *Cache) Delete(key []byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.data, string(key))

	return nil
}

func (c *Cache) Get(key []byte) ([]byte, error) {
	c.lock.RLock()
	defer c.lock.Unlock()

	value, ok := c.data[string(key)]
	if !ok {
		return nil, fmt.Errorf("Key (%s) not found", string(key))
	}

	return value, nil

}

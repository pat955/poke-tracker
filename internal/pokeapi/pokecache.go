package pokeapi

// TODO
// Remove cache after 1-30mins
import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu      *sync.RWMutex
	entries map[string]CacheEntry
}

type CacheEntry struct {
	timeCreated time.Time
	val         []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = CacheEntry{time.Now(), val}
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, found := c.entries[key]
	if !found {
		return nil, false
	}
	return entry.val, true
}

func NewCache(interval time.Duration) Cache {
	c := Cache{&sync.RWMutex{}, make(map[string]CacheEntry)}
	go c.ReapLoop(interval)
	return c

}

func (c *Cache) ReapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now(), interval)
	}

}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.entries {
		if v.timeCreated.Before(now.Add(-last)) {
			fmt.Println(k, " deleted!")
			delete(c.entries, k)
		}
	}
}

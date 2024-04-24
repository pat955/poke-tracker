package pokeapi

// TODO
// Remove cache after 1-30mins
import (
	"fmt"
	"sync"
	"time"
)

type DataTypes interface {
	String() string
	GetID() int
	GetURL() string
}

type Cache struct {
	mu      *sync.RWMutex
	entries map[string]CacheEntry // string is a unique id
}

type CacheEntry struct {
	timeCreated time.Time
	dataStruct  DataTypes
}

func (c *Cache) Print() {
	for i, entry := range c.entries {
		fmt.Println("Nr."+i, entry.dataStruct.String())
	}
}
func (c *Cache) Add(keyID string, dataType DataTypes) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[keyID] = CacheEntry{
		timeCreated: time.Now(),
		dataStruct:  dataType,
	}
}
func (c *Cache) Get(key string) (DataTypes, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, found := c.entries[key]
	if !found {
		return nil, false
	}
	return entry.dataStruct, true
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

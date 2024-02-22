package main

import (
    "fmt"
)

type Cache interface {
    Get(k string) (string, bool)
    Set(k, v string)
}

var _ Cache = (*cacheImpl)(nil)

// Will refine the cache constructor and methods so that they match the Cache interface
func newCacheImpl() *cacheImpl {
    return &cacheImpl{data: make(map[string]string)}
type cacheImpl struct {
    data map[string]string
}

func (c *cacheImpl) Get(k string) (string, bool) {
    v, ok := c.data[k]
    return v, ok
}

func (c *cacheImpl) Set(k, v string) {
    c.data[k] = v
}

func newDbImpl(cache Cache) *dbImpl {
    return &dbImpl{cache: cache, dbs: map[string]string{"hello": "world", "test": "test"}}
}

type dbImpl struct {
    cache Cache
    dbs   map[string]string
}

func (d *dbImpl) Get(k string) (string, bool) {
    v, ok := d.cache.Get(k)
    if ok {
        return fmt.Sprintf("answer from cache: key: %s, val: %s", k, v), ok
    }
    v, ok = d.dbs[k]
    return fmt.Sprintf("answer from dbs: key: %s, val: %s", k, v), ok
}

func main() {
    c := newCacheImpl()
    db := newDbImpl(c)
    fmt.Println(db.Get("test"))
    fmt.Println(db.Get("hello"))
}

// package main

import (
    "fmt"
    "sync"
)

type DBSimple struct {
    data map[string]string
    mu   sync.RWMutex
}

type Cache struct {
    data map[string]string
    mu   sync.RWMutex
}

func (d *DBSimple) AddData(key, value string) {
    d.mu.Lock()
    defer d.mu.Unlock()
    d.data[key] = value
}

func (c *Cache) FillCache(d *DBSimple) {
    d.mu.RLock()
    defer d.mu.RUnlock()
    for key, value := range d.data {
        c.mu.Lock()
        c.data[key] = value
        c.mu.Unlock()
    }
}

func (c *Cache) GetData(key string) string {
    c.mu.RLock()
    defer c.mu.RUnlock()
    return c.data[key]
}

func main() {
    db := &DBSimple{data: make(map[string]string)}
    db.AddData("key1", "value1")

    cache := &Cache{data: make(map[string]string)}
    cache.FillCache(db)

    // Check if the data is being returned from the cache
    fmt.Println(cache.GetData("key1")) // Output: value1
}
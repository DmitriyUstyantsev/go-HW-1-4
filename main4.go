//Простой механизм кэширования с функциями “Get” и “Set
//Код определяет реализацию кэша (cacheImpl), которая соответствует интерфейсу кэша, структуре DBSimple и структуре dbImpl, которая использует кэш.
//Функция main демонстрирует простой сценарий использования.
//Когда db.Вызывается Get, он сначала проверяет кэш, а затем базовое хранилище данных соответствующим образом.к

package main

import (
	"fmt"
	"sync"
)

type Cache interface {
	Get(k string) (string, bool)
	Set(k, v string)
}

func newCacheImpl() *cacheImpl {
	return &cacheImpl{data: make(map[string]string)}
}

type cacheImpl struct {
	data map[string]string
	sync.RWMutex
}

func (c *cacheImpl) Get(k string) (string, bool) {
	c.RLock()
	defer c.RUnlock()
	v, ok := c.data[k]
	return v, ok
}

func (c *cacheImpl) Set(k, v string) {
	c.Lock()
	defer c.Unlock()
	c.data[k] = v
}

type DBSimple struct {
	data map[string]string
	sync.RWMutex
}

func (d *DBSimple) AddData(key, value string) {
	d.Lock()
	defer d.Unlock()
	d.data[key] = value
}

type dbImpl struct {
	cache Cache
	dbs   DBSimple
}

func newDbImpl(cache Cache) *dbImpl {
	return &dbImpl{
		cache: cache,
		dbs:   DBSimple{data: map[string]string{"hello": "world", "test": "test"}},
	}
}

func (d *dbImpl) Get(k string) (string, bool) {
	v, ok := d.cache.Get(k)
	if ok {
		return fmt.Sprintf("answer from cache: key: %s, val: %s", k, v), ok
	}
	d.RLock()
	defer d.RUnlock()
	v, ok = d.dbs.data[k]
	return fmt.Sprintf("answer from dbs: key: %s, val: %s", k, v), ok
}

func main() {
	c := newCacheImpl()
	db := newDbImpl(c)
	fmt.Println(db.Get("test"))
	fmt.Println(db.Get("hello"))
}

package cache

import (
//	"fmt"
	"time"
)

type vt struct {
	v string
	t time.Time
}

type Cache struct {
	m map[string]vt
}

func NewCache() Cache {
	return Cache{map[string]vt{}}
}

func (c Cache) Get(key string) (string, bool) {
	if vt, ok := c.m[key]; !ok {
		return "", false 
	} else if time.Now().Before(vt.t) {
		return "", false
	}
	return c.m[key].v, true
}

func (c *Cache) Put(key, value string) {
	c.m[key] = vt{v: value, t: time.Unix(1<<63-1,0)}
}

func (c Cache) Keys() []string {
	sl := []string{}
	for key, _ := range c.m {
		sl = append(sl, key)
	}
	return sl
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.Update()
	c.m[key] = vt{v: value, t: deadline}
}

func (c *Cache) Update() {
	for key, vt := range c.m {
		if time.Now().Before(vt.t) {
			delete(c.m, key)
		}
	}
}

package cache

import "time"

// Struct for values in Cache map
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
	//Before getting values - update Cache
	c.Update()
	if _, ok := c.m[key]; !ok {
		return "", false
	}
	return c.m[key].v, true
}

// First idea to avoid complex solution with time. In more reliable solution it
// can be boolean flag for checking or not expire of the time
func (c *Cache) Put(key, value string) {
	c.m[key] = vt{v: value, t: time.Unix(1<<62-1, 0)}
}

func (c Cache) Keys() []string {
	sl := []string{}
	for key, _ := range c.m {
		sl = append(sl, key)
	}
	return sl
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.m[key] = vt{v: value, t: deadline}
}

func (c *Cache) Update() {
	for key, vt := range c.m {
		// Checks expire of the time value
		if time.Now().After(vt.t) {
			delete(c.m, key)
		}
	}
}

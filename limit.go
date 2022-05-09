package limit

import (
	"sync"
)

type Limit struct {
	sync.Mutex
	LimitKeyMap map[interface{}]int64
}

// InitLimit 初始化Limit
func InitLimit() *Limit {
	return &Limit{
		LimitKeyMap: make(map[interface{}]int64),
	}
}

// Exists 判断限制是否存在
func (c *Limit) Exists(key interface{}) bool {
	c.Lock()
	defer c.Unlock()

	_, ok := c.LimitKeyMap[key]
	if !ok {
		c.LimitKeyMap[key] = 1
	}
	return ok
}

// Delete 删除限制
func (c *Limit) Delete(key interface{}) {
	c.Lock()
	defer c.Unlock()

	delete(c.LimitKeyMap, key)
}


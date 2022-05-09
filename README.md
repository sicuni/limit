# limit
-- import "github.com/sicuni/limit"

## Usage
**func InitLimit**
```bigquery
func InitLimit() *Limit
```
InitLimit returns Initialize the constraint structure

**func Exists**
```bigquery
func (c *Limit) Exists(key interface{}) bool
```
Exists returns whether the key has a limit, if not add a limit

**func Delete**

```bigquery
func (c *Limit) Delete(key interface{})
```
Delete restrictions on the key


**Type Limit**
```bigquery
type Limit struct {
	sync.Mutex
	LimitKeyMap map[interface{}]int64
}
```

**Test**
```bigquery
package limit

import (
	"fmt"
	"testing"
)

func TestInitLimit(t *testing.T) {
	limit := InitLimit()

	str := "111"
if limit.Exists(str) {
		fmt.Println("limit str:", str)
	}
	// 此处被limit
if limit.Exists(str) {
		fmt.Println("limit str:", str)
	}

	key := 0
if limit.Exists(key) {
		fmt.Println("limit key:", key)
	}
	limit.Delete(key) // 解除限制
if limit.Exists(key) {
		fmt.Println("limit key:", key)
	}
}

```
package HandleIpRequest

import (
	"errors"
	"sync"
	"sync/atomic"
)

// Counter Stores counts associated with a key.
type Counter struct {
	m sync.Map
}

// Get Retrieves the count without modifying it
func (c *Counter) Get(key string) (int64, error) {
	err := errors.New("failed")
	count, ok := c.m.Load(key) // khi nhập key vào hàm này thì nó sẽ trả về giá trị của key đó có trong map nếu có hoặc nil nếu key đó không có giá trị,nếu key này có giá trị thì nó sẽ trả về biến bool là OK
	if ok {                    //
		return atomic.LoadInt64(count.(*int64)), nil // hàm atomic.LoadInt64 này sẽ trả về giá trị int64 khi được truyền tham số vào
	}
	return 0, err
}

// Add Adds value to the stored underlying value if it exists.
// If it does not exist, the value is assigned to the key.
func (c *Counter) Add(key string, value int64) int64 {

	count, loaded := c.m.LoadOrStore(key, &value) //hàm này sẽ trả về giá trị của key có trong map nếu key đó có săn trong map. nếu không thì key đó sẽ được lưu và trả về giá trị đã cho mới gán val. nếu key đó có giá trị trong map thì trả về true nếu không có thì trả về false
	if loaded {
		return atomic.AddInt64(count.(*int64), value)
	}
	return *count.(*int64)
}

// DeleteAndGetLastValue Deletes the value associated with the key and retrieves it.
func (c *Counter) DeleteAndGetLastValue(key string) (int64, bool) {

	lastValue, loaded := c.m.LoadAndDelete(key)
	if loaded {
		return *lastValue.(*int64), loaded
	}
	return 0, false
}

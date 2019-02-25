package wurfl

import (
	"fmt"
	"unsafe"
)

// newCStrongCache is not thread safe
// so you have to warm up this cache before using
type cstringCache struct {
	data map[unsafe.Pointer]string
}

func newCStrongCache() *cstringCache {
	return &cstringCache{
		data: map[unsafe.Pointer]string{},
	}
}

func (ch *cstringCache) Get(key unsafe.Pointer) (val string, ok bool) {
	val, ok = ch.data[key]
	return
}

func (ch *cstringCache) Set(key unsafe.Pointer, val string) {
	ch.data[key] = val
}

func (ch *cstringCache) Print() {
	for key, val := range ch.data {
		fmt.Println(">", key, val)
	}
}

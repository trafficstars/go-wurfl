package wurfl

/*
#include <stdlib.h>
*/
import "C"

import "unsafe"

// unsafeCStringCache is not thread safe
// so you have to warm up this cache before using
type unsafeCStringCache struct {
	data map[string]*C.char
}

func newUnsafeCStringCache() *unsafeCStringCache {
	return &unsafeCStringCache{
		data: map[string]*C.char{},
	}
}

func (ch *unsafeCStringCache) Get(key string) (val *C.char, ok bool) {
	val, ok = ch.data[key]
	return
}

func (ch *unsafeCStringCache) Set(key string, val *C.char) {
	ch.data[key] = val
}

func (ch *unsafeCStringCache) Release() {
	for _, v := range ch.data {
		C.free(unsafe.Pointer(v))
	}
}

package wurfl

import (
	"sync"
	"unsafe"
)

type pair struct {
	OriginalKey unsafe.Pointer
	Key         string
	Value       string
}

var (
	al           = 0
	keystorePool = sync.Pool{
		New: func() interface{} { return make(KeyStoreList, 0, 500) },
	}
)

// KeyStoreList contains key and value of the property
type KeyStoreList []pair

// NewKeyStireList get previous object or allocate the new one
func NewKeyStireList() KeyStoreList {
	return keystorePool.Get().(KeyStoreList)
}

// Push item to the store
func (list *KeyStoreList) Push(originalKey unsafe.Pointer, key, value string) {
	*list = append(*list, pair{OriginalKey: originalKey, Key: key, Value: value})
}

// Release the keystore and return to the pool
func (list KeyStoreList) Release() {
	if list != nil {
		keystorePool.Put(list[:0])
	}
}

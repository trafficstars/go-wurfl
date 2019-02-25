package wurfl

/*
// http://www.scientiamobile.com/docs/c-api-user-guide.html/
#cgo LDFLAGS: -lwurfl -L/usr/local/lib
#include <stdlib.h>
#include "./wurfl.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

// Wurfl connection object
type Wurfl struct {
	wurfl              C.wurfl_handle
	cstringCache       *cstringCache
	unsafeCStringCache *unsafeCStringCache
}

// New wurfl database accessor
func New(wurflDatabase string, patches ...string) (*Wurfl, error) {
	var (
		w          = &Wurfl{cstringCache: newCStrongCache(), unsafeCStringCache: newUnsafeCStringCache()}
		wurflError C.wurfl_error
	)

	w.wurfl = C.wurfl_create()

	// Define wurfl DB
	wurflDatabaseStr := C.CString(wurflDatabase)
	wurflError = C.wurfl_set_root(w.wurfl, wurflDatabaseStr)
	C.free(unsafe.Pointer(wurflDatabaseStr))
	if wurflError != C.WURFL_OK {
		errMsg := C.wurfl_get_error_message(w.wurfl)
		return nil, errors.New(C.GoString(errMsg))
	}

	for _, pxml := range patches {
		p := C.CString(pxml)
		C.wurfl_add_patch(w.wurfl, p)
		C.free(unsafe.Pointer(p))
	}

	if C.wurfl_load(w.wurfl) != C.WURFL_OK {
		err := C.wurfl_get_error_message(w.wurfl)
		return nil, errors.New(C.GoString(err))
	}

	// Warm up the cache
	{
		// Init name cache
		for _, name := range keywords {
			w.unsafeCStringCache.Set(name, C.CString(name))
		}

		// Init parameters cache
		for _, agent := range agentsList {
			device := w.Lookup(agent)
			for _, pair := range device.Capabilities {
				if _, ok := w.cstringCache.Get(pair.OriginalKey); !ok {
					w.cstringCache.Set(pair.OriginalKey, pair.Key)
				}
			}
		}
	}

	return w, nil
}

// LookupProperties useragent with specific properties
func (w *Wurfl) LookupProperties(useragent string, proplist []string, vproplist []string) *Device {
	ua := C.CString(useragent)
	device := C.wurfl_lookup_useragent(w.wurfl, ua)
	C.free(unsafe.Pointer(ua))

	if device == nil {
		return nil
	}

	m := NewKeyStireList()
	for _, prop := range proplist {
		cprop, gen := w.getCString(prop)
		val := C.wurfl_device_get_capability(device, cprop)
		if gen {
			C.free(unsafe.Pointer(cprop))
		}
		m.Push(nil, prop, C.GoString(val))
	}

	// get the virtual properties
	mv := NewKeyStireList()
	for _, prop := range vproplist {
		cprop, gen := w.getCString(prop)
		val := C.wurfl_device_get_virtual_capability(device, cprop)
		if gen {
			C.free(unsafe.Pointer(cprop))
		}
		mv.Push(nil, prop, C.GoString(val))
	}

	d := &Device{
		Device:              C.GoString(C.wurfl_device_get_id(device)),
		Capabilities:        m,
		VirtualCapabilities: mv,
	}
	C.wurfl_device_destroy(device)

	return d
}

// Lookup user agent
func (w *Wurfl) Lookup(useragent string) *Device {
	ua := C.CString(useragent)
	device := C.wurfl_lookup_useragent(w.wurfl, ua)
	C.free(unsafe.Pointer(ua))

	if device == nil {
		return nil
	}

	result := C.wurfl_device_get_capabilities(device, (**C.char)(getPointer()))
	m := w.wurflCapabilityEnumerate(result)

	result = C.wurfl_device_get_virtual_capabilities(device, result)
	mv := w.wurflCapabilityEnumerate(result)
	putPointer(unsafe.Pointer(result))

	d := &Device{
		Device:              C.GoString(C.wurfl_device_get_id(device)),
		Capabilities:        m,
		VirtualCapabilities: mv,
	}
	C.wurfl_device_destroy(device)

	return d
}

// Close wurfl connection
func (w *Wurfl) Close() error {
	if w != nil && w.wurfl != nil {
		w.unsafeCStringCache.Release()
		C.wurfl_destroy(w.wurfl)
		w.wurfl = nil
	}
	return nil
}

///////////////////////////////////////////////////////////////////////////////
/// Helper methods
///////////////////////////////////////////////////////////////////////////////

func (w *Wurfl) wurflCapabilityEnumerate(items **C.char) KeyStoreList {
	if items == nil || *items == nil {
		return nil
	}

	m := NewKeyStireList()

	for ; ; items = nextCharStringArrayItem(items, 2) {
		if *items == nil {
			break
		}

		gname := w.getGoString(*items)
		gval := C.GoString(*nextCharStringArrayItem(items, 1))

		m.Push(unsafe.Pointer(*items), gname, gval)
	}

	return m
}

func (w *Wurfl) getGoString(str *C.char) string {
	s, ok := w.cstringCache.Get(unsafe.Pointer(str))
	if !ok {
		s = C.GoString(str)
	}
	return s
}

func (w *Wurfl) getCString(str string) (_ *C.char, gen bool) {
	s, ok := w.unsafeCStringCache.Get(str)
	if !ok {
		s = C.CString(str)
		gen = true
	}
	return s, gen
}

func nextCharStringArrayItem(items **C.char, count uintptr) **C.char {
	return (**C.char)((unsafe.Pointer)((uintptr)(unsafe.Pointer(items)) + unsafe.Sizeof(items)*count))
}

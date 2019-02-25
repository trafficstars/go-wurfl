package wurfl

import (
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// Old library test 30s
// goos: linux
// goarch: amd64
// pkg: informer/wurfl
// Benchmark_Lookup-2   	  500000	     98068 ns/op	    1084 B/op	      23 allocs/op
// PASS
// ok  	informer/wurfl	68.372s

// New library test 30s
// goos: linux
// goarch: amd64
// pkg: informer/lib/wurfl
// Benchmark_Lookup-2             	  300000	    123769 ns/op	   18418 B/op	     491 allocs/op
// Benchmark_LookupProperties-2   	 2000000	     27232 ns/op	    1248 B/op	      13 allocs/op
// PASS
// ok  	informer/lib/wurfl	133.134s
//
// Benchmark_LookupProperties-2 equal to Benchmark_Lookup-2

var (
	props = []string{
		"mobile_browser",
		"is_tablet",
	}
	vprops = []string{
		"form_factor",
		"advertised_browser",
		"advertised_device_os",
		"advertised_device_os_version",
		"is_robot",
		"is_full_desktop",
		"is_mobile",
	}
)

func Benchmark_Lookup(b *testing.B) {
	_, file, _, _ := runtime.Caller(0)
	absfilepath := filepath.Join(filepath.Dir(file), "wurfl.xml.gz")

	if _, err := os.Stat(absfilepath); os.IsNotExist(err) {
		b.Skip("wurfl.xml.gz is not available. We can`t place commercial DB in the public repository")
	}

	wrfl, err := New(absfilepath)
	defer wrfl.Close()

	if err != nil {
		b.Error(err)
		return
	}

	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			agent := agentsList[rand.Int31n(int32(len(agentsList)))]
			device := wrfl.Lookup(agent)
			device.Release()
		}
	})
}

func Benchmark_LookupProperties(b *testing.B) {
	_, file, _, _ := runtime.Caller(0)
	absfilepath := filepath.Join(filepath.Dir(file), "wurfl.xml.gz")

	if _, err := os.Stat(absfilepath); os.IsNotExist(err) {
		b.Skip("wurfl.xml.gz is not available. We can`t place commercial DB in the public repository")
	}

	wrfl, err := New(absfilepath)
	defer wrfl.Close()

	if err != nil {
		b.Error(err)
		return
	}

	b.ResetTimer()
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			agent := agentsList[rand.Int31n(int32(len(agentsList)))]
			device := wrfl.LookupProperties(agent, props, vprops)
			device.Release()
		}
	})
}

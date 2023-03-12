package pprof

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
)

var (
	cpuf       *os.File
	heapf      *os.File
	allocsf    *os.File
	mutexf     *os.File
	blockf     *os.File
	goroutinef *os.File
)

func Stop() {
	pprof.StopCPUProfile()
	cpuf.Close()

	pprof.Lookup("heap").WriteTo(heapf, 0)
	heapf.Close()
	pprof.Lookup("allocs").WriteTo(allocsf, 0)
	allocsf.Close()

	if mp := pprof.Lookup("mutex"); mp != nil {
		mp.WriteTo(mutexf, 0)
	}
	mutexf.Close()
	runtime.SetMutexProfileFraction(0)

	pprof.Lookup("block").WriteTo(blockf, 0)
	blockf.Close()
	runtime.SetBlockProfileRate(0)

	if mp := pprof.Lookup("goroutine"); mp != nil {
		mp.WriteTo(goroutinef, 0)
	}
	goroutinef.Close()
}

func Start(path string) {
	if path == "" {
		path = "./"
	}

	var err error

	cpuf, err = os.Create(filepath.Join(path, "cpu.pprof"))
	if err != nil {
		log.Fatalf("profile: could not create cpu profile: %v", err)
	}
	pprof.StartCPUProfile(cpuf)

	heapf, err = os.Create(filepath.Join(path, "heap.pprof"))
	if err != nil {
		log.Fatalf("profile: could not create memory profile: %v", err)
	}

	allocsf, err = os.Create(filepath.Join(path, "allocs.pprof"))
	if err != nil {
		log.Fatalf("profile: could not create memory profile: %v", err)
	}

	mutexf, err = os.Create(filepath.Join(path, "mutex.pprof"))
	if err != nil {
		log.Fatalf("profile: could not create mutex profile: %v", err)
	}
	runtime.SetMutexProfileFraction(1)

	blockf, err = os.Create(filepath.Join(path, "block.pprof"))
	if err != nil {
		log.Fatalf("profile: could not create block profile: %v", err)
	}
	runtime.SetBlockProfileRate(1)

	goroutinef, err = os.Create(filepath.Join(path, "goroutine.pprof"))
	if err != nil {
		log.Fatalf("profile: could not create goroutine profile: %v", err)
	}
}

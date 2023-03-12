.PHONY: genpdf
genpdf: cpu.pdf heap.pdf allocs.pdf mutex.pdf block.pdf goroutine.pdf

cpu.pdf: cpu.pprof
	go tool pprof --pdf cpu.pprof > cpu.pdf

heap.pdf: heap.pprof
	go tool pprof --pdf heap.pprof > heap.pdf

allocs.pdf: allocs.pprof
	go tool pprof --pdf allocs.pprof > allocs.pdf

mutex.pdf: mutex.pprof
	go tool pprof --pdf mutex.pprof > mutex.pdf

block.pdf: block.pprof
	go tool pprof --pdf block.pprof > block.pdf

goroutine.pdf: goroutine.pprof
	go tool pprof --pdf goroutine.pprof > goroutine.pdf

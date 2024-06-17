.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt:
	go fmt main.go
	go fmt rsort2a.go

vet: fmt
	go vet main.go rsort2a.go

build: vet
	go build -o rdxsort main.go rsort2a.go

profile:
	go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

test:
	go test

clean:
	/bin/rm -f rdxsort gocrdxsort.test
	/bin/rm -f cpu.prof mem.prof profile001.callgraph.out


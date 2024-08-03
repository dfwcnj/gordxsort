.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt:
	go fmt *.go

vet: fmt
	go vet *.go

build: vet
	go build -o rdxsort *.go

profile:
	go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

test:
	go test

clean:
	/bin/rm -f rdxsort gordxsort.test
	/bin/rm -f *.prof *.out
	/bin/rm -f *.pdf


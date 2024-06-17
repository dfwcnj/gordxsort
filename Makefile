.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt:
	go fmt main.go
	go fmt rsort2a.go

vet: fmt
	go vet main.go rsort2a.go

build: vet
	go build -o rdxsort main.go rsort2a.go

test:
	go test

clean:
	/bin/rm -f rdxsort


.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt:
	go fmt main.go
	go fmt rdxsort.go

vet: fmt
	go vet main.go
	go vet rdxsort.go

build: vet
	go build -o rdxsort main.go

test:
	go test

clean:
	/bin/rm -f rdxsort


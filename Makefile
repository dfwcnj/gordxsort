.DEFAULT_GOAL := build

.PHONY:fmt vet build

fmt:
	go fmt main.go
	go fmt rdxsort.go

vet: fmt
	go vet main.go rdxsort.go

build: vet
	go build -o rdxsort main.go rdxsort.go

clean: rdxsort
	/bin/rm rdxsort


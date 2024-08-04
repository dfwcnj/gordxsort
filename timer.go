package main

import (
	"cmp"
	"fmt"
	"slices"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func timeslicessort[S ~[]E, E cmp.Ordered](lns S) {
	a := fmt.Sprintf("slices sort %d", len(lns))
	defer timer(a)()
	slices.Sort(lns)
}

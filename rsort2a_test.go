package main

import (
	"fmt"
	"log"
	"slices"
	"testing"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func timemyradixsort(lns lines) {
	a := fmt.Sprintf("rsort2a %d", len(lns))
	defer timer(a)()
	rsort2a(lns, 0)
}

func timeslicessort(lns []string) {
	a := fmt.Sprintf("slices sort  %d", len(lns))
	defer timer(a)()
	slices.Sort(lns)
}

func Test_rsort2a(t *testing.T) {

	var l int = 32
	var r bool = true
	ls := []int{1, 2, 1 << 4, 1 << 8, 1 << 16, 1 << 20, 1 << 24}

	for _, nl := range ls {

		var lns lines
		log.Print("testing sort of ", nl)
		rsl := randomstrings(nl, l, r)
		if len(rsl) != int(nl) {
			log.Fatal("rsort2a test rsl: wanted len ", nl, " got ", len(rsl))
		}
		for _, s := range rsl {
			bln := []byte(s)
			lns = append(lns, bln)
		}
		if len(lns) != int(nl) {
			log.Print(lns)
			log.Fatal("rsort2a test lns: before sort wanted len ", nl, " got ", len(lns))
		}
		slns := rsort2a(lns, 0)
		if len(slns) != int(nl) {
			//log.Print(slns)
			log.Fatal("rsort2a test slns: after sort wanted len ", nl, " got ", len(slns))
		}
		var ssl []string
		for _, s := range slns {
			ssl = append(ssl, string(s))
		}

		if len(ssl) != int(nl) {
			//log.Print(ssl)
			log.Fatal("rsort2a test ssl: wanted len ", nl, " got ", len(ssl))
		}
		//if !sort.StringsAreSorted(ssl) {
		if !slices.IsSorted(ssl) {
			log.Fatal("rsort2a failed for size ", nl)
		} else {
			log.Print("sort test passed for ", nl)
		}

		timemyradixsort(lns)
		timeslicessort(rsl)
	}
}

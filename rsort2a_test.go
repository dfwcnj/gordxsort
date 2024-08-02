package main

import (
	"cmp"
	"encoding/binary"
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
	a := fmt.Sprintf("rsort2a     %d", len(lns))
	defer timer(a)()
	rsort2a(lns, 0)
}

func timeslicessort[S ~[]E, E cmp.Ordered](lns S) {
	a := fmt.Sprintf("slices sort %d", len(lns))
	defer timer(a)()
	slices.Sort(lns)
}

func Test_rsort2a(t *testing.T) {

	ls := []int{1, 2, 1 << 4, 1 << 8, 1 << 16, 1 << 25}

	for _, nl := range ls {

		var lns lines
		var l int = 32
		var r bool = true
		log.Print("testing sort of ", nl, " random strings")
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
			//log.Print(ulns)
			log.Fatal("rsort2a test ulns: after sort wanted len ", nl, " got ", len(slns))
		}
		var ssl []string
		for _, s := range slns {
			ssl = append(ssl, string(s))
		}

		if !slices.IsSorted(ssl) {
			log.Fatal("rsort2a failed for size ", nl)
		} else {
			log.Print("sort test passed for ", nl)
		}

		timemyradixsort(lns)
		timeslicessort(rsl)

		// }

		// for _, nl := range ls {

		// var lns lines

		log.Print("testing sort of ", nl, " random uints")
		lns = randomuintb(nl)
		if len(lns) != int(nl) {
			log.Fatal("rsort2a test rui: wanted len ", nl, " got ", len(lns))
		}
		slns = rsort2a(lns, 0)
		if len(slns) != int(nl) {
			//log.Print(ulns)
			log.Fatal("rsort2a test ulns: after sort wanted len ", nl, " got ", len(slns))
		}
		var ulns []uint64
		for _, s := range slns {
			ui := binary.BigEndian.Uint64(s)
			ulns = append(ulns, ui)
		}

		if len(ulns) != int(nl) {
			//log.Print(ssl)
			log.Fatal("rsort2a test ssl: wanted len ", nl, " got ", len(ulns))
		}
		if !slices.IsSorted(ulns) {
			log.Fatal("rsort2a failed for size ", nl)
		} else {
			log.Print("sort test passed for ", nl)
		}

		timemyradixsort(lns)
		timeslicessort(ulns)
	}
}

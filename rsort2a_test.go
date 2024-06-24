package main

import (
	"log"
	"sort"
	"testing"
)

func Test_rsort2a(t *testing.T) {

	var l uint = 32
	ls := []uint{1, 2, 1 << 4, 1 << 8, 1 << 16, 1 << 20, 1 << 24}

	for _, nl := range ls {

		var lns lines
		log.Print("testing sort of ", nl)
		rsl := randomstrings(nl, l)
		if len(rsl) != int(nl) {
			log.Fatal("rsl: wanted len ", nl, " got ", len(rsl))
		}
		for _, s := range rsl {
			bln := []byte(s)
			lns = append(lns, bln)
		}
		if len(lns) != int(nl) {
			log.Print(lns)
			log.Fatal("lns: before sort wanted len ", nl, " got ", len(lns))
		}
		slns := rsort2a(lns, 0)
		if len(slns) != int(nl) {
			//log.Print(slns)
			log.Fatal("slns: after sort wanted len ", nl, " got ", len(slns))
		}
		var ssl []string
		for _, s := range slns {
			ssl = append(ssl, string(s))
		}

		if len(ssl) != int(nl) {
			//log.Print(ssl)
			log.Fatal("ssl: wanted len ", nl, " got ", len(ssl))
		}
		for i, _ := range ssl {
			if len(ssl[i]) != int(l) {
				log.Fatal("ssl[i]: wanted len ", l, " got ", len(ssl[i]))
			}
		}
		if !sort.StringsAreSorted(ssl) {
			t.Error("rsort2a failed for size ", nl)
		} else {
			log.Print("sort test passed for ", nl)
		}
	}
}

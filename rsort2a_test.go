package main

import (
	"log"
	"sort"
	"testing"
)

func Test_rsort2a(t *testing.T) {

	var l uint = 32
	ls := []uint{1, 2, 1 << 4, 1 << 8, 1 << 16, 1 << 20, 1 << 24}

	for _, i := range ls {

		var lns lines
		log.Print("testing sort of ", i)
		rsl := randomstrings(i, l)
		if len(rsl) != int(i) {
			log.Fatal("rsl: wanted len ", i, " got ", len(rsl))
		}
		for _, s := range rsl {
			bln := []byte(s)
			lns = append(lns, bln)
		}
		if len(lns) != int(i) {
			log.Print(lns)
			log.Fatal("lns: before sort wanted len ", i, " got ", len(lns))
		}
		slns := rsort2a(lns, 0)
		if len(slns) != int(i) {
			//log.Print(slns)
			log.Fatal("slns: after sort wanted len ", i, " got ", len(slns))
		}
		var ssl []string
		for _, s := range slns {
			ssl = append(ssl, string(s))
		}

		if len(ssl) != int(i) {
			//log.Print(ssl)
			log.Fatal("ssl: wanted len ", i, " got ", len(ssl))
		}
		for i, _ := range ssl {
			if len(ssl[i]) != int(l) {
				log.Fatal("ssl[i]: wanted len ", l, " got ", len(ssl[i]))
			}
		}
		if !sort.StringsAreSorted(ssl) {
			t.Error("rsort2a failed for size ", i)
		} else {
			log.Print("sort test passed for ", i)
		}
	}
}

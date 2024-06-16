package gocrdxsort

import (
	"log"
	"sort"
	"testing"
)

func Test_rsort2a(t *testing.T) {

	var lns lines
	var l uint = 32
	ls := []uint{1, 2, 16, 128, 65536, 1 << 20}

	for _, i := range ls {

		log.Print("testing sort of ", i)
		rsl := randomstrings(i, l)
		if len(rsl) != int(i) {
			log.Fatal("rsl: wanted len ", i, " got ", len(rsl))
		}
		lns = lns[:0]
		for _, s := range rsl {
			bln := []byte(s)
			lns = append(lns, bln)
		}
		if len(lns) != int(i) {
			log.Print(lns)
			log.Fatal("lns: before sort wanted len ", i, " got ", len(lns))
		}
		rsort2a(lns, 0)
		if len(lns) != int(i) {
			log.Print(lns)
			log.Fatal("lns: after sort wanted len ", i, " got ", len(lns))
		}

		var ssl []string
		for _, s := range lns {
			ssl = append(ssl, string(s))
		}
		if len(ssl) != int(i) {
			log.Print(ssl)
			log.Fatal("ssl: wanted len ", i, " got ", len(ssl))
		}
		if !sort.StringsAreSorted(ssl) {
			t.Error("rsort2a failed for size ", i)
		} else {
			log.Print("sort test passed for ", i)
		}
	}
}

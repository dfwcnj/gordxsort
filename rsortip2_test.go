package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"slices"
	"testing"
)

func timemyrsortip2(lns lines) {
	a := fmt.Sprintf("rsortip2     %d", len(lns))
	defer timer(a)()
	rsortip2(lns, 0)
}

func Test_rsortip2(t *testing.T) {

	ls := []int{1 << 5}
	ns := []int{1 << 3, 1 << 4, 1 << 5, 1 << 6, 1 << 16, 1 << 20, 1 << 24}

	for _, ll := range ls {
		for _, nl := range ns {

			var lns lines
			var l int = ll
			var r bool = true
			log.Print("testing rsortip2 of ", nl, " random strings length ", l)
			rsl := randomstrings(nl, l, r)
			if len(rsl) != int(nl) {
				log.Fatal("rsortip2 test rsl: wanted len ", nl, " got ", len(rsl))
			}
			// log.Print("strings generated")
			for _, s := range rsl {
				bln := []byte(s)
				lns = append(lns, bln)
			}
			if len(lns) != int(nl) {
				// log.Print(lns)
				log.Fatal("rsortip2 test lns: before rsortip2 wanted len ", nl, " got ", len(lns))
			}
			// log.Print("strings converted to bytes")
			slns := rsortip2(lns, 0)
			if len(slns) != int(nl) {
				//log.Print(ulns)
				log.Fatal("rsortip2 test ulns: after rsortip2 wanted len ", nl, " got ", len(slns))
			}
			// log.Print("byte strings sorted")
			var ssl []string
			for _, s := range slns {
				ssl = append(ssl, string(s))
			}
			// log.Print("byte strings converted to strings")

			if !slices.IsSorted(ssl) {
				for i, _ := range ssl {
					// log.Println(string(slns[i]))
					log.Println(ssl[i])
				}
				log.Fatal("rsortip2 not sorted for size ", nl)
			} else {
				log.Print("rsortip2 test passed for ", nl)
			}

			log.Print("string rsortip2 comparison")
			timemyrsortip2(lns)
			timeslicessort(rsl)

			log.Print("testing rsortip2 of ", nl, " random uints")
			lns = randomuintb(nl)
			if len(lns) != int(nl) {
				log.Fatal("rsortip2 test rui: wanted len ", nl, " got ", len(lns))
			}
			// log.Print("uint64 byte strings generated")
			slns = rsortip2(lns, 0)
			if len(slns) != int(nl) {
				//log.Print(ulns)
				log.Fatal("rsortip2 test ulns: after rsortip2 wanted len ", nl, " got ", len(slns))
			}
			// log.Print("uint64 byte strings sorted")
			var ulns []uint64
			for _, s := range slns {
				ui := binary.BigEndian.Uint64(s)
				ulns = append(ulns, ui)
			}
			// log.Print("uint64 byte strings converted to uint64")

			if len(ulns) != int(nl) {
				//log.Print(ssl)
				log.Fatal("rsortip2 test ssl: wanted len ", nl, " got ", len(ulns))
			}
			if !slices.IsSorted(ulns) {
				log.Fatal("rsortip2 failed for size ", nl)
			} else {
				log.Print("rsortip2 test passed for ", nl)
			}

			log.Print("uint64 rsortip2 comparison")
			timemyrsortip2(lns)
			timeslicessort(ulns)
		}
	}
}

package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"slices"
	"testing"
)

//func timer(name string) func() {
//	start := time.Now()
//	return func() {
//		fmt.Printf("%s took %v\n", name, time.Since(start))
//	}
//}

func timemyrsortsl(lns lines) {
	a := fmt.Sprintf("rsortsl     %d", len(lns))
	defer timer(a)()
	rsortsl(lns, 0)
}

func Test_rsortsl(t *testing.T) {

	ls := []int{1 << 3, 1 << 4, 1 << 5, 1 << 6}
	ns := []int{1 << 3, 1 << 16, 1 << 25}

	for _, ll := range ls {
		for _, nl := range ns {

			var lns lines
			var l int = ll
			var r bool = true
			log.Print("testing rsortsl of ", nl, " random strings length ", l)
			rsl := randomstrings(nl, l, r)
			if len(rsl) != int(nl) {
				log.Fatal("rsortsl test rsl: wanted len ", nl, " got ", len(rsl))
			}
			for _, s := range rsl {
				bln := []byte(s)
				lns = append(lns, bln)
			}
			if len(lns) != int(nl) {
				log.Print(lns)
				log.Fatal("rsortsl test lns: before rsortsl wanted len ", nl, " got ", len(lns))
			}
			slns := rsortsl(lns, 0)
			if len(slns) != int(nl) {
				//log.Print(ulns)
				log.Fatal("rsortsl test ulns: after rsortsl wanted len ", nl, " got ", len(slns))
			}
			var ssl []string
			for _, s := range slns {
				ssl = append(ssl, string(s))
			}

			if !slices.IsSorted(ssl) {
				log.Fatal("rsortsl failed for size ", nl)
			} else {
				log.Print("rsortsl test passed for ", nl)
			}

			timemyrsortsl(lns)
			timeslicessort(rsl)

			// }

			// for _, nl := range ls {

			// var lns lines

			log.Print("testing rsortsl of ", nl, " random uints")
			lns = randomuintb(nl)
			if len(lns) != int(nl) {
				log.Fatal("rsortsl test rui: wanted len ", nl, " got ", len(lns))
			}
			slns = rsortsl(lns, 0)
			if len(slns) != int(nl) {
				//log.Print(ulns)
				log.Fatal("rsortsl test ulns: after rsortsl wanted len ", nl, " got ", len(slns))
			}
			var ulns []uint64
			for _, s := range slns {
				ui := binary.BigEndian.Uint64(s)
				ulns = append(ulns, ui)
			}

			if len(ulns) != int(nl) {
				//log.Print(ssl)
				log.Fatal("rsortsl test ssl: wanted len ", nl, " got ", len(ulns))
			}
			if !slices.IsSorted(ulns) {
				log.Fatal("rsortsl failed for size ", nl)
			} else {
				log.Print("rsortsl test passed for ", nl)
			}

			timemyrsortsl(lns)
			timeslicessort(ulns)
		}
	}
}

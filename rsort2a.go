package main

import (
	"log"
)

const THRESHOLD int = 1 << 5

//type line []byte
//type lines []line

func binsertionsort(lns lines) lines {
	n := len(lns)
	if n == 1 {
		return lns
	}
	for i := 0; i < n; i++ {
		for j := i; j > 0 && string(lns[j-1]) > string(lns[j]); j-- {
			lns[j], lns[j-1] = lns[j-1], lns[j]
		}
	}
	return lns
}

// bostic
func rsort2a(lns lines, recix int) lines {
	var piles = make([][]line, 256)
	var nc int
	var li int
	nl := len(lns)
	pilelen := make([]int, 256)

	if nl == 0 {
		log.Fatal("rsort2a: 0 len lines: ", recix)
	}
	if nl < THRESHOLD {
		return binsertionsort(lns)
	}

	for i, _ := range lns {

		if recix >= len(lns[i]) {
			continue
		}

		// aooend line to the pile indexed by c
		c := int(lns[i][recix])
		piles[c] = append(piles[c], line(lns[i]))
		if len(piles[c]) == 1 {
			nc++ // number of piles so far
		}
		li = c
	}
	if nc == 1 {
		return binsertionsort(piles[li])
	}

	for i, _ := range piles {
		if len(piles[i]) == 0 {
			continue
		}

		pilelen[i] = len(piles[i])
		// sort pile
		if len(piles[i]) < THRESHOLD {
			piles[i] = binsertionsort(piles[i])
			if len(piles[i]) != pilelen[i] {
				log.Fatal("pilelen[", i, "] ", pilelen[i], "len(piles[i]) ", len(piles[i]))
			}
		} else {
			piles[i] = rsort2a(piles[i], recix+1)
			if len(piles[i]) != pilelen[i] {
				log.Fatal("pilelen[", i, "] ", pilelen[i], "len(piles[i]) ", len(piles[i]))
			}
		}
		nc--
		if nc == 0 {
			break
		}
	}
	var slns lines
	for i, _ := range piles {
		if len(piles[i]) != pilelen[i] {
			log.Fatal("pilelen[", i, "] ", pilelen[i], "len(piles[i]) ", len(piles[i]))
		}
		for j, _ := range piles[i] {
			slns = append(slns, piles[i][j])
		}
	}
	if len(slns) != nl {
		log.Fatal("slns: ", len(slns), " nl ", nl)
	}
	return slns
}

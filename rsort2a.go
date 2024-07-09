package main

import (
	"bytes"
	"log"
)

const THRESHOLD int = 1 << 5

type line []byte
type lines []line

func binsertionsort(lns lines) lines {
	n := len(lns)
	if n == 1 {
		return lns
	}
	for i := 0; i < n; i++ {
		for j := i; j > 0 && bytes.Compare(lns[j-1], lns[j]) > 0; j-- {
			lns[j], lns[j-1] = lns[j-1], lns[j]
		}
	}
	return lns
}

// bostic
func rsort2a(lns lines, recix int) lines {
	var piles = make([][]line, 256)
	var nc int
	nl := len(lns)

	if nl == 0 {
		log.Fatal("rsort2a: 0 len lines: ", recix)
	}
	if nl < THRESHOLD {
		return binsertionsort(lns)
	}

	for i, _ := range lns {
		var c int

		if len(lns[i]) == 0 {
			log.Fatal("rsort2a 0 length string")
		}
		if recix >= len(lns[i]) {
			c = 0
		} else {
			c = int(lns[i][recix])
		}
		piles[c] = append(piles[c], line(lns[i]))
		if len(piles[c]) == 1 {
			nc++ // number of piles so far
		}
	}
	if len(piles[0]) > 1 {
		piles[0] = binsertionsort(piles[0])
	}
	if nc == 1 {
		if len(piles[0]) > 1 {
			return piles[0]
		}
		return binsertionsort(lns)
	}

	for i, _ := range piles {
		if len(piles[i]) == 0 {
			continue
		}

		// sort pile
		if len(piles[i]) < THRESHOLD {
			piles[i] = binsertionsort(piles[i])
		} else {
			piles[i] = rsort2a(piles[i], recix+1)
		}
		nc--
		if nc == 0 {
			break
		}
	}
	var slns lines
	for i, _ := range piles {
		for j, _ := range piles[i] {
			slns = append(slns, piles[i][j])
		}
	}
	return slns
}

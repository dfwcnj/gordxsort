package main

import (
	"bytes"
	"log"
)

// type line []byte
// type lines []line

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
	const THRESHOLD int = 1 << 5
	var sizes = make([]int, 256)
	var piles = make([][]line, 256)
	var nc int
	nl := len(lns)

	if nl == 0 {
		log.Fatal("rsort2a: 0 len lines: ", recix)
	}
	if nl < THRESHOLD {
		return binsertionsort(lns)
	}

	// count the number of lines that will fall each pile
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
		sizes[c]++
	}
	// preallocate the piles so that they don't have to be resized
	for i, _ := range sizes {
		if sizes[i] != 0 {
			piles[i] = make([]line, 0, sizes[i])
		}
	}

	// deal lines into piles
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

	// sort the piles
	if nc == 1 {
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

	// combine the sorted piles
	var slns lines
	for i, _ := range piles {
		for j, _ := range piles[i] {
			slns = append(slns, piles[i][j])
		}
	}
	return slns
}

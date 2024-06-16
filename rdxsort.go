package gocrdxsort

const THRESHOLD int = 1 << 5

type line []byte
type lines []line

func simplesort(lns lines) {
	n := len(lns)
	if n == 1 {
		return
	}
	for i := 0; i < len(lns); i++ {
		for j := i; j > 0 && string(lns[j-1]) > string(lns[j]); j-- {
			lns[j], lns[j-1] = lns[j-1], lns[j]
		}
	}
}

// bostic
func rsort2a(lns lines, recix int) {
	var piles = make([][]line, 256)
	var nc int

	if len(lns) == 0 {
		return
	}
	if len(lns) < THRESHOLD {
		simplesort(lns)
		return
	}

	for _, l := range lns {

		if recix >= len(l) {
			continue
		}

		// aooend line to the pile indexed by c
		c := int(l[recix])
		piles[c] = append(piles[c], line(l))
		if len(piles[c]) == 1 {
			nc++ // number of piles so far
		}
	}

	for _, p := range piles {
		if len(p) == 0 {
			continue
		}
		// sort pile
		rsort2a(p, recix+1)
		nc--
		if nc == 0 {
			break
		}
	}
	clear(lns)
	for _, p := range piles {
		if len(p) == 0 {
			continue
		}
		for _, l := range p {
			lns = append(lns, l)
		}
	}
}

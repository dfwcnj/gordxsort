package gocrdxsort

const THRESHOLD int = 1 << 5

type line []byte
type lines []line

func simplesort(lns lines) {
	n := len(lns)
	if n == 1 {
		return
	}
	var j int
	for i := 1; i < n; i++ {
		l := lns[i]
		for j = i - 1; j >= 0; j-- {
			// make room
			if string(lns[j]) > string(l) {
				lns[j+1] = lns[j]
			} else {
				break
			}
		}
		lns[j+1] = l
	}
}

// bostic
func rsort2a(lns lines, recix int) {
	var piles = make([][]line, 256)
	var nc int

	if len(lns) < THRESHOLD {
		simplesort(lns)
		return
	}

	for _, l := range lns {

		if recix >= len(l) {
			continue
		}

		// aooend offset in line to the pile indexed by c
		c := int(l[recix])
		piles[c] = append(piles[c], line(l))
		if len(piles[c]) == 1 {
			nc++ // number of piles so far
		}
	}

	// no lns with offset ix
	if nc == 0 {
		return // if no piles, done
	}

	for _, p := range piles {
		if len(p) == 0 {
			continue
		}
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

package gocrdxsort

const THRESHOLD int = 1 << 5

type line []byte

// bostic
func rsort2a(lines []line, recix int) {
	var piles = make([][]int, 256)
	var tlines = make([]line, len(lines))
	var nc int
	var ix int = recix

	// iterate over indices in lines
	for {
		for i, l := range lines {
			if ix >= len(l) {
				continue
			}
			c := int(l[ix])

			// aooend offset in line to the pile indexed by c
			piles[c] = append(piles[c], i)
			if len(piles[c]) == 1 {
				nc++ // number of piles so far
			}
		}
		if nc == 0 {
			return // if no piles, done
		}
		for i, _ := range piles {
			// copy piles to temporary storage in order
			for j, _ := range piles[i] {
				tlines = append(tlines, lines[piles[i][j]])
			}
			nc--
			if nc == 0 {
				break
			}
		}
		clear(piles)
		clear(lines)
		copy(tlines, lines)
		ix++
	}
}

package main

import (
	"log"
)

//https://www.usenix.org/legacy/publications/compsystems/1993/win_mcilroy.pdf
// https://arxiv.org/pdf/0706.4107
// https://stackoverflow.com/questions/463105/in-place-radix-sort
// https://informatica.vu.lt/journal/INFORMATICA/article/953/info

// type line []byte
// type lines []line

// Al-Badarneh
func rsortip(lns lines, recix int) lines {
	// log.Print("rsortip ", recix)
	const THRESHOLD int = 1 << 3
	if len(lns) <= THRESHOLD {
		// log.Print("rsortip inssort ix ", recix)
		return inssort(lns)
	}
	type bin struct {
		count, start, end int
	}
	bins := make([]bin, 256)

	// count the sizes of the bins
	// log.Print("computing bin sizes")
	for i, _ := range lns {
		var c int
		if len(lns[i]) == 0 {
			log.Fatal("rsortip 0 length string")
		}
		if recix >= len(lns[i]) {
			c = 0
		} else {
			c = int(lns[i][recix])
		}
		bins[c].count++
	}

	// compute the offsets of the bins
	// log.Print("computing bin dimensions")
	var offset int
	var nc int
	for i, _ := range bins {
		if bins[i].count != 0 {
			bins[i].start = offset
			bins[i].end = offset
			offset += bins[i].count
			nc++
		}
	}

	// log.Print("rsortip nc ", nc)
	if bins[0].count != 0 {
		inssort(lns[:bins[0].count-1])
		if nc == 1 {
			return lns
		}
	}
	// swap ln[i] with current bin contents
	// update bins[i].end
	// repeat unto lns[i] is where it should be
	// log.Print("loop swapping")
	for i, _ := range lns {
		//log.Print("	looping lns index ", i)
		for {
			var bin int
			if recix >= len(lns[i]) {
				bin = 0
			} else {
				bin = int(lns[i][recix])
			}
			dst := bins[bin].end
			if i >= bins[bin].start {
				break
			}
			if bins[bin].end >= bins[bin].start+bins[bin].count {
				log.Fatal("rsortip fatal dst bin ", bin, " start ", bins[bin].start, " count ", bins[bin].count, " used ", bins[bin].end-bins[bin].start)
			}
			lns[i], lns[dst] = lns[dst], lns[i]
			bins[bin].end++
		}
	}

	for i, _ := range bins {
		if i == 0 || bins[i].count == 0 {
			continue
		}
		rsortip(lns[bins[i].start:bins[i].end], recix+1)
		nc--
		if nc == 0 {
			break
		}
	}

	if recix == 0 {
		inssort(lns)
	}
	return lns
}

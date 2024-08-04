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
func rsortip2(lns lines, recix int) lines {
	// log.Print("rsortip2 ", recix)
	const THRESHOLD int = 1 << 2
	if len(lns) <= THRESHOLD {
		// log.Print("rsortip2 inssort ix ", recix)
		return inssort(lns)
	}
	type bin struct {
		count, start, end int
	}
	bins := make([]bin, 256)
	type move struct {
		bin, from, to int
	}
	moves := make([]move, len(lns))

	// count the sizes of the bins
	// log.Print("computing bin sizes")
	for i, _ := range lns {
		var c int
		if len(lns[i]) == 0 {
			log.Fatal("rsortip2 0 length string")
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

	// log.Print("rsortip2 nc ", nc)
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
		var bin int
		if recix >= len(lns[i]) {
			bin = 0
		} else {
			bin = int(lns[i][recix])
		}
		dst := bins[bin].end
		moves[i].bin = bin
		moves[i].from = i
		moves[i].to = dst
		bins[bin].end++
	}

	for i, _ := range moves {
		for {
			if i == moves[i].to {
				break
			}
			lns[i], lns[moves[i].to] = lns[moves[i].to], lns[i]
			moves[i], moves[moves[i].to] = moves[moves[i].to], moves[i]
		}
	}

	for i, _ := range bins {
		if i == 0 || bins[i].count == 0 {
			continue
		}
		rsortip2(lns[bins[i].start:bins[i].end], recix+1)
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

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

// type line []byte
// type lines []line

func main() {
	var fn string
	flag.StringVar(&fn, "file", "", "name of file to sort")
	flag.Parse()
	var lns lines

	var err error

	fp := os.Stdin
	if fn != "" {
		fp, err = os.Open(fn)
		if err != nil {
			log.Fatal(err)
		}
		defer fp.Close()
	}

	scanner := bufio.NewScanner(fp)
	// option, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		l := scanner.Text()
		bln := []byte(l)
		lns = append(lns, bln)
	}
	slns := rsortsl(lns, 0)
	for i, _ := range slns {
		fmt.Println(string(slns[i]))
	}

}

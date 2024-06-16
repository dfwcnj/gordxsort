package gocrdxsort

import (
	//	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	// "strings"
)

func main() {
	var fn string
	flag.StringVar(&fn, "file", "", "name of file to sort")
	flag.Parse()

	var fp = os.Stdin
	var err error
	const bsz int64 = 1 << 20
	var ba [bsz]byte
	var buf []byte = ba[0:]
	var offset int64

	if fn != "" {
		fp, err = os.Open(fn)
		if err != nil {
			log.Fatal(err)
		}
	}

	for {
		n, err := fp.ReadAt(buf, offset)
		if n == 0 {
			if !errors.Is(err, io.EOF) {
				fmt.Println("ReadAt", offset, n, err)
			}
		}
		offset = offset + int64(n)

	}
}

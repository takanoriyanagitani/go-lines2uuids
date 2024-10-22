package rdr2iter

import (
	"bufio"
	"iter"
)

func ScannerToIter(s *bufio.Scanner) iter.Seq[[]byte] {
	return func(yield func([]byte) bool) {
		for s.Scan() {
			var line []byte = s.Bytes()
			var ok bool = yield(line)
			if !ok {
				return
			}
		}
	}
}

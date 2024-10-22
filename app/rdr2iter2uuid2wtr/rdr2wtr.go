package rdr2wtr

import (
	"bufio"
	"context"
	"io"
	"iter"
	"os"

	l2u "github.com/takanoriyanagitani/go-lines2uuids"
	s2i "github.com/takanoriyanagitani/go-lines2uuids/input/rdr/rdr2iter"
	u2o "github.com/takanoriyanagitani/go-lines2uuids/output"
)

type ReaderToIterToUuidToOutput struct {
	l2u.LineToUuid
	u2o.OutputUuid
}

func (r ReaderToIterToUuidToOutput) LinesToUuidsToOutput(
	ctx context.Context,
	lines iter.Seq[[]byte],
) error {
	for line := range lines {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		u, e := r.LineToUuid(ctx, line)
		if nil != e {
			return e
		}

		e = r.OutputUuid(ctx, u)
		if nil != e {
			return e
		}
	}
	return nil
}

func (r ReaderToIterToUuidToOutput) ReaderToUuidsToOutput(
	ctx context.Context,
	rdr io.Reader,
) error {
	var br io.Reader = bufio.NewReader(rdr)
	var s *bufio.Scanner = bufio.NewScanner(br)
	var i iter.Seq[[]byte] = s2i.ScannerToIter(s)
	return r.LinesToUuidsToOutput(ctx, i)
}

func (r ReaderToIterToUuidToOutput) StdinToUuidsToOutput(
	ctx context.Context,
) error {
	return r.ReaderToUuidsToOutput(ctx, os.Stdin)
}

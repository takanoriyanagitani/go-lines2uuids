package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"

	l2u "github.com/takanoriyanagitani/go-lines2uuids"
	h2u "github.com/takanoriyanagitani/go-lines2uuids/hex"

	u2o "github.com/takanoriyanagitani/go-lines2uuids/output"
	u2w "github.com/takanoriyanagitani/go-lines2uuids/output/wtr"

	r2w "github.com/takanoriyanagitani/go-lines2uuids/app/rdr2iter2uuid2wtr"
)

var hex2bytes h2u.HexStringToBytes = h2u.HexStrToBytesStd
var line2uuid l2u.LineToUuid = hex2bytes.ToLineToUuid()

var uuid2writer u2w.UuidToWriter = u2w.UuidToWriterDefault

func main() {
	var wtr io.Writer = os.Stdout
	var bw *bufio.Writer = bufio.NewWriter(wtr)
	defer bw.Flush()

	var uuid2output u2o.OutputUuid = uuid2writer.ToOutputUuid(bw)
	reader2iter2uuid2output := r2w.ReaderToIterToUuidToOutput{
		LineToUuid: line2uuid,
		OutputUuid: uuid2output,
	}

	e := reader2iter2uuid2output.StdinToUuidsToOutput(context.Background())
	if nil != e {
		log.Printf("%v\n", e)
	}
}

package hex2ints2uuid

import (
	"context"
	"encoding/hex"
	"errors"

	l2u "github.com/takanoriyanagitani/go-lines2uuids"
)

var (
	ErrInvalidUuid error = errors.New("invalid uuid")
)

type HexStringToBytes func(context.Context, [32]uint8) (l2u.Uuid, error)

func HexStrToBytesStd(_ context.Context, h [32]uint8) (l2u.Uuid, error) {
	var buf [16]uint8

	cnt, e := hex.Decode(buf[:], h[:])
	if nil != e {
		return buf, e
	}

	if 16 != cnt {
		return buf, ErrInvalidUuid
	}

	return buf, nil
}

func (h HexStringToBytes) ToLineToUuid() l2u.LineToUuid {
	return func(ctx context.Context, line []byte) (l2u.Uuid, error) {
		var buf [32]byte
		var copied int = copy(buf[:], line)
		if 32 != copied {
			return l2u.Uuid{}, ErrInvalidUuid
		}
		return h(ctx, buf)
	}
}

package sha256uuid

import (
	"context"

	"crypto/sha256"

	l2u "github.com/takanoriyanagitani/go-lines2uuids"
)

type LineToSha256 func(context.Context, []byte) ([32]uint8, error)

func (l LineToSha256) ToLineToUuid() l2u.LineToUuid {
	return func(ctx context.Context, line []byte) (l2u.Uuid, error) {
		s8, e := l(ctx, line)
		if nil != e {
			return l2u.Uuid{}, e
		}

		return [16]uint8{
			s8[0x00] ^ s8[0x10],
			s8[0x01] ^ s8[0x11],
			s8[0x02] ^ s8[0x12],
			s8[0x03] ^ s8[0x13],
			s8[0x04] ^ s8[0x14],
			s8[0x05] ^ s8[0x15],
			s8[0x06] ^ s8[0x16],
			s8[0x07] ^ s8[0x17],
			s8[0x08] ^ s8[0x18],
			s8[0x09] ^ s8[0x19],
			s8[0x0a] ^ s8[0x1a],
			s8[0x0b] ^ s8[0x1b],
			s8[0x0c] ^ s8[0x1c],
			s8[0x0d] ^ s8[0x1d],
			s8[0x0e] ^ s8[0x1e],
			s8[0x0f] ^ s8[0x1f],
		}, nil
	}
}

func LineToSha256StdDefault(_ context.Context, line []byte) ([32]uint8, error) {
	return sha256.Sum256(line), nil
}

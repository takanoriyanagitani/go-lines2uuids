package hex2ints2uuid_test

import (
	"context"
	"testing"

	h2u "github.com/takanoriyanagitani/go-lines2uuids/hex"
)

func TestHex(t *testing.T) {
	t.Parallel()

	t.Run("HexStrToBytesStd", func(t *testing.T) {
		t.Parallel()

		t.Run("empty", func(t *testing.T) {
			t.Parallel()

			var h2b h2u.HexStringToBytes = h2u.HexStrToBytesStd
			var empty [32]uint8
			_, e := h2b(context.Background(), empty)
			if nil == e {
				t.Fatalf("must fail\n")
			}
		})

		t.Run("zero", func(t *testing.T) {
			t.Parallel()

			var h2b h2u.HexStringToBytes = h2u.HexStrToBytesStd
			var zero [32]uint8 = [32]uint8{
				0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
				0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
				0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
				0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
			}
			b16, e := h2b(context.Background(), zero)
			if nil != e {
				t.Fatalf("unexpected err: %v\n", e)
			}
			var z2 [16]uint8
			if z2 != b16 {
				t.Fatalf("expected: %v, got: %v\n", z2, b16)
			}
		})
	})
}

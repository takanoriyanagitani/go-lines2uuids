package lines2uuids

import (
	"context"
)

type Uuid [16]uint8

type LineToUuid func(context.Context, []byte) (Uuid, error)

package output

import (
	"context"

	l2u "github.com/takanoriyanagitani/go-lines2uuids"
)

type OutputUuid func(context.Context, l2u.Uuid) error

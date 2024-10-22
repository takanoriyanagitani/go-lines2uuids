package wtr

import (
	"context"
	"io"

	l2u "github.com/takanoriyanagitani/go-lines2uuids"
	u2o "github.com/takanoriyanagitani/go-lines2uuids/output"
)

type UuidToWriter func(context.Context, l2u.Uuid, io.Writer) error

func (u UuidToWriter) ToOutputUuid(wtr io.Writer) u2o.OutputUuid {
	return func(ctx context.Context, id l2u.Uuid) error {
		return u(ctx, id, wtr)
	}
}

func UuidToWriterDefault(_ context.Context, u l2u.Uuid, w io.Writer) error {
	var s []byte = u[:]
	_, e := w.Write(s)
	return e
}

package model

import (
	"time"

	"github.com/stormasm/noms/go/d"
	"github.com/stormasm/noms/go/marshal"
	"github.com/stormasm/noms/go/types"
)

type CommitMeta struct {
	Date string
}

func NewCommitMeta() CommitMeta {
	return CommitMeta{
		time.Now().Format(time.RFC3339),
	}
}

func (c CommitMeta) Marshal() types.Struct {
	v, err := marshal.Marshal(c)
	d.Chk.NoError(err)
	return v.(types.Struct)
}

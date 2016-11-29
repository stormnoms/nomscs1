// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package spec

import (
	"github.com/stormasm/noms/go/chunks"
	"github.com/stormasm/noms/go/d"
)

type refCountingBoltStore struct {
	*chunks.BoltDBStore
	refCount int
	closeFn  func()
}

func newrefCountingBoltStore(path string, closeFn func()) *refCountingBoltStore {
	return &refCountingBoltStore{chunks.NewBoltDBStoreUseFlags(path, ""), 1, closeFn}
}

func (r *refCountingBoltStore) AddRef() {
	r.refCount++
}

func (r *refCountingBoltStore) Close() (err error) {
	d.PanicIfFalse(r.refCount > 0)
	r.refCount--
	if r.refCount == 0 {
		err = r.BoltDBStore.Close()
		r.closeFn()
	}
	return
}

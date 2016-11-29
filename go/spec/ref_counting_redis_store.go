// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package spec

import (
	"github.com/stormasm/noms/go/chunks"
	"github.com/stormasm/noms/go/d"
)

type refCountingRedisStore struct {
	*chunks.LevelDBStore
	refCount int
	closeFn  func()
}

func newrefCountingRedisStore(path string, closeFn func()) *refCountingRedisStore {
	return &refCountingRedisStore{chunks.NewLevelDBStoreUseFlags(path, ""), 1, closeFn}
}

func (r *refCountingRedisStore) AddRef() {
	r.refCount++
}

func (r *refCountingRedisStore) Close() (err error) {
	d.PanicIfFalse(r.refCount > 0)
	r.refCount--
	if r.refCount == 0 {
		err = r.LevelDBStore.Close()
		r.closeFn()
	}
	return
}

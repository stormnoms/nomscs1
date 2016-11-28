// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

// Package spec provides builders and parsers for spelling Noms databases,
// datasets and values.
package spec

import (
	"github.com/stormasm/noms/go/chunks"
)

func getBoltStore(path string) chunks.ChunkStore {
	if store, ok := ldbStores[path]; ok {
		store.AddRef()
		return store
	}

	store := newRefCountingLdbStore(path, func() {
		delete(ldbStores, path)
	})
	ldbStores[path] = store
	return store
}

func getRedisStore(path string) chunks.ChunkStore {
	if store, ok := ldbStores[path]; ok {
		store.AddRef()
		return store
	}

	store := newRefCountingLdbStore(path, func() {
		delete(ldbStores, path)
	})
	ldbStores[path] = store
	return store
}

// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package spec

import (
	//"fmt"
	"io/ioutil"
	//"os"
	"path"
	"testing"

	"github.com/stormasm/noms/go/chunks"
	"github.com/stormasm/noms/go/datas"
	"github.com/stormasm/noms/go/types"
	"github.com/attic-labs/testify/assert"
)

func TestBoltDatabaseSpec(t *testing.T) {
	assert := assert.New(t)

	run := func(prefix string) {
		tmpDir, err := ioutil.TempDir("/tmp67", "spec_test")
		assert.NoError(err)
		//defer os.RemoveAll(tmpDir)

		s1 := types.String("corvallis")
		s2 := types.String("raton")
		s3 := types.String("santafe")

		// Existing database in the database are read from the spec.
		store1 := path.Join(tmpDir, "store1")
		
		cs := chunks.NewBoltStoreUseFlags(store1, "ralph")
		db := datas.NewDatabase(cs)
		db.WriteValue(s1)
		db.Close() // must close immediately to free bolt

		cs = chunks.NewBoltStoreUseFlags(store1, "bill")
		db = datas.NewDatabase(cs)
		db.WriteValue(s2)
		db.Close() // must close immediately to free bolt

		cs = chunks.NewBoltStoreUseFlags(store1, "stu")
		db = datas.NewDatabase(cs)
		db.WriteValue(s3)
		db.Close() // must close immediately to free bolt


/*
		spec1, err := ForDatabase(prefix + store1)
		assert.NoError(err)
		defer spec1.Close()

		assert.Equal("bolt", spec1.Protocol)
		assert.Equal(store1, spec1.DatabaseName)

		assert.Equal(s, spec1.GetDatabase().ReadValue(s.Hash()))

		// New databases can be created and read/written from.

		store2 := path.Join(tmpDir, "store2")
		spec2, err := ForDatabase(prefix + store2)
		assert.NoError(err)
		defer spec2.Close()

		assert.Equal("bolt", spec2.Protocol)
		assert.Equal(store2, spec2.DatabaseName)

		db = spec2.GetDatabase()
		db.WriteValue(s)
		assert.Equal(s, db.ReadValue(s.Hash()))
*/
	}

	run("bolt:")
}

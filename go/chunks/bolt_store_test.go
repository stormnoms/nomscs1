// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package chunks

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/attic-labs/testify/suite"
)

func TestBoltStoreTestSuite(t *testing.T) {
	suite.Run(t, &BoltStoreTestSuite{})
}

type BoltStoreTestSuite struct {
	ChunkStoreTestSuite
	factory Factory
	dir     string
}

func (suite *BoltStoreTestSuite) SetupTest() {
	var err error
	suite.dir, err = ioutil.TempDir(os.TempDir(), "")
	suite.NoError(err)
	suite.factory = NewBoltStoreFactory(suite.dir, false)
	store := suite.factory.CreateStore("htestb").(*BoltStore)
	suite.putCountFn = func() int {
		return int(store.putCount)
	}

	suite.Store = store
}

func (suite *BoltStoreTestSuite) TearDownTest() {
	suite.Store.Close()
	suite.factory.Shutter()
	os.Remove(suite.dir)
}

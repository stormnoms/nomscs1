// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package chunks

import (
	"io/ioutil"
	//"os"
	"testing"

	"github.com/attic-labs/testify/suite"
)

func TestBoltStoreBatchTestSuite(t *testing.T) {
	suite.Run(t, &BoltStoreBatchTestSuite{})
}

type BoltStoreBatchTestSuite struct {
	BoltStoreCommonBatchTestSuite
	factory Factory
	dir     string
}

func (suite *BoltStoreBatchTestSuite) SetupTest() {
	var err error
	//suite.dir, err = ioutil.TempDir(os.TempDir(), "")
	suite.dir, err = ioutil.TempDir("/tmp67", "")
	suite.NoError(err)
	suite.factory = NewBoltStoreFactory(suite.dir, false)
	store := suite.factory.CreateStore("hbatchb").(*BoltStore)
	suite.putCountFn = func() int {
		return int(store.putCount)
	}

	suite.Store = store
}

func (suite *BoltStoreBatchTestSuite) TearDownTest() {
	suite.Store.Close()
	suite.factory.Shutter()
	//os.Remove(suite.dir)
}

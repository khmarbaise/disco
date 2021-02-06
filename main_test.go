// Copyright 2021 The Disco Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main_test

import (
	"github.com/khmarbaise/disco/modules/execute"
	"testing"
)

//Test_Main_first Integration test to execute our own executable within a test.
func Test_Main_first(t *testing.T) {
	t.Run("Execute disco with --help", func(t *testing.T) {
		//Execute our own produced executable for testing purposes.
		execute.ExternalCommand("./disco", "--help")
	})
}

//Second...
func Test_Main_Second(t *testing.T) {
	t.Run("Execute disco xxx", func(t *testing.T) {
		//Execute our own produced executable for testing purposes.
		execute.ExternalCommand("./disco", "dist", "--help")
	})
}

//Second...
func Test_Main_Thrid(t *testing.T) {
	t.Run("Execute disco xxx", func(t *testing.T) {
		//Execute our own produced executable for testing purposes.
		execute.ExternalCommand("./disco", "dist", "--name", "oracle")
	})
}

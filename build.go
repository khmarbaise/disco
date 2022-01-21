// Copyright 2021, 2022 The Disco Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
//go:build vendor
// +build vendor

package main

// Libraries that are included to vendor utilities used during build.
// These libraries will not be included in a normal compilation.

import (
	// for vet
	_ "code.gitea.io/gitea-vet"
)

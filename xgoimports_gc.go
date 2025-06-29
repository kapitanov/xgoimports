// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// This file is a modified version of the original goimports tool from the Go project.
// Its original source can be found at:
// https://cs.opensource.google/go/x/tools/+/refs/tags/v0.34.0:cmd/goimports/goimports_gc.go
//
// The modified version is licensed under the BSD 3-Clause License (see LICENSE and NOTICE files).
// Copyright 2025 Albert Kapitanov.
// SPDX-License-Identifier: BSD-3-Clause

//go:build gc

package main

import (
	"flag"
	"runtime/trace"
)

var traceProfile = flag.String("trace", "", "trace profile output")

func doTrace() func() {
	if *traceProfile != "" {
		bw, flush := bufferedFileWriter(*traceProfile)
		trace.Start(bw)
		return func() {
			trace.Stop()
			flush()
		}
	}
	return func() {}
}

// Copyright 2023 The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"goki.dev/gi/v2/gi"
	"goki.dev/gi/v2/gimain"
	"goki.dev/gi/v2/giv"
	"goki.dev/goosi/events"
)

func main() { gimain.Run(app) }

func app() {
	b := gi.NewAppBody("files")

	fv := giv.NewFileView(b)
	fv.OnDoubleClick(func(e events.Event) {
		if fv.SelectedDoubleClick {
			gi.OpenURL(fv.SelectedFile())
		}
	})

	b.NewWindow().Run().Wait()
}

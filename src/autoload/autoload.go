// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package autoload

import (
	"github.com/golang-framework/mvc"
	err "github.com/golang-framework/mvc/modules/error"
	"src/storage"
	"src/storage/language"
)

type autoload struct {
}

func init() { (&autoload{}).src(mvc.New()) }

func (ad *autoload) src(fw *mvc.Framework) {
	ad.before()

	// initialized self error message
	fw.Err = &err.M{EMsg: storage.Em}
	fw.Fw()

	// initialized Translation
	fw.FwTranslation(language.Tm)

	// initialized router
	fw.Route.E, fw.Route.M = ad.mvcInitializedRouter()
	fw.FwRouter()

	fw.Run()
}

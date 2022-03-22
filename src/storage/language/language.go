// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package language

import (
	"github.com/golang-framework/mvc/storage"
)

var Tm = &storage.E{
	storage.EN: *en,
	storage.CN: *cn,
}

// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package storage

import (
	"github.com/golang-framework/mvc/storage"
)

var Tm = &storage.E{
	storage.EN: {
		"TEST_T_1": "test success 1",
		"TEST_T_2": "test success 2",
	},
	storage.CN: {
		"TEST_T_1": "测试成功 1",
		"TEST_T_2": "测试成功 2",
	},
}

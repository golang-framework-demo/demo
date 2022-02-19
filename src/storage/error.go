// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package storage

import (
	err "github.com/golang-framework/mvc/modules/error"
	"github.com/golang-framework/mvc/storage"
)

const (
	KeyDemo10001 = "demo_msg_key_10001"

	valDemo10001 = "[10001]Show Error Success"
)

var Em = &storage.E{
	"demo": {
		KeyDemo10001: valDemo10001,
	},
}

func Ed(k string, content ...interface{}) error {
	pfx := "demo"
	p(&pfx, k)

	return err.Err(pfx, k, content...)
}

func Nd(k string) string {
	pfx := "demo"
	p(&pfx, k)

	return err.Num(pfx, k)
}

func p(pfx *string, k string) {
	msg := (*Em)[*pfx]
	_, ok := msg[k]
	if ok == false {
		*pfx = storage.ErrPrefix
	}
}

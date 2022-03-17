// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package demo

import "src/app/db/repositories/demo"

const (
	db = "demo"
)

type (
	Demo    demo.Demo
	D       demo.D
	Details demo.DemoDetails

	DemoToDetails struct {
		Demo    `xorm:"extends"`
		Details `xorm:"extends"`
	}

	DemoToD struct {
		Demo `xorm:"extends"`
		D    `xorm:"extends"`
	}
)

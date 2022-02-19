// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-framework/mvc/storage"
	"src/app/services/demo/service"
	store "src/storage"
)

type DemoController struct {
	src *service.DemoService
}

func NewDemoController() *DemoController {
	return &DemoController{
		src: service.NewDemoService(),
	}
}

func (c *DemoController) TestDemo(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, storage.Y{
		"demo":    "success",
		"service": c.src.Demo(),
		"error": fmt.Sprintf(
			"%v: %v",
			store.Nd(store.KeyDemo10001),
			store.Ed(store.KeyDemo10001),
		),
	})
	ctx.Abort()
	return
}

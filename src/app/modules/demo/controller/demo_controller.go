// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-framework/mvc/storage"
	"src/app/modules/demo/service"
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

func (c *DemoController) SetRedisDemo(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, c.src.SetRedisDemo())
	ctx.Abort()

	return
}

func (c *DemoController) GetRedisDemo(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, c.src.GetRedisDemo())
	ctx.Abort()

	return
}

func (c *DemoController) Demo(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, c.src.GetDemo())
	ctx.Abort()

	return
}

func (c *DemoController) DemoAll(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, c.src.GetDemoAll())
	ctx.Abort()

	return
}

func (c *DemoController) DemoDetails(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, c.src.GetDemoDetails())
	ctx.Abort()

	return
}

func (c *DemoController) DemoD(ctx *gin.Context) {
	ctx.JSON(storage.StatusOK, c.src.GetDd())
	ctx.Abort()

	return
}
